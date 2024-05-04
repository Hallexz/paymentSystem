package src

import (
	"context"
	"errors"

	pb "paymentSystem/proto"
)

type TransactionServer struct {
	pb.UnimplementedTransactionServiceServer
	Transactions map[string][]*pb.Transaction
}

func (s *TransactionServer) GetTransactionHistory(ctx context.Context, in *pb.TransactionRequest) (*pb.TransactionResponse, error) {
	user := in.GetUser()
	if user == "" {
		return nil, errors.New("Invalid user")
	}

	userTransactions, ok := s.Transactions[user]
	if !ok {
		return nil, errors.New("No transactions found for this user")
	}

	return &pb.TransactionResponse{Transactions: userTransactions}, nil
}
