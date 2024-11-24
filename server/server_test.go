package main

import (
	"context"
	"testing"

	pb "test_train/protobuf"

	"github.com/stretchr/testify/assert"
)

func TestPurchaseTicket(t *testing.T) {
	server := NewServer()

	req := &pb.PurchaseTicketRequest{
		User: &pb.User{
			FirstName: "John",
			LastName:  "Doe",
			Email:     "john.doe@example.com",
		},
		From: "London",
		To:   "France",
	}

	receipt, err := server.PurchaseTicket(context.Background(), req)
	assert.NoError(t, err, "error purchasing ticket")
	assert.NotNil(t, receipt, "receipt should not be nil")
	assert.Equal(t, "A", receipt.Seat.Section, "expected section A")
	assert.Equal(t, int32(1), receipt.Seat.Seat, "expected seat 1")
	assert.Equal(t, req.User.Email, receipt.User.Email, "user email should match")
}

func TestGetReceipt(t *testing.T) {
	server := NewServer()

	// Add a ticket
	req := &pb.PurchaseTicketRequest{
		User: &pb.User{
			FirstName: "John",
			LastName:  "Doe",
			Email:     "john.doe@example.com",
		},
		From: "London",
		To:   "France",
	}
	server.PurchaseTicket(context.Background(), req)

	// Fetch the receipt
	userReq := &pb.UserRequest{Email: "john.doe@example.com"}
	receipt, err := server.GetReceipt(context.Background(), userReq)
	assert.NoError(t, err, "error fetching receipt")
	assert.NotNil(t, receipt, "receipt should not be nil")
	assert.Equal(t, req.User.Email, receipt.User.Email, "user email should match")
}

func TestGetUsersBySection(t *testing.T) {
	server := NewServer()

	// Add users to sections
	server.PurchaseTicket(context.Background(), &pb.PurchaseTicketRequest{
		User: &pb.User{
			FirstName: "John",
			LastName:  "Doe",
			Email:     "john.doe@example.com",
		},
		From: "London",
		To:   "France",
	})
	server.PurchaseTicket(context.Background(), &pb.PurchaseTicketRequest{
		User: &pb.User{
			FirstName: "Jane",
			LastName:  "Smith",
			Email:     "jane.smith@example.com",
		},
		From: "London",
		To:   "France",
	})

	// Fetch users in Section A
	sectionReq := &pb.SectionRequest{Section: "A"}
	resp, err := server.GetUsersBySection(context.Background(), sectionReq)
	assert.NoError(t, err, "error fetching users by section")
	assert.Len(t, resp.Users, 2, "expected 2 users in section A")
}

func TestRemoveUser(t *testing.T) {
	server := NewServer()

	// Add a user
	req := &pb.PurchaseTicketRequest{
		User: &pb.User{
			FirstName: "John",
			LastName:  "Doe",
			Email:     "john.doe@example.com",
		},
		From: "London",
		To:   "France",
	}
	server.PurchaseTicket(context.Background(), req)

	// Remove the user
	userReq := &pb.UserRequest{Email: "john.doe@example.com"}
	_, err := server.RemoveUser(context.Background(), userReq)
	assert.NoError(t, err, "error removing user")

	// Verify the user is removed
	_, err = server.GetReceipt(context.Background(), userReq)
	assert.Error(t, err, "user should not exist after removal")
}

func TestModifyUserSeat(t *testing.T) {
	server := NewServer()

	// Add a user
	req := &pb.PurchaseTicketRequest{
		User: &pb.User{
			FirstName: "John",
			LastName:  "Doe",
			Email:     "john.doe@example.com",
		},
		From: "London",
		To:   "France",
	}
	server.PurchaseTicket(context.Background(), req)

	// Modify the user's seat
	modifyReq := &pb.ModifySeatRequest{
		Email:      "john.doe@example.com",
		NewSection: "B",
		NewSeat:    5,
	}
	receipt, err := server.ModifyUserSeat(context.Background(), modifyReq)
	assert.NoError(t, err, "error modifying user seat")
	assert.Equal(t, "B", receipt.Seat.Section, "expected section B")
	assert.Equal(t, int32(5), receipt.Seat.Seat, "expected seat 5")
}

