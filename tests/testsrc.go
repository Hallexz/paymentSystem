package tests

import (
	"context"
	"log"
	"net"
	"testing"

	pb "paymentSystem/proto"
	"paymentSystem/src"

	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
)

func TestAddBankAccount(t *testing.T) {
	bank := src.NewBank()
	account := &pb.BankAccount{AccountNumber: "123456", CardNumber: "100000"}

	resp := bank.AddBankAccount(account)
	if resp.Confirmation != "Bank account added successfully" {
		t.Errorf("Expected 'Bank account added successfully', but got %s", resp.Confirmation)
	}

	resp = bank.AddBankAccount(account)
	if resp.Confirmation != "Bank account already exists" {
		t.Errorf("Expected 'Bank account already exists', but got %s", resp.Confirmation)
	}
}

func TestGetTransactionHistory(t *testing.T) {
	s := &src.TransactionServer{
		Transactions: map[string][]*pb.Transaction{
			"user1": {
				{
					User:      "Alice",
					Amount:    100,
					Currency:  "USD",
					Recipient: "Bob",
					Status:    "ok",
				},
				{
					User:      "Charlie",
					Amount:    2000,
					Currency:  "RUB",
					Recipient: "Dave",
					Status:    "fatal",
				},
			},
		},
	}

	t.Run("Success", func(t *testing.T) {
		req := &pb.TransactionRequest{User: "user1"}
		resp, err := s.GetTransactionHistory(context.Background(), req)
		assert.NoError(t, err)
		assert.NotNil(t, resp)
		assert.Equal(t, 2, len(resp.Transactions))
	})

	t.Run("InvalidUser", func(t *testing.T) {
		req := &pb.TransactionRequest{User: ""}
		resp, err := s.GetTransactionHistory(context.Background(), req)
		assert.Error(t, err)
		assert.Nil(t, resp)
		assert.Equal(t, "Invalid user", err.Error())
	})

	t.Run("NoTransactions", func(t *testing.T) {
		req := &pb.TransactionRequest{User: "user2"}
		resp, err := s.GetTransactionHistory(context.Background(), req)
		assert.Error(t, err)
		assert.Nil(t, resp)
		assert.Equal(t, "No transactions found for this user", err.Error())
	})
}

func TestUserServiceServer(t *testing.T) {
	s := src.NewUserServiceServer()

	t.Run("CreateUser", func(t *testing.T) {
		user := &pb.User{Username: "testuser", Password: "testpass"}
		resp, err := s.CreateUser(context.Background(), user)
		assert.NoError(t, err)
		assert.Equal(t, "User created successfully", resp.Confirmation)

		_, err = s.CreateUser(context.Background(), user)
		assert.Error(t, err)
	})

	t.Run("UserLogin", func(t *testing.T) {
		user := &pb.User{Username: "testuser", Password: "testpass"}
		resp, err := s.UserLogin(context.Background(), user)
		assert.NoError(t, err)
		assert.Equal(t, "User logged in successfully", resp.Confirmation)

		userBadPass := &pb.User{Username: "testuser", Password: "wrongpass"}
		_, err = s.UserLogin(context.Background(), userBadPass)
		assert.Error(t, err)
	})
}

const bufSize = 1024 * 1024

var lis *bufconn.Listener

func init() {
	lis = bufconn.Listen(bufSize)
	s := grpc.NewServer()
	pb.RegisterPaymentServiceServer(s, &mockPaymentServiceServer{})
	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatalf("Server exited with error: %v", err)
		}
	}()
}

type mockPaymentServiceServer struct {
	pb.UnimplementedPaymentServiceServer
}

func (s *mockPaymentServiceServer) InitiatePayment(ctx context.Context, in *pb.PaymentRequest) (*pb.PaymentResponse, error) {
	return &pb.PaymentResponse{Confirmation: "Payment Successful"}, nil
}

func bufDialer(context.Context, string) (net.Conn, error) {
	return lis.Dial()
}

func TestClientPay(t *testing.T) {
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}
	defer conn.Close()
	client := pb.NewPaymentServiceClient(conn)

	src.ClientPay(client)
}
