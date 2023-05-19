package service

import (
	"context"
	"github.com/sirupsen/logrus"
	"grpcAvito/internal/entity"
	"grpcAvito/internal/repository/postgres"
	"grpcAvito/internal/usecase"
	"grpcAvito/proto"
)

type UsersServiceImpl struct {
	repo      postgres.UsersRepository
	txService *usecase.TransactionServiceImpl
	log       *logrus.Logger
}

func NewUserService(repo postgres.UsersRepository, txService *usecase.TransactionServiceImpl, log *logrus.Logger) *UsersServiceImpl {
	return &UsersServiceImpl{
		repo:      repo,
		txService: txService,
		log:       log,
	}
}

type ServerServer interface {
	Create(context.Context, *proto.CreateReq) (*proto.CreateReply, error)
	Sum(context.Context, *proto.SumReq) (*proto.SumReply, error)
}

func (s Service) Create(ctx context.Context, req *proto.CreateReq) (*proto.CreateReply, error) {
	//s.log.Printf("Received: %v", req.Id)
	userDTO := entity.User{Id: req.Id, Balance: req.Balance}
	err := s.useCase.Create(ctx, userDTO)
	if err != nil {
		s.log.Errorf("creating user %v", err)
		return &proto.CreateReply{Result: "not ok "}, err
	}
	return &proto.CreateReply{Result: "ok "}, nil
}

func (s Service) Sum(ctx context.Context, req *proto.SumReq) (*proto.SumReply, error) {
	//s.log.Printf("Received: %v", req.Id)
	userDTO := entity.User{Id: req.Id, Balance: req.Balance}
	err := s.useCase.Sum(ctx, userDTO)
	if err != nil {
		s.log.Errorf("summing user %v", err)
		return &proto.SumReply{Result: "not ok "}, err
	}
	return &proto.SumReply{Result: "ok "}, nil
}
