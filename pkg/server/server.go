package server

import (
	"fmt"
	ds "github.com/hashsequence/Linux-Job-Worker/pkg/dataStore"
	linuxJobWorkerPb "github.com/hashsequence/Linux-Job-Worker/pkg/pb"
	utils "github.com/hashsequence/Linux-Job-Worker/pkg/utils"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	//"os/signal"
	"context"
	"crypto/tls"
	"crypto/x509"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/status"
	"net"
	"syscall"
)

type LinuxJobWorkerServer struct {
	dataStore *ds.DataStore
}

func NewLinuxJobWorkerServer() *LinuxJobWorkerServer {
	dataStore, _ := ds.NewDataStore()
	return &LinuxJobWorkerServer{
		dataStore: dataStore,
	}
}

func (s *LinuxJobWorkerServer) Serve(caCrt, serverCrt, serverKey, addr string) error {

	// Load the certificates from disk
	certificate, err := tls.LoadX509KeyPair(serverCrt, serverKey)
	if err != nil {
		return fmt.Errorf("could not load server key pair: %v", err)
	}

	// Create a certificate pool from the certificate authority
	certPool := x509.NewCertPool()
	ca, err := ioutil.ReadFile(caCrt)
	if err != nil {
		return fmt.Errorf("could not read ca certificate: %v", err)
	}

	// Append the client certificates from the CA
	if ok := certPool.AppendCertsFromPEM(ca); !ok {
		return fmt.Errorf("failed to append ca certs")
	}

	// Create the channel to listen on
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		return fmt.Errorf("could not list on %v: %v", addr, err)
	}

	// Create the TLS credentials
	creds := credentials.NewTLS(&tls.Config{
		ClientAuth: tls.RequireAndVerifyClientCert,
		//InsecureSkipVerify: true,
		Certificates: []tls.Certificate{certificate},
		ClientCAs:    certPool,
	})

	// Create the gRPC server with the credentials
	srv := grpc.NewServer(grpc.Creds(creds))

	// Register the handler object
	linuxJobWorkerPb.RegisterLinuxJobWorkerServiceServer(srv, s)

	// Serve and Listen
	if err := srv.Serve(lis); err != nil {
		return fmt.Errorf("grpc serve error: %v", err)
	}

	return nil
}

