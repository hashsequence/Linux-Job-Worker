FROM golang

RUN mkdir ./Linux-Job-Worker

WORKDIR ./Linux-Job-Worker

COPY . .

RUN go build -o ./server_exe ./cmd/server/ 

ENTRYPOINT ["./server_exe"]

EXPOSE 50051