# go-service-grpc
## This reposiroy is a example of Go with gRPC.

## Installation
`git clone https://github.com/RoodrigoRoot/go-service-grpc.git`

### Go with gRPC
`cd go`

#### Installation

1. Create the container:
`docker build -t go-grpc .`

2. Run container:
`docker run -p 50051:50051 go-grpc`




### To run Stress Test (Python-gRPC) 
We need have ghz

and execute:

`ghz  -c 100 -n 1000 --insecure  --proto proto/users.proto --call users.User.GetUser  localhost:50051`





### Micro with gRPC 
`cd micro`

#### Installation

1. Create the container:
`docker build -t micro-grpc .`

2. Run container:
`docker run -p 8081:8081 go-grpc`


### To run Stress Test (Go-Micro)

execute:

`ghz  -c 100 -n 1000 --insecure  --proto proto/greeter.proto --call greeter.Greeter.GetUser  localhost:8081`







