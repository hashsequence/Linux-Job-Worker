package main

import (
	"fmt"
	clientLib "github.com/hashsequence/Linux-Job-Worker/pkg/client"
	linuxJobWorkerPb "github.com/hashsequence/Linux-Job-Worker/pkg/pb"
	"log"
	"context"
	"time"
)

var (
	crt = "./ssl/client-cert.pem"
	key = "./ssl/client-key.pem"
	ca = "./ssl/ca-cert.pem"
	addr = "localhost:50051"

)

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r)
        }
		fmt.Println("Client Closing")
	}()
	client, err := clientLib.CreateClient(ca, crt, key, addr)
	if err != nil {
        log.Fatalf("could not create client stream %v: %v", addr, err)
	}
	//time.Sleep(5*time.Second)
	startReq := &linuxJobWorkerPb.StartRequest{
		Command : "./testProgram1_exe",
		Args : []string{"30"},
	}
	startResp, err := client.Start(context.Background(),startReq)
	fmt.Printf("startResp: \n%v\n err: %v \n",startResp, err)

	time.Sleep(3 * time.Second)
	queryOneReq := &linuxJobWorkerPb.QueryOneProcessRequest {
			Uuid : startResp.GetUuid(),
	}
	queryOneResp, err := client.QueryOneProcess(context.Background(), queryOneReq)
	fmt.Printf("queryOneResp: \n%v\n err: %v \n",queryOneResp, err)

	time.Sleep(6 * time.Second)
	queryOneResp, err = client.QueryOneProcess(context.Background(), queryOneReq)
	fmt.Printf("queryOneResp: \n%v\n err: %v \n",queryOneResp, err)

	time.Sleep(9 * time.Second)
	stopReq := &linuxJobWorkerPb.StopRequest{
		Uuid : startResp.GetUuid(),
	}
	stopResp, err := client.Stop(context.Background(),stopReq)
	fmt.Printf("stopResp: \n%v\n err: %v \n",stopResp, err)
	time.Sleep(1*time.Second)
	stopReq = &linuxJobWorkerPb.StopRequest{
		Uuid : startResp.GetUuid(),
	}
	stopResp, err = client.Stop(context.Background(),stopReq)
	fmt.Printf("stopResp: \n%v\n err: %v \n",stopResp, err)
	queryOneResp, err = client.QueryOneProcess(context.Background(), queryOneReq)
	fmt.Printf("queryOneResp: \n%v\n err: %v \n",queryOneResp, err)


	startReq = &linuxJobWorkerPb.StartRequest{
		Command : "./testProgram1_exe",
		Args : []string{"7"},
	}
	startResp, err = client.Start(context.Background(),startReq)
	fmt.Printf("startResp: \n%v\n err: %v \n",startResp, err)

	time.Sleep(2 * time.Second)
	queryRunProcReq := &linuxJobWorkerPb.QueryRunningProcessesRequest	{}
	queryRunProcResp, err := client.QueryRunningProcesses(context.Background(),queryRunProcReq)
	fmt.Printf("queryRunProcResp:\n")
	clientLib.PrintLinuxWorkerJobPbProcessTable(queryRunProcResp.GetProcessTable())

	//time.Sleep(7 * time.Second)
	//stopReq = &linuxJobWorkerPb.StopRequest{
	//	Uuid : startResp.GetUuid(),
	//}
	//stopResp, err = client.Stop(context.Background(),stopReq)
	//fmt.Printf("stopResp: \n%v\n err: %v \n",stopResp, err)


	//startReq = &linuxJobWorkerPb.StartRequest{
	//	Command : "head",
	//	Args : []string{"-c 1000000" ,"/dev/urandom"},
	//}
	//startResp, _ = client.Start(context.Background(),startReq)
	//fmt.Printf("startResp: \n%v\n",startResp)
//
	//time.Sleep(1 * time.Second)
	//queryOneReq := &linuxJobWorkerPb.QueryOneProcessRequest {
	//	Uuid : startResp.GetUuid(),
	//}
	//queryOneResp, err := client.QueryOneProcess(context.Background(), queryOneReq)
	//fmt.Println("------------------------------------------------------------------")
	//fmt.Printf("queryOneResp: \n%v\n err: %v \n",queryOneResp, err)
	//
	//time.Sleep(1 * time.Second)
	//queryOneReq = &linuxJobWorkerPb.QueryOneProcessRequest {
	//	Uuid : startResp.GetUuid(),
	//}
	//queryOneResp, err = client.QueryOneProcess(context.Background(), queryOneReq)
	//fmt.Println("------------------------------------------------------------------")
	//fmt.Printf("queryOneResp: \n%v\n err: %v \n",queryOneResp, err)
//
	//stopReq = &linuxJobWorkerPb.StopRequest{
	//	Uuid : startResp.GetUuid(),
	//}
	//stopResp, err = client.Stop(context.Background(),stopReq)
	//fmt.Printf("stopResp: \n%v\n err: %v \n",stopResp, err)

}