func (this *LinuxJobWorkerServer) Start(ctx context.Context, req *linuxJobWorkerPb.StartRequest) (*linuxJobWorkerPb.StartResponse, error) {
	fmt.Println("making start request")
	var cmd *exec.Cmd
	var err error
	pid := 0
	command := req.GetCommand()
	args := req.GetArgs()
	uuid := utils.GetNewUUID()
	cmd = exec.Command(command, args...)
	cmd.Env = req.GetEnv()
	cmd.Dir = req.GetDir()
	//sets the child process to have the same pgid as the parent
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Setpgid: true,
	}
	cmdArr := append([]string{command}, args...)
	this.dataStore.Add(uuid, cmdArr...)
	startTimeStamp, ok := this.dataStore.GetStartTimeStamp(uuid)
	if !ok {
		err = fmt.Errorf("GetStartTimeStamp failed when sending StartResponse with uuid: %v with pid: %v with status %v", uuid, pid, ok)
		return nil, status.Errorf(codes.NotFound, err.Error())
	}

	//get writers to stdouterr.log and stderr.log
	stdoutPath, _ := this.dataStore.GetStdoutPath(uuid)
	stderrPath, _ := this.dataStore.GetStderrPath(uuid)
	stdoutFile, err := os.OpenFile(stdoutPath, os.O_APPEND|os.O_WRONLY, 0644)
	stderrFile, err := os.OpenFile(stderrPath, os.O_APPEND|os.O_WRONLY, 0644)
	stdoutWriter := utils.NewLogWriter(stdoutFile)
	stderrWriter := utils.NewLogWriter(stderrFile)
	cmd.Stdout = io.MultiWriter(stdoutWriter)
	cmd.Stderr = io.MultiWriter(stderrWriter)

	startErr := cmd.Start()
	if cmd == nil || cmd.Process == nil || startErr != nil {
		if err = this.dataStore.UpdateFailedProcessDidNotStart(uuid); err != nil {
			return nil, status.Errorf(codes.FailedPrecondition, fmt.Errorf("%v | %v", startErr, err).Error())
		} else {
			return nil, status.Errorf(codes.FailedPrecondition, fmt.Errorf("cmd.Process is null due to error %v", startErr).Error())
		}
	} else {
		pid = cmd.Process.Pid
		if updateWithPidErr := this.dataStore.UpdateWithPid(uuid, pid, cmd); err != nil {
			err = fmt.Errorf("%v | updateWithPid failed updating %v with pid: %v", updateWithPidErr, uuid, pid)
			return nil, status.Errorf(codes.FailedPrecondition, err.Error())
		}
	}

	wait := func() {
		err = cmd.Wait()
		stdoutFile.Close()
		stderrFile.Close()
		var exitCode int

		if err == nil {
			exitCode = cmd.ProcessState.ExitCode()
			fmt.Printf("%v exited at Completion with exitCode: %v\n", uuid, exitCode)
		} else {
			fmt.Printf("%v Did Not Exit at Completion", uuid)
			return
		}
		this.dataStore.UpdateFinishProcess(uuid, exitCode)
		fmt.Println("Finished Cmd: ", cmd)
	}

	go wait()

	return &linuxJobWorkerPb.StartResponse{
		Pid:            int32(pid),
		Uuid:           uuid,
		StartTimeStamp: startTimeStamp,
	}, nil
}

func (this *LinuxJobWorkerServer) Stop(ctx context.Context, req *linuxJobWorkerPb.StopRequest) (*linuxJobWorkerPb.StopResponse, error) {
	var cmd *exec.Cmd
	var stdoutBuf []byte
	var stderrBuf []byte
	isKilled := false
	var exitCode int
	var endTimeStamp string
	uuid := req.GetUuid()
	fmt.Printf("Stopping %s\n", uuid)

	cmd, ok := this.dataStore.GetCmd(uuid)
	if !ok {
		return nil, status.Errorf(codes.FailedPrecondition, "Unable to get Cmd")
	}
	if cmd == nil || cmd.Process == nil {
		return nil, status.Errorf(codes.FailedPrecondition, "cmd is nil")
	}
	//if procErr := cmd.Process.Kill(); procErr != nil {
	//	return nil, status.Errorf(codes.FailedPrecondition, fmt.Sprintf("failed to kill process: %v", procErr))
	//}
	if pgid, pgidErr := syscall.Getpgid(cmd.Process.Pid); pgidErr == nil {
		fmt.Printf("killing processes and descendant process for pgid: %v\n",pgid)
		if sysKillErr := syscall.Kill(-pgid, syscall.SIGKILL); sysKillErr != nil {
			return nil, status.Errorf(codes.FailedPrecondition, fmt.Sprintf("failed to kill process, child and desecendant processes: %v", sysKillErr))
		}
	} else {
		return nil, status.Errorf(codes.FailedPrecondition, fmt.Sprintf("failed to get pgid of pid to kill process, child and desecendant processes: %v", pgidErr))
	}
	
	isKilled = true
	exitCode = cmd.ProcessState.ExitCode()


	if updateFinishProcessErr := this.dataStore.UpdateFinishProcess(uuid, exitCode); updateFinishProcessErr != nil {
		return nil, status.Errorf(codes.FailedPrecondition, updateFinishProcessErr.Error())
	}
	stdoutPath, ok := this.dataStore.GetStdoutPath(uuid)
	if !ok {
		return nil, status.Errorf(codes.NotFound, fmt.Sprintf("Unable to find stdout.log for %v", req.GetUuid()))
	}
	stderrPath, ok := this.dataStore.GetStderrPath(uuid)
	if !ok {
		return nil, status.Errorf(codes.NotFound, fmt.Sprintf("Unable to find stderr.log for %v", req.GetUuid()))
	}
	stdoutBuf, stdoutBufErr := ioutil.ReadFile(stdoutPath)
	if stdoutBufErr != nil {
		return nil, status.Errorf(codes.FailedPrecondition, fmt.Sprintf("%v", stdoutBufErr))
	}
	stderrBuf, stderrBufErr := ioutil.ReadFile(stderrPath)
	if stderrBufErr != nil {
		return nil, status.Errorf(codes.FailedPrecondition, fmt.Sprintf("%v", stderrBufErr))
	}
	if endTimeStamp, ok = this.dataStore.GetEndTimeStamp(uuid); !ok {

		return nil, status.Errorf(codes.FailedPrecondition, fmt.Sprintf("Unable to find endTimeStamp for %v", uuid))
	}

	return &linuxJobWorkerPb.StopResponse{
		Stdout:       stdoutBuf,
		Stderr:       stderrBuf,
		EndTimeStamp: endTimeStamp,
		IsKilled:     isKilled,
		ExitCode:     int32(exitCode),
	}, nil
}

