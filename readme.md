# gRPC quick start 
gRPC simple client and server in docker container 

### structure 
- `./pkg/grpc-client` stands for  grpc client  
- `./pkg/grpc-listener` stands for grpc server
- `./pkg/grpc-servers` stands for grpc services used by grpc server
- `./pkg/logs` contains proto files and generated files 
- `./cmd` application entry 

### Quick start  
`docker-compose up -d`
starts a grpc server on port 8000

### Run client 
`docker exec -it go-container sh` then `go run ./cmd/client` it will call client 

### Generate proto
`docker exec -it go-container sh` then `cd to-dir && protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative ./example.proto` 
