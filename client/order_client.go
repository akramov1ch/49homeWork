package main

import (
    "context"
    "log"
    "time"

    "google.golang.org/grpc"
    pb "49HW/gen/order"
)

func createOrder(client pb.OrderServiceClient, userID string, productName string, quantity int32, price float64) {
    ctx, cancel := context.WithTimeout(context.Background(), time.Second)
    defer cancel()

    req := &pb.CreateOrderRequest{UserId: userID, ProductName: productName, Quantity: quantity, Price: price}
    res, err := client.CreateOrder(ctx, req)
    if err != nil {
        log.Fatalf("could not create order: %v", err)
    }

    log.Printf("Created Order: %s", res.Order.Id)
}

func main() {
    conn, err := grpc.Dial("localhost:50052", grpc.WithInsecure())
    if err != nil {
        log.Fatalf("did not connect: %v", err)
    }
    defer conn.Close()

    client := pb.NewOrderServiceClient(conn)

    createOrder(client, "1", "Laptop", 1, 1000.0)
    createOrder(client, "2", "Phone", 2, 500.0)
}
