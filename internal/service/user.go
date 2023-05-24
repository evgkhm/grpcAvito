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

func (s Service) Create(ctx context.Context, req *proto.CreateReq) (*proto.CreateReply, error) {
	userDTO := entity.User{Id: req.Id, Balance: req.Balance}
	err := s.useCase.Create(ctx, userDTO)
	if err != nil {
		s.log.Errorf("creating user: %v", err)
		return &proto.CreateReply{Success: false}, err
	}
	return &proto.CreateReply{Success: true}, nil
}

func (s Service) Sum(ctx context.Context, req *proto.SumReq) (*proto.SumReply, error) {
	userDTO := entity.User{Id: req.Id, Balance: req.Balance}
	err := s.useCase.Sum(ctx, userDTO)
	if err != nil {
		s.log.Errorf("summing user: %v", err)
		return &proto.SumReply{Success: false}, err
	}
	return &proto.SumReply{Success: true}, nil
}

func (s Service) GetBalance(ctx context.Context, req *proto.BalanceReq) (*proto.BalanceReply, error) {
	userDTO := entity.User{Id: req.Id}
	err := s.useCase.GetBalance(ctx, userDTO)
	if err != nil {
		s.log.Errorf("get balance user: %v", err)
		return &proto.BalanceReply{Success: false}, err
	}
	return &proto.BalanceReply{Id: userDTO.Id, Balance: userDTO.Balance, Success: true}, nil
}
