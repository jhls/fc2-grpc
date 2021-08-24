
package main;

import (
	"net"
	"log"
	"google.golang.org/grpc"
)



func  main()  {
	
	lis, err := net.Listen("tcp","localhost:50051")


	if err != nil {
		log.Fatalf("Could not connect: %v",err)
	}

	grpcServer := grpc.NewServer()

	if err := grpcServer.Serve(lis); err != nil{
		log.Fatalf("Could not serve: %v",err)
	}

}