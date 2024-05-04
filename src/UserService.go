package src

import (
	"context"
	"errors"
	"sync"

	pb "paymentSystem/proto"
)

type UserServiceServer interface {
	CreateUser(context.Context, *pb.User) (*pb.UserResponse, error)
	UserLogin(context.Context, *pb.User) (*pb.UserResponse, error)
}

type userServiceServer struct {
	pb.UnimplementedUserServiceServer
	users map[string]*pb.User
	mu    sync.Mutex
}

func NewUserServiceServer() *userServiceServer {
	return &userServiceServer{
		users: make(map[string]*pb.User),
	}
}

func (s *userServiceServer) CreateUser(ctx context.Context, in *pb.User) (*pb.UserResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, ok := s.users[in.Username]; ok {
		return nil, errors.New("user already exists")
	}
	s.users[in.Username] = in

	return &pb.UserResponse{Confirmation: "User created successfully"}, nil
}

func (s *userServiceServer) UserLogin(ctx context.Context, in *pb.User) (*pb.UserResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	// Проверьте, существует ли пользователь с таким именем
	user, ok := s.users[in.Username]
	if !ok {
		return nil, errors.New("user does not exist")
	}

	// Проверьте, совпадает ли пароль
	if user.Password != in.Password {
		return nil, errors.New("incorrect password")
	}

	return &pb.UserResponse{Confirmation: "User logged in successfully"}, nil
}
