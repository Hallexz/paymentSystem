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

	// Создание экземпляра TransactionServer
	ts := &src.TransactionServer{
		Transactions: make(map[string][]*pb.Transaction),
	}

	// Создание экземпляра UserServiceServer
	us := src.NewUserServiceServer()

	// Создание экземпляра Bank
	bank := src.NewBank()

	// Создание нового банковского счета
	newAccount := &pb.BankAccount{
		User:          "John Doe",
		AccountNumber: "123456789",
		BankName:      "My Bank",
		CardNumber:    "1000000000000",
		CardExpiry:    "12/24",
		CardCVV:       "123",
	}

	// Добавление нового банковского счета
	response := bank.AddBankAccount(newAccount)
	fmt.Println(response.Confirmation)

	// Запуск сервера
	lis, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()

	// Регистрация серверов
	pb.RegisterPaymentServiceServer(grpcServer, &s)
	pb.RegisterPaymentServiceServer(grpcServer, &ps)

	// Регистрация TransactionServer
	pb.RegisterTransactionServiceServer(grpcServer, ts)

	// Регистрация UserServiceServer
	pb.RegisterUserServiceServer(grpcServer, us)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
