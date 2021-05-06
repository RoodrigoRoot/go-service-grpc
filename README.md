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

3. Enter the container:
`docker exec -ti <name_container> bash`

4. Run client:
`go run client/main.go`

5. Response:
`name:"Rodrigo"  age:"25"  email:"francisco@miflink.com"  phone:"741"`


### Micro with gRPC
`cd micro`


Under construction ....









