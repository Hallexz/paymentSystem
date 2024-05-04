package src

import (
	"sync"

	pb "paymentSystem/proto"
)

type Bank struct {
	accounts map[string]*pb.BankAccount
	mu       sync.RWMutex
}

func NewBank() *Bank {
	return &Bank{
		accounts: make(map[string]*pb.BankAccount),
	}
}

func (b *Bank) AddBankAccount(account *pb.BankAccount) *pb.BankAccountResponse {
	b.mu.Lock()
	defer b.mu.Unlock()

	// Проверьте, существует ли уже счет
	if _, ok := b.accounts[account.AccountNumber]; ok {
		return &pb.BankAccountResponse{Confirmation: "Bank account already exists"}
	}

	// Добавьте новый счет
	b.accounts[account.AccountNumber] = account

	return &pb.BankAccountResponse{Confirmation: "Bank account added successfully"}
}

func addBankAccount(_ *pb.BankAccount) *pb.BankAccountResponse {
	return &pb.BankAccountResponse{Confirmation: "Bank account added successfully"}
}
