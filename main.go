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
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterPaymentServiceServer(s, &src.Server{})
	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()
	src.ClientPay()

	bank := src.NewBank()
	account := &pb.BankAccount{
		User:          "User1",
		AccountNumber: "1234567890",
		BankName:      "Bank1",
		CardNumber:    "1111222233334444",
		CardExpiry:    "01/23",
		CardCVV:       "123",
	}

	// Добавьте банковский счет в банк
	response := bank.AddBankAccount(account)

	// Выведите подтверждение
	fmt.Println(response.Confirmation)

}
