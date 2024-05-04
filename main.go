package main

import (
	"fmt"
	"log"
	"net"
	"paymentSystem/src"

	"google.golang.org/grpc"

	pb "paymentSystem/proto"
)

func main() {
	src.ClientPay()

	s := src.Server{}
	ps := src.PaymentService{}

	ts := &src.TransactionServer{
		Transactions: make(map[string][]*pb.Transaction),
	}

	us := src.NewUserServiceServer()

	bank := src.NewBank()

	newAccount := &pb.BankAccount{
		User:          "John Doe",
		AccountNumber: "123456789",
		BankName:      "My Bank",
		CardNumber:    "1000000000000",
		CardExpiry:    "12/24",
		CardCVV:       "123",
	}

	response := bank.AddBankAccount(newAccount)
	fmt.Println(response.Confirmation)


	lis, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()

	pb.RegisterPaymentServiceServer(grpcServer, &s)
	pb.RegisterPaymentServiceServer(grpcServer, &ps)

	pb.RegisterTransactionServiceServer(grpcServer, ts)

	pb.RegisterUserServiceServer(grpcServer, us)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
