package main

import (
	"context"
	"fmt"
	pb "github.com/roodrigoroot/go-service/proto"
	"google.golang.org/grpc"
	"net"
)

type server struct {
	pb.UnimplementedUserServer
}


func (s *server) GetUser(ctx context.Context, req *pb.Empty) (*pb.UserInfo, error) {
	fmt.Println("Starting.......") 
	return &pb.UserInfo{Name:"Rodrigo", Email:"francisco@miflink.com", Age:"25", Phone:"741"}, nil
}

func main() {
	listner, err := net.Listen("tcp", ":50051")

	if err != nil {
		panic("cannot create tcp connection" + err.Error())
	}
	fmt.Println("Running on port 50051")
	serv := grpc.NewServer()
	pb.RegisterUserServer(serv, &server{})
	if err = serv.Serve(listner); err != nil {
		panic("cannot initialize the server" + err.Error())
	}

}