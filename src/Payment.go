package src

import (
	"context"
	"paymentSystem/src/api"

	pb "paymentSystem/proto"
)

type PaymentServiceServer interface {
	InitiatePayment(context.Context, *pb.PaymentRequest) (*pb.PaymentResponse, error)
	ProcessPayment(context.Context, *pb.PaymentRequest) (*pb.PaymentResponse, error)
	ConfirmPayment(context.Context, *pb.PaymentRequest) (*pb.PaymentResponse, error)
	RefundPayment(context.Context, *pb.PaymentRequest) (*pb.PaymentResponse, error)
}

type PaymentService struct {
	pb.UnimplementedPaymentServiceServer
}

func (s *PaymentService) InitiatePayment(ctx context.Context, in *pb.PaymentRequest) (*pb.PaymentResponse, error) {
	return &pb.PaymentResponse{Confirmation: "Payment initiated successfully"}, nil
}

func (s *PaymentService) ProcessPayment(ctx context.Context, in *pb.PaymentRequest) (*pb.PaymentResponse, error) {
	return &pb.PaymentResponse{Confirmation: "Payment processed successfully"}, nil
}

func (s *PaymentService) ConfirmPayment(ctx context.Context, in *pb.PaymentRequest) (*pb.PaymentResponse, error) {
	return &pb.PaymentResponse{Confirmation: "Payment confirmed successfully"}, nil
}

func (s *PaymentService) RefundPayment(ctx context.Context, in *pb.PaymentRequest) (*pb.PaymentResponse, error) {
	return &pb.PaymentResponse{Confirmation: "Payment refunded successfully"}, nil
}

