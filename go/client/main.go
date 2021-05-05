package main

import (
	"context"
	"fmt"
	pb "github.com/roodrigoroot/go-service/proto"
	"google.golang.org/grpc"
	
)




func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	if err != nil {
		panic("cannot connect with server " + err.Error())
	}

	serviceClient := pb.NewUserClient(conn)

	res, err := serviceClient.GetUser(context.Background(), &pb.Empty{})

	if err != nil {
		panic("wishlist is not created  " + err.Error())
	}

	fmt.Println(res)

}