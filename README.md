# Linux-Job-Worker
Distributed Linux Job Worker

## Summary
A job worker service that provides an API to run arbitrary Linux processes.

## Requirements

### API
* Job worker will provide RPC API to start/stop/query status and get an output of a running job process. Any RPC mechanism that works for the task: GRPC and barebones SSL will be used to guarantee secure and reliable client-server communication as well as a simple but secure authentication and authorization mechanism.

### Client

* Client command will be able to connect to worker service and schedule several jobs

* Client will be able to query result of the job execution and fetch the logs

## Design Document

### Limitations And Scope

#### Data Management
Though the use of a database to store persistant data would be ideal, I will be instead storing the outputs and error outputs into logs stored on 
the file system of the linux worker. The logs will be generated at start time with the foldername \<uuid\>-\<startTimeStamp\>
The client can query the server to see what is running to determine what can be killed and a query a list of jobs that were executed,
or the client can store response data from server. I'll be using a self-implemented Data Store to store in memory data to manage jobs ran by server.

#### Scale
The scope of this project would only deal with a single linux worker server interfacing with multiple clients

### API Design

* The client will have an APIs' that handle the three types of request, Start, Stop, and Query

* The client can then execute the commands over the server via the APIs'

#### Start

* The Start command is called with a StartRequest that has the client's command and required arguments and optional env, dir params

* A uuid (universal unique identification) will be generated and a folder called START-\<startTimeStamp\>will be created, two logs called stdout.log and stderr.log will be created

* The start command will execute the job and return with the uuid, pid, startTimeStamp, if it fails the process table will be correspondingly updated

* Goroutines should manage running processes in the background (outputing into logs, updating dataStore)

* When the job is done the process table will be updated

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
    string startTimeStamp    
}

func Start(StartRequest) returns(StartResponse)
```

#### Stop

* The User should be able to stop the request based on the uuid

* When stopped the process should be killed 

* A response will be sent to client indicating process have been stopped along with the contents of the log

* Job should be marked as completed with the exit code in the dataStore

```
type StopRequest {
    string uuid
}

type StopResponse {
    []byte stdout
    []byte stderr
    bool isKilled
    string endTimeStamp
    int exitCode
}

func Stop(StopRequest) returns(StopResponse)

```

#### Query 


* There will be only two types of Query

    * QueryOneProcess:

        * Return the logs of a job using a valid uuid, along with ProcessInfo


    * QueryRunningProcesses:

        * Get a list of job's executed and processInfo for the jobs

```
type ProcessInfo {
    int pid 
    string startTimeStamp 
    string endTimeStamp
    string processName 
    bool isRunning 
    int exitCode
}

type QueryOneProcessRequest {
    string uuid
}

type QueryOneProcessResponse {
    processInfo procInfo
    []byte stdout 
    []byte stderr 
}

type QueryRunningProcessesRequest {
    //will be empty since the server just needs to verify its a QueryRunningProcessesRequest
}
    
type QueryRunningProcessesResponse {
     ProcessInfo[] processTable 
}

func QueryOneProcess(QueryOneProcessRequest) returns(QueryOneProcessResponse)
func QueryRunningProcesses(QueryRunningProcessesRequest) returns(QueryRunningProcessesResponse)

