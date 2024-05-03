package src

import (
	"context"
	"log"
	"time"

	pb "paymentSystem/proto"

	"google.golang.org/grpc"
)

const (
	address     = "localhost:50051"
	defaultName = "world"
)

func ClientPay() {
	conn, err := grpc.Dial(address, grpc.WithAuthority("your_authority"), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewPaymentServiceClient(conn)

	ctx, cansel := context.WithTimeout(context.Background(), time.Second)
	defer cansel()
	r, err := c.SendPayment(ctx, &pb.PaymentRequest{User: "Alice", Amount: 100})
	if err != nil {
		log.Fatalf("could not send payment: %v", err)
	}
	log.Printf("Payment Confirmation: %s", r.GetConfirmation())
}
