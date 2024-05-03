package src

import (
	"context"
	"log"

	pb "paymentSystem/proto"
)

type Server struct {
	pb.UnimplementedPaymentServiceServer
}

func (s *Server) SendPayment(ctx context.Context, in *pb.PaymentRequest) (*pb.PaymentResponse, error) {
	log.Printf("Received payment from %s for amount %d", in.User, in.Amount)
	return &pb.PaymentResponse{Confirmation: "Payment received!"}, nil
}
