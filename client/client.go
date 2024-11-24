package main

import (
	"context"
	"fmt"
	"log"

	pb "test_train/protobuf" // Adjust according to your protobuf path

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// Set up connection to the server
	//conn, err := grpc.Dial("localhost:8111", grpc.WithInsecure())
	conn, err := grpc.NewClient("localhost:8111", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	defer conn.Close()

	// Create a new client
	client := pb.NewTrainServiceClient(conn)

	// Test PurchaseTicket
	user := &pb.User{
		FirstName: "John",
		LastName:  "Doe",
		Email:     "john.doe@example.com",
	}
	req := &pb.PurchaseTicketRequest{
		User: user,
		From: "London",
		To:   "France",
	}

	// Call PurchaseTicket
	receipt, err := client.PurchaseTicket(context.Background(), req)
	if err != nil {
		log.Fatalf("Error purchasing ticket: %v", err)
	}
	fmt.Printf("Ticket purchased: %+v\n", receipt)

	// Test GetReceipt
	receiptReq := &pb.UserRequest{
		Email: user.Email,
	}
	receiptResp, err := client.GetReceipt(context.Background(), receiptReq)
	if err != nil {
		log.Fatalf("Error getting receipt: %v", err)
	}
	fmt.Printf("Ticket receipt for %s: %+v\n", user.Email, receiptResp)

	// Test GetUsersBySection for Section A
	sectionReq := &pb.SectionRequest{Section: "A"}
	usersResp, err := client.GetUsersBySection(context.Background(), sectionReq)
	if err != nil {
		log.Fatalf("Error getting users by section: %v", err)
	}
	fmt.Printf("Users in Section A: %+v\n", usersResp)

	// Test ModifyUserSeat
	modifyReq := &pb.ModifySeatRequest{
		Email:      user.Email,
		NewSection: "B",
		NewSeat:    2,
	}
	modifiedReceipt, err := client.ModifyUserSeat(context.Background(), modifyReq)
	if err != nil {
		log.Fatalf("Error modifying user seat: %v", err)
	}
	fmt.Printf("User seat modified: %+v\n", modifiedReceipt)

	// Test RemoveUser
	removeResp, err := client.RemoveUser(context.Background(), receiptReq)
	if err != nil {
		log.Fatalf("Error removing user: %v", err)
	}
	fmt.Printf("User removed: %+v\n", removeResp)

}
