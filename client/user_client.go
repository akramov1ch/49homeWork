package main

import (
    "context"
    "log"
    "time"

    "google.golang.org/grpc"
    pb "49HW/gen/user"
)

func createUser(client pb.UserServiceClient, name string, email string) {
    ctx, cancel := context.WithTimeout(context.Background(), time.Second)
    defer cancel()

    req := &pb.CreateUserRequest{Name: name, Email: email}
    res, err := client.CreateUser(ctx, req)
    if err != nil {
        log.Fatalf("could not create user: %v", err)
    }

    log.Printf("Created User: %s", res.User.Id)
}

func main() {
    conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
    if err != nil {
        log.Fatalf("did not connect: %v", err)
    }
    defer conn.Close()

    client := pb.NewUserServiceClient(conn)

    createUser(client, "Alice", "alice@example.com")
    createUser(client, "Bob", "bob@example.com")
}
