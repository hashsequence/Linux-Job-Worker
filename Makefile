test:
	go test -v ./...

clean: 
	rm -rf logs client_exe server_exe testProgram1_exe

server:
	go build -o ./server_exe ./cmd/server/ 

killServer:
	killall server_exe

build:
	make clean 
	go build -o ./server_exe ./cmd/server/ 
	go build -o ./client_exe ./cmd/client/ 
	go build -o ./testProgram1_exe ./cmd/testProgram1/ 

command:
	./server_exe &
	./client_exe &




	