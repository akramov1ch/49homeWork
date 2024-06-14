package main

import (
	"context"
	"fmt"
	"sync"

	pb "49HW/gen/order"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type orderServer struct {
	pb.UnimplementedOrderServiceServer
	mu     sync.Mutex
	orders map[string]*pb.Order
}

func newOrderServer() *orderServer {
	return &orderServer{
		orders: make(map[string]*pb.Order),
	}
}

func (s *orderServer) CreateOrder(ctx context.Context, req *pb.CreateOrderRequest) (*pb.OrderResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	id := fmt.Sprintf("%d", len(s.orders)+1)
	order := &pb.Order{
		Id:          id,
		UserId:      req.UserId,
		ProductName: req.ProductName,
		Quantity:    req.Quantity,
	}
	s.orders[id] = order

	return &pb.OrderResponse{Order: order}, nil
}

func (s *orderServer) GetOrder(ctx context.Context, req *pb.GetOrderRequest) (*pb.OrderResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	order, exists := s.orders[req.Id]
	if !exists {
		return nil, status.Errorf(codes.NotFound, "order not found")
	}

	return &pb.OrderResponse{Order: order}, nil
}

func (s *orderServer) ListOrders(ctx context.Context, req *pb.Empty) (*pb.OrderListResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	orders := make([]*pb.Order, 0, len(s.orders))
	for _, order := range s.orders {
		orders = append(orders, order)
	}

	return &pb.OrderListResponse{Orders: orders}, nil
}
