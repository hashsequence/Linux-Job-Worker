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
the file system of the linux worker. The logs will be generated at start time with the foldername \<pid\>_\<startingTimestamp\>
The client can query the server to see what is running to determine what can be killed and a query a list of jobs that were executed,
or the client can store response data from server.

#### Scale
The scope of this project would only deal with a single linux worker server interfacing with multiple clients

### API Design

* The client will have an Command Api that takes in three types of request, Start, Stop, and Query.

* The client can then execute the commands over the server via Execute APIs'

#### Start

* The Start command is called with a StartRequest that has the client's command and required arguments and optional env, dir params

* A uuid (universal unique identification) will be generated and a folder called \<uuid\>-\<startingtimestamp\> will be created, two logs called PID-<pid>-stdout.log and PID-<pid>-stderr.log will be created, 

* the start command will execute the job and return with the uuid, pid, startingtimestamp, if it fails to execute then a log called FAILED.log will be created to indcate that the job failed to execute

* goroutines should manage running processes in the background (outputing into logs, updating dataStore)

* when the job is done it will generate a log called END-\<endtimestamp\>

```
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
    //univeral unique identifier that tags each unique request made to server
    string uuid
    //starting time of start request
    string startingTimeStamp

}

func ExecuteStart(StartRequest) returns(StartResponse)
```

#### Stop

* The User should be able to stop the request based on the pid

* When stopped the pid should be killed 

* A response will be sent to client indicating process have been stopped along with the contents of the log

* job should be marked as completed with the exit code in the dataStore

```

type LogOutput {
    //contents of log
    []byte contents 
}

type StopRequest {
    int pid
}

type StopResponse {
    LogOutput stdout
    LogOutput stderr
    bool isKilled
}

func ExecuteStop(StopRequest) returns(StopResponse)

```

#### Query 


* There will be only two types of Query

    * QueryOneProcess:

        * return the logs of a job using a valid pid (if it was started) or uuid, along with ProcessInfo


    * QueryRunningProcesses:

        * get a list of job's executed and processInfo for the jobs

```
type ProcessInfo {
    int pid 
    string startingTimeStamp 
    string endTimeStamp
    string processName 
    string uuid
    bool isRunning 
    int exitCode
}

type QueryOneProcessRequest {
    int pid 
    string startingTimeStamp
    string uuid
}

type QueryOneProcessResponse {
    processInfo procInfo
    LogOutput stdout 
    LogOutput stderr 
}

type QueryRunningProcessesRequest {
    //will be empty since the server just needs to verify its a QueryRunningProcessesRequest
}
    
type QueryRunningProcessesResponse {
     ProcessInfo[] processTable 
}

func ExecuteQueryOneProcess(QueryOneProcessRequest) returns(QueryOneProcessResponse)
func ExecuteQueryRunningProcesses(QueryRunningProcessesRequest) returns(QueryRunningProcessesResponse)

```


### Implementation Overview 

* Using GRPC so the above methods and types with be generated via protocol buffers (will be using libprotoc 3.11.4)

####  DataStore

* we can use Map in Go to implement a set of structs to store process info and use sync.mutex to handle concurrent transactions, the key to the map will be \<uuid\>-\<startTimeStamp\>
     ```go
        type ProcessInfo struct {
            pid int
            string startTimeStamp
            string endTimeStamp
            string processName
            string uuid
            string logPath
            bool isRunning
            int exitCode
        }
        
        type dataStore map[string]ProcessInfo
        //methods to manage access to dataStore
        //...
     ```


#### Server 

* use a DataStore Structure to function as a in-memory database

* function to execute commands 

    * os/exec and syscall packages can take care of this

* logfile system management

    * can use log, writer and reader packages

    * must manage concurrency issues with read and write
    
    * will be similar to how linux implements it using the /proc path:
        ```
        $ ls -al  /proc | head -n 10
        total 4
        dr-xr-xr-x 312 root             root                           0 Aug 24 11:41 .
        drwxr-xr-x  23 root             root                        4096 Aug 17 16:01 ..
        dr-xr-xr-x   9 root             root                           0 Aug 24 11:41 1
        dr-xr-xr-x   9 root             root                           0 Aug 26 16:40 10
        dr-xr-xr-x   9 avwong13         avwong13                       0 Aug 26 16:40 1003
        dr-xr-xr-x   9 root             root                           0 Aug 26 17:08 1004
        dr-xr-xr-x   9 avwong13         avwong13                       0 Aug 25 23:08 10202
        dr-xr-xr-x   9 avwong13         avwong13                       0 Aug 25 23:09 10256
        dr-xr-xr-x   9 avwong13         avwong13                       0 Aug 24 13:25 10485
        ```

* The server will only store running pids in memory and have the pass logs available in their respective \<uuid\>_\<startingTimeStamp\> folder

* concurrent go routines to handle cmd executions starting and finishing 

* killing of pid can be implemented through sigkill via os.Process.Signal()

* Data for query for all processes ever executed by linux worker server should be generated by iterating through file system and extracting data from
logs and filepath using path packages

* I~~mplement simple logging for server, as in the server stdout and stderr will output into a log in the server's filesystem~~
    
    * Remarks: don't worry about the logs of the server itself, they can go to stdout/stderr if needed, something like systemd can redirect those to a file.

#### Client

* client is responsible for constructing the commands to be passed to server via the Execute APIs'

* client is responsible for interpreting response from Execute APIs'

* client should remember the information from the response of the requests from the server


### Authentication

* Using standard GRPC tsl/ssl encryption and authentication via certificates

* we will be using grpc with mutual tls

* certificates will be generated using openssl and self-signed, though in a real production environment, certificates will be generated and signed 
by a valid certificate authority (CA). there is :

    * public domains:

        *  Let’s Encrypt - a free automated open certificate authority got cert generation and distribution  

    * private domains:

        * vault to generate signing requests, centrify for renewels of certificates

### Contentious Issues

* Should the client remember commands they executed and remember the pids, starting timestamp, process command?

    * yes, but not across restarts

* Will logs that are too old be deleted by server, or should the logs just stay in the file system storage, should this be within the scope of the project?
    
    * log cleanup is not in scope

* For the output of logs, should the should the contents of the log be loaded into a string array in the response message, or should I leave it as string?
    
    * will be using []byte

### Development Timeline

(Will be Updated througout development)

* Setup local development environment 

    * working on Ubuntu 16.04 
    
    * install protoc (will be using proto3)
        
        * current version
        ```
        $ protoc --version
        libprotoc 3.11.4

        ```
    * install golang 
        
        * current version
        ```
        $ go version
        go version go1.15 linux/amd64
        ```
    * install openssl (to generate certificates and keys)

        * current version
        ```
        $ openssl version
        OpenSSL 1.0.2g  1 Mar 2016
        ```

* Setup project directory

    * setup go modules (dependency management)

* Write Protocol Buffers for grpc and generate go package

* Implement Authentication and Encryption for grpc in go

* Implement Start 

* Implement Stop

* Implement Query 

* Write Tests

