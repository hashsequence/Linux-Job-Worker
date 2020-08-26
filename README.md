# Linux-Job-Worker
Distributed Linux Job Worker

## Summary
Implement a prototype job worker service that provides an API to run arbitrary Linux processes.

## Requirements

### API
* Job worker should provide RPC API to start/stop/query status and get an output of a running job process. Any RPC mechanism that works for the task and is familiar to you is OK: GRPC, HTTPS/JSON API or anything else that can guarantee secure and reliable client-server communication.

* The API should provide simple but secure authentication and authorization mechanism.

### Client

* Client command should be able to connect to worker service and schedule several jobs

* Client should be able to query result of the job execution and fetch the logs

## Design Document

### Limitations And Scope

#### Data Management
Though the use of a database to store persistant data would be ideal, I will be instead storing the outputs and error outputs into logs stored on 
the file storage of the linux worker. The logs will be generated at start time with the foldername <pid>_<startingTimestamp>
The client will have to maintain a queue of commands, the client can query the server to see what is running to determine what can be killed,
or the client can store a data store pids that were recieved from the execution of start commands

#### Scale
The scope of this project would only deal with a single linux worker interfacing with one client

### API Design

#### Start

* The Start Api is called with a StartRequest that has the client's command and required arguments and optional env, dir params

* The Process be executed, two logs will be generated, the standard output and standard error will be redirected into a log file, one for stdout and anothe for stderr and both logs will be located in the same folder with the foldername <pid>_<startingTimestamp>, a response with the pid is sent, outputs should be empty since it was just started

* when the process is completed or exit out due to errors the log folder will have a timestamp appended so pids can be reused

* last n (configured via a batch_size constant) lines of logs will be outputted to the output and errOutput of the response and sent to the client via a response when done 

* the Start Api will be a server streaming API and have to stream two responses per request, one for starting the command and one for ending, unless the starting command failed, in any case the stream will close on the server side when its done with the request

```
func Start(StartRequest) returns (StartResponse, error)

type StartRequest {
    //path of a program or executable
    string command
    //arguments to invoke program or executable with
    []string args 
    // environment variables for the execution
    //OPTIONAL variable
    []string env
    // current working directory of the execution
    //OPTIONAL variable, default will be the working directory of server 
    string dir 
}

type StartResponse {
    //process Id of the running process executed by the command
    //will be 0 if process failed to execute
    int pid 
    //stdout(standard output) of command/program executed
    string output
    //stderr(standard Error) of command/program executed
    string ErrorOutput 
    //if the command was finished it will be 0, otherwise if command successfully completes it will be 1
    bool finished
    //timestamp returns the starting time if just executed or finish time if finished
    string timeStamp 
}

```
#### Stop

* The User should be able to stop the request based on the pid and starting timestamp

* when stopped the pid should be killed 

* a response will be sent to client indicating process have been stopped along with a the last n lines of the logs

* the Stop API will be a unary api since we want a response to indicate the process is stopped when requested
```
type StopRequest {
    int pid
    string processName
}

type StopResponse {
    string Output
    string errorOutput
}

func Stop(StopRequest) returns (StopResponse, error)

```
* The client should be able to check what processes are still running

* The Client should be able fetch the logs and fetch the status of the process by pid and starting timestamp

* Client should be able to read the logs into memory and send it and  will only grab up to a certain amout of bytes from
the file to prevent maxing out string size or memory 

* Query API should be a unary API since we need an immediate response from state of system at request time

* splitting up Query API into

    * QueryOneProcess:

        * get information on one pid

    * QueryRunningProcesses:

        * get list of running processes from server

Query 
```

type QueryOneRequest {
    int pid
    string startingTimeStamp
    string processName
    string errorLog
    string outputLog
}

type QueryOneResponse {
    //output of requested range of output log
    string output
    //output of requested range of error log
    string errorOutput
    //return status of execution, 0 for finished or stopped, 1 for still running
    bool running
}

type RunningProcess {
    int pid
    string startingTimeStamp
    string processName
}

type RunningProcessResponse {
    []RunningProcess dataTable
}

func QueryOneProcess(QueryOneRequest) returns (QueryOneResponse, error)

func QueryRunningProcesses() returns (RunningProcessResponse, error)

```

### Client - Server Communication

* The server will store and maintain the running pids and timestamps in memory in a concurrent map, since pids are reusable and multiple instance of an executable can be executed with the same parameter, possibly at the same time, pids and timestamp should be unique enough to be identified with

* The Server will have to maintain the logs in the file system

* The Client will have a a queue of the the different requests 

    * start commands will insert new pids into store, deleting pids that are finished

    * the stop commands will execute stops and will delete pids from store when successful

    * the query commands grab contents of logs and running status for the specified pid



### Implementation Details 

* Using GRPC so the above methods and types with be generated via protocol buffers (will be using libprotoc 3.11.4)

####  DataStore

* we can use sync.Map in Go to implement a set of structs to store {pid,startTimeStamp,processName}


#### Server 

* data store to pid, starting timestamp
    
    * can be implemented with a concurrent map

* function to execute commands

    * os/exec and syscall packages can take care of this

* file system management

    * can use log and writer and reader packages

* concurrent go routines to handle cmd executions starting and finishing 

* killing of pid can be implemented can be done through sigkill via syscall

#### Client

* flags package to parse use commands or have preset commands from a file or hardcoded

* Basic Queue data structure to execute commands in FIFO order

* Can use the same data store structure for server and client


### Authentication

* Using standard GRPC tsl/ssl encryption and authentication via certificates

* we will be using grpc with mutual tls

* certificates will be generated usign openssl and self-signed, though in a real production environment, certificates will be generated and signed 
by a valid certificate authority (CA). there is :

    * public domains:

        *  Let’s Encrypt - a free automatef open certificate authority got cert generation and distribution  

    * private domains:

        * vault to generate signing requests, centrify for renewels of certs