```

#### Error Handling

* Errors will be handled using the grpc error handling package in Go: [gprc/status](https://pkg.go.dev/google.golang.org/grpc/status?tab=doc)

* List of possible grpc Codes can be found here: [grpc error codes](https://pkg.go.dev/google.golang.org/grpc@v1.31.1/codes?tab=doc#Code)

* For Example in the Start API:
    ```go
        Start(context.Context, *StartRequest) (*StartResponse, error) {
            //...logic for start...


            //error return
            return nil, status.Errorf(codes.FailedPrecondition,
			"Start Process Did Not Start")
	}
        }
    ```
    The error is formated using the grpc/status pkg and returned accordingly in the return statement

### Implementation Overview 

* Using GRPC so the above methods and types with be generated via protocol buffers (will be using libprotoc 3.11.4)

* The rpc APIs' will be all unary to keep the client - server communication simple

####  DataStore

* We can use Map in Go to implement a set of structs to store process info and use sync.mutex to handle concurrent transactions, the key to the map will be \<uuid\>
     ```go

       type ProcessInfo struct {
		    pid int
		    startTimeStamp string
		    endTimeStamp string
		    processName string
		    logPath string
		    stdoutPath string
		    stderrPath string
		    isRunning bool
            exitCode int
            string status
            //possible other metadata fields to manage processes
	    }

        //key will be uuid
	    type ProcessTable map[string]*ProcessInfo

	    type DataStore struct {
		    sync.RWMutex
		    ProcessTable
		    logFolder string
	    }
        //methods to update, delete, add, access, create accompanying log folders


     ```
     * There are better sources for in memory datastore like redis, but for the scope of the project I will use the ones I will implement with Go
        
        * the tradeoff will be performance by using an simple datastore with mutexes rather than something more industry standard like redis


#### Server 

* Use a DataStore Structure to function as a in-memory database

* Function to execute commands 

    * os/exec and syscall packages can take care of this

* Logfile system management

    * Can use log, writer and reader packages

    * Must manage concurrency issues with read and write
    
    * Will be similar to how linux implements it using the /proc path:
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

* The server will store process info of jobs that were requeste in memory in the dataStore and have logs for each request named \<uuid\>-\<startTimeStamp\> folder

* Concurrent go routines to handle cmd executions starting and finishing 

* Killing of pid can be implemented via os.Process.Signal() by sending a kill signal to process

* Querys will use information from dataStore and logs from server

* Implement simple logging for server, as in the server stdout and stderr will output into a log in the server's filesystem
    
    * Remarks: don't worry about the logs of the server itself, they can go to stdout/stderr if needed, something like systemd can redirect those to a file.

* \<uuid\>-\<startTimeStamp\> can be implemented like this for example:

    ```go
    import(
        "fmt"
        guuid "github.com/google/uuid"
    )

    func main() {
        requestId := guuid.New().String()
        fmt.Println(requestId)
    }

    ```

#### Client

* Client is responsible for constructing the commands to be passed to server via the Execute APIs'

* Client is responsible for interpreting response from Execute APIs'

* Client should remember the information from the response of the requests from the server


### Authentication

* Using standard GRPC tsl/ssl encryption and authentication via certificates

* We will be using grpc with mutual tls

* Certificates will be generated using openssl and self-signed to keep things simple for this project

#### Mutual TLS 

* Since the client is basically sending commands to the server, both the server and client must know they are indeed safe and valid. Basically we are trying to solve the problem of encrypting messages between client and server and the client and server must know that they are indeed client and server, but how do we solve this?

* Here's the basic algorithm to verify:
    
    * We have the client, the server, and and an authority

    * The client ask if the server is indeed the valid server

    * The server send its certificate 

    * The client verifys the certificate with the authority

    * The client then sends its certificate to the server after validating the server's certificate is goo

    * The server verifys the certificate with the authority

    * The server then verifies that the certficate is good 

    * Now both the server and client knows that the connection is secure

    * The server and client have the respective keys (public key via in the certificates transfered) to encrypt the message, and the respective sides have they own
    private keys to decrypt the message

* Does this algo work?

    * First the client and server can verify its validity by having a central authority to validate both certs

    * When the client/server sends over their certs over connection, if it was tampered with it will be rejected by the authority, and the intermediaries cant really use the information without knowing how to decrypt the message

    * A problem arises when the certificate authority is compromised and is not trustworthy, like someone had access to the ca.key, so how do we have a strong certificate authority, one simple way you could do is have that ca.crt and key in a box without any outside connection, and whenever you need new certs you physically go to the body and sign new certs for new servers/clients. this is hugely un-scalable so we can probably have the box accept secure request from the outside coming in to request certs to be signed. 

* Now how do we create all the needed certs and keys?

    * First we need a certificate authority(ca), the certificate authority provides the server and client its first certificate (ca.crt) it trusts
        
        * The certificate authority has a private key (ca.key) that is used to create valid certificates for client and server
        
        * The the server/client generates its own certificate accompanied with the private key to decrypt 

        * When a server or client asks to have the certificate validated, the server/client sends a request (.csr) to the ca to create a signed certificate

* In this project we will be making things simpler, since I am the one building the client and server I can make my own ca.crt + ca.key and have it sign
 the server.crt and client.crt, which is self signing. In production, I would probably have a legitimate certificate authority and a legitimate ca.crt

 * The format of the crt wil be [x509](https://en.wikipedia.org/wiki/X.509) 

 * Will be using openssl to generate these certificates




### Contentious Issues

* Should the client remember commands they executed and remember the pids, starting timestamp, process command?

    * Yes, but not across restarts

* Will logs that are too old be deleted by server, or should the logs just stay in the file system storage, should this be within the scope of the project?
    
    * Log cleanup is not in scope

* For the output of logs, should the should the contents of the log be loaded into a string array in the response message, or should I leave it as string?
    
    * Will be using []byte

* Memory management of dataStore: Since running jobs over time and having the server keeping track of new entries of jobs in the data Store gets expensive, should their be a process to delete entries or truncate the data Store over time? Perhaps clear the dataStore after a period of time passed? Maybe delete jobs that have been done for a period of time?

    * Not within scope

### Development Timeline

* Setup local development environment 

    * Working on Ubuntu 16.04 
    
    * Install protoc (will be using proto3)
        
        * current version
        ```
        $ protoc --version
        libprotoc 3.11.4

        ```
    * Install golang 
        
        * current version
        ```
        $ go version
        go version go1.15 linux/amd64
        ```
    * Install openssl (to generate certificates and keys)

        * current version
        ```
        $ openssl version
        OpenSSL 1.0.2g  1 Mar 2016
        ```

* Setup project directory

    * Setup go modules (dependency management)

* Write Protocol Buffers for grpc and generate go package [LinuxWorker.proto](https://github.com/hashsequence/Linux-Job-Worker/blob/feature/Design-Doc_Avery_V2/pb/LinuxWorker.proto)

* Implement Authentication and Encryption for grpc in go - [x]

* Implement Start - [x]

* Implement Stop - [x]

* Implement Query - [x]

* Write Tests - [x]
