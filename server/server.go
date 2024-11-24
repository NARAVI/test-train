package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"sync"

	pb "test_train/protobuf"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedTrainServiceServer
	mu       sync.Mutex
	tickets  map[string]*pb.TicketReceipt
	sections map[string]map[string]*pb.SeatAllocation
	userData map[string]*pb.User // Track user data by email
}

func NewServer() *server {
	return &server{
		tickets: make(map[string]*pb.TicketReceipt),
		sections: map[string]map[string]*pb.SeatAllocation{
			"A": make(map[string]*pb.SeatAllocation),
			"B": make(map[string]*pb.SeatAllocation),
		},
		userData: make(map[string]*pb.User),
	}
}

func (s *server) PurchaseTicket(ctx context.Context, req *pb.PurchaseTicketRequest) (*pb.TicketReceipt, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	// Assign section and seat number assuming  each section can contain 50  passengers
	section := "A"
	seat := len(s.sections[section]) + 1
	if seat > 50 {
		section = "B"
		seat = len(s.sections[section]) + 1
		if seat > 100 {
			return nil, fmt.Errorf("Train is full")
		}
	}

	// Create ticket receipt
	receipt := &pb.TicketReceipt{
		User:  req.User,
		From:  req.From,
		To:    req.To,
		Price: 20,
		Seat: &pb.SeatAllocation{
			Section: section,
			Seat:    int32(seat),
		},
	}

	// Save ticket and user data
	s.tickets[req.User.Email] = receipt
	s.sections[section][req.User.Email] = receipt.Seat
	s.userData[req.User.Email] = req.User

	log.Printf("Ticket purchased: %+v", receipt)
	return receipt, nil
}

func (s *server) GetReceipt(ctx context.Context, req *pb.UserRequest) (*pb.TicketReceipt, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	receipt, ok := s.tickets[req.Email]
	if !ok {
		return nil, fmt.Errorf("user not found")
	}

	return receipt, nil
}

func (s *server) GetUsersBySection(ctx context.Context, req *pb.SectionRequest) (*pb.UsersResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	users := []*pb.User{}
	for email, _ := range s.sections[req.Section] {
		// Retrieve the user using their email
		user, exists := s.userData[email]
		if exists {
			users = append(users, &pb.User{
				FirstName: user.FirstName,
				LastName:  user.LastName,
				Email:     user.Email,
			})
		}
	}

	return &pb.UsersResponse{Users: users}, nil
}

func (s *server) RemoveUser(ctx context.Context, req *pb.UserRequest) (*pb.EmptyResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	receipt, ok := s.tickets[req.Email]
	if !ok {
		return nil, fmt.Errorf("user not found")
	}

	delete(s.tickets, req.Email)
	delete(s.sections[receipt.Seat.Section], req.Email)

	return &pb.EmptyResponse{}, nil
}

func (s *server) ModifyUserSeat(ctx context.Context, req *pb.ModifySeatRequest) (*pb.TicketReceipt, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	receipt, ok := s.tickets[req.Email]
	if !ok {
		return nil, fmt.Errorf("user not found")
	}

	delete(s.sections[receipt.Seat.Section], req.Email)

	receipt.Seat.Section = req.NewSection
	receipt.Seat.Seat = req.NewSeat
	s.sections[req.NewSection][req.Email] = receipt.Seat

	return receipt, nil
}

func main() {
	listener, err := net.Listen("tcp", ":8111")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterTrainServiceServer(grpcServer, NewServer())

	log.Println("Server is listening on port 8111")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