func (this *LinuxJobWorkerServer) QueryOneProcess(ctx context.Context, req *linuxJobWorkerPb.QueryOneProcessRequest) (*linuxJobWorkerPb.QueryOneProcessResponse, error) {
	var stdoutBuf []byte
	var stderrBuf []byte
	var stdoutBufErr error
	var stderrBufErr error
	uuid := req.GetUuid()

	//get stdouPath from dataStore using uuid
	if stdoutPath, ok := this.dataStore.GetStdoutPath(uuid); !ok {
		return nil, status.Errorf(codes.NotFound, fmt.Sprintf("Unable to find stdout.log for %v", uuid))
	} else {
		//get contents of stdout.log into buffer
		stdoutBuf, stdoutBufErr = ioutil.ReadFile(stdoutPath)
		if stdoutBufErr != nil {
			return nil, status.Errorf(codes.FailedPrecondition, fmt.Sprintf("%v", stdoutBufErr))
		}
	}
	//get stdouterrPath from dataStore using uuid
	if stderrPath, ok := this.dataStore.GetStderrPath(uuid); !ok {
		return nil, status.Errorf(codes.NotFound, fmt.Sprintf("Unable to find stderr.log for %v", uuid))
	} else {
		//get contents of stderr.Log into buffer
		stderrBuf, stderrBufErr = ioutil.ReadFile(stderrPath)
		if stderrBufErr != nil {
			return nil, status.Errorf(codes.FailedPrecondition, fmt.Sprintf("%v", stderrBufErr))
		}
	}
	//get resp processInfo
	respProcessInfo, ok := this.dataStore.GetRespProcessInfo(uuid)
	if !ok {
		return nil, status.Errorf(codes.NotFound, fmt.Sprintf("unable to get processInfo for %v", uuid))
	}

	return &linuxJobWorkerPb.QueryOneProcessResponse{
		ProcInfo: respProcessInfo,
		Stdout:   stdoutBuf,
		Stderr:   stderrBuf,
	}, nil
}

func (this *LinuxJobWorkerServer) QueryRunningProcesses(ctx context.Context, req *linuxJobWorkerPb.QueryRunningProcessesRequest) (*linuxJobWorkerPb.QueryRunningProcessesResponse, error) {
	processTable, ok := this.dataStore.GetProcessTable()
	var err error
	if !ok {
		this.dataStore, err = ds.NewDataStore()
		return nil, status.Errorf(codes.FailedPrecondition, fmt.Sprintf("processTable has bad data, will clear out all data %v", err.Error()))
	}
	return &linuxJobWorkerPb.QueryRunningProcessesResponse{
		ProcessTable: processTable,
	}, nil
}
