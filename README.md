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
the file system of the linux worker. The logs will be generated at start time with the foldername \<uuid\>-\<startTimeStamp\>
The client can query the server to see what is running to determine what can be killed and a query a list of jobs that were executed,
or the client can store response data from server.

#### Scale
The scope of this project would only deal with a single linux worker server interfacing with multiple clients

### API Design

* The client will have an Command Api that takes in three types of request, Start, Stop, and Query

* The client can then execute the commands over the server via Execute APIs'

#### Start

* The Start command is called with a StartRequest that has the client's command and required arguments and optional env, dir params

* A uuid (universal unique identification) will be generated and a folder called START-\<startTimeStamp\>will be created, two logs called stdout.log and stderr.log will be created

* the start command will execute the job and return with the uuid, pid, startTimeStamp, if it fails the process table will be correspondingly updated

* goroutines should manage running processes in the background (outputing into logs, updating dataStore)

* when the job is done the process table will be updated

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
    string status
    int statusCode
    

}

func Start(StartRequest) returns(StartResponse)
```

#### Stop

* The User should be able to stop the request based on the pid

* When stopped the pid should be killed 

* A response will be sent to client indicating process have been stopped along with the contents of the log

* job should be marked as completed with the exit code in the dataStore

```
type StopRequest {
    string uuid
}

type StopResponse {
    []byte stdout
    []byte stderr
    bool isKilled
    string endTimeStamp
    string status
    int statusCode
}

func Stop(StopRequest) returns(StopResponse)

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
    string startTimeStamp 
    string endTimeStamp
    string processName 
    bool isRunning 
    int exitCode
    string status
}

type QueryOneProcessRequest {
    string uuid
}

type QueryOneProcessResponse {
    processInfo procInfo
    []byte stdout 
    []byte stderr 
    string status
    int statusCode
}

type QueryRunningProcessesRequest {
    //will be empty since the server just needs to verify its a QueryRunningProcessesRequest
}
    
type QueryRunningProcessesResponse {
     ProcessInfo[] processTable 
     string status
     int statusCode
}

func QueryOneProcess(QueryOneProcessRequest) returns(QueryOneProcessResponse)
func QueryRunningProcesses(QueryRunningProcessesRequest) returns(QueryRunningProcessesResponse)

```


### Implementation Overview 

* Using GRPC so the above methods and types with be generated via protocol buffers (will be using libprotoc 3.11.4)

* the rpc APIs' will be all unary to keep the client - server communication simple

####  DataStore

* we can use Map in Go to implement a set of structs to store process info and use sync.mutex to handle concurrent transactions, the key to the map will be \<uuid\>
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

* The server will store process info of jobs that were requeste in memory in the dataStore and have logs for each request named \<uuid\>-\<startTimeStamp\> folder

* concurrent go routines to handle cmd executions starting and finishing 

* killing of pid can be implemented via os.Process.Signal() by sending a kill signal to process

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

* client is responsible for constructing the commands to be passed to server via the Execute APIs'

* client is responsible for interpreting response from Execute APIs'

* client should remember the information from the response of the requests from the server


### Authentication

* Using standard GRPC tsl/ssl encryption and authentication via certificates

* we will be using grpc with mutual tls

* certificates will be generated using openssl and self-signed to keep things simple for this project

#### Mutual TLS 

* since the client is basically sending commands to the server, both the server and client must know they are indeed safe and valid. Basically we are trying to solve the problem of encrypting messages between client and server and the client and server must know that they are indeed client and server, but how do we solve this?

* here's the basic algorithm to verify:
    
    * We have the client, the server, and and an authority

    * the client ask if the server is indeed the valid server

    * the server send its certificate 

    * the client verifys the certificate with the authority

    * the client then sends its certificate to the server after validating the server's certificate is goo

    * the server verifys the certificate with the authority

    * the server then verifies that the certficate is good 

    * now both the server and client knows that the connection is secure

    * the server and client have the respective keys (public key via in the certificates transfered) to encrypt the message, and the respective sides have they own
    private keys to decrypt the message

* does this algo work?

    * first the client and server can verify its validity by having a central authority to validate both certs

    * when the client/server sends over their certs over connection, if it was tampered with it will be rejected by the authority, and the intermediaries cant really use the information without knowing how to decrypt the message

    * a problem arises when the certificate authority is compromised and is not trustworthy, like someone had access to the ca.key, so how do we have a strong certificate authority, one simple way you could do is have that ca.crt and key in a box without any outside connection, and whenever you need new certs you physically go to the body and sign new certs for new servers/clients. this is hugely un-scalable so we can probably have the box accept secure request from the outside coming in to request certs to be signed. 

* Now how do we create all the needed certs and keys?

    * First we need a certificate authority(ca), the certificate authority provides the server and client its first certificate (ca.crt) it trusts
        
        * the certificate authority has a private key (ca.key) that is used to create valid certificates for client and server
        
        * the the server/client generates its own certificate accompanied with the private key to decrypt 

        * when a server or client asks to have the certificate validated, the server/client sends a request (.csr) to the ca to create a signed certificate

* In this project we will be making things simpler, since I am the one building the client and server I can make my own ca.crt + ca.key and have it sign
 the server.crt and client.crt, which is self signing. In production, I would probably have a legitimate certificate authority and a legitimate ca.crt

 * the format of the crt wil be [x509](https://en.wikipedia.org/wiki/X.509) 

 * will be using openssl to generate these certificates




### Contentious Issues

* Should the client remember commands they executed and remember the pids, starting timestamp, process command?

    * yes, but not across restarts

* Will logs that are too old be deleted by server, or should the logs just stay in the file system storage, should this be within the scope of the project?
    
    * log cleanup is not in scope

* For the output of logs, should the should the contents of the log be loaded into a string array in the response message, or should I leave it as string?
    
    * will be using []byte

* Memory management of dataStore: Since running jobs over time and having the server keeping track of new entries of jobs in the data Store gets expensive, should their be a process to delete entries or truncate the data Store over time? Perhaps clear the dataStore after a period of time passed? Maybe delete jobs that have been done for a period of time?

    * not within scope

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

* Write Protocol Buffers for grpc and generate go package [LinuxWorker.proto](https://github.com/hashsequence/Linux-Job-Worker/blob/feature/Design-Doc_Avery_V2/pb/LinuxWorker.proto)

* Implement Authentication and Encryption for grpc in go

* Implement Start 

* Implement Stop

* Implement Query 

* Write Tests

