package main

import (
	"log"
	"net"

	orderpb "49HW/gen/order"
	pb "49HW/gen/user"

	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterUserServiceServer(s, newUserServer())

	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	lisOrder, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	sOrder := grpc.NewServer()
	orderpb.RegisterOrderServiceServer(sOrder, newOrderServer())

	if err := sOrder.Serve(lisOrder); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
