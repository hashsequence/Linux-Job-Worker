package main

import (
	"fmt"
	server "github.com/hashsequence/Linux-Job-Worker/pkg/server"
)

const (
	crt = "./ssl/server-cert.pem"
	key = "./ssl/server-key.pem"
	caCrt = "./ssl/ca-cert.pem"
	addr = "localhost:50051"
)

func main() {
	
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r)
        }
		fmt.Println("Server Closing")
	}()

	fmt.Println("Starting Server")
	s := server.NewLinuxJobWorkerServer()
	if err := s.Serve(caCrt, crt, key, addr); err != nil {
		panic(err)
	}

}