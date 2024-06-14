package main

import (
    "context"
    "fmt"
    "sync"

    pb "49HW/gen/user"
    "google.golang.org/grpc"
    "google.golang.org/grpc/codes"
    "google.golang.org/grpc/status"
)

type userServer struct {
    pb.UnimplementedUserServiceServer
    mu    sync.Mutex
    users map[string]*pb.User
}

func newUserServer() *userServer {
    return &userServer{
        users: make(map[string]*pb.User),
    }
}

func (s *userServer) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.UserResponse, error) {
    s.mu.Lock()
    defer s.mu.Unlock()

    id := fmt.Sprintf("%d", len(s.users)+1)
    user := &pb.User{
        Id:    id,
        Name:  req.Name,
        Email: req.Email,
    }
    s.users[id] = user

    return &pb.UserResponse{User: user}, nil
}

func (s *userServer) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.UserResponse, error) {
    s.mu.Lock()
    defer s.mu.Unlock()

    user, exists := s.users[req.Id]
    if !exists {
        return nil, status.Errorf(codes.NotFound, "user not found")
    }

    return &pb.UserResponse{User: user}, nil
}

func (s *userServer) ListUsers(ctx context.Context, req *pb.Empty) (*pb.UserListResponse, error) {
    s.mu.Lock()
    defer s.mu.Unlock()

    users := make([]*pb.User, 0, len(s.users))
    for _, user := range s.users {
        users = append(users, user)
    }

    return &pb.UserListResponse{Users: users}, nil
}
