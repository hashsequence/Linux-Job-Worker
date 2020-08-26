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

