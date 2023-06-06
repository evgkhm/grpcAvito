package service

import (
	"context"
	"github.com/pkg/errors"
	"grpcAvito/internal/entity"
	"grpcAvito/internal/service/spec"
)

func (s Service) CreateUser(ctx context.Context, req *spec.CreateUserRequest) (*spec.CreateUserReply, error) {
	userDTO := entity.User{Id: req.Id, Balance: req.Balance}
	err := s.useCase.Create(ctx, &userDTO)
	if err != nil {
		s.log.Errorf("service - Service - CreateUser - s.useCase.CreateUser: %v", err)
		return &spec.CreateUserReply{Success: false}, errors.Unwrap(err)
	}
	return &spec.CreateUserReply{Success: true}, nil
}

func (s Service) UserBalanceAccrual(ctx context.Context, req *spec.UserBalanceAccrualRequest) (*spec.UserBalanceAccrualReply, error) {
	userDTO := entity.User{Id: req.Id, Balance: req.Balance}
	err := s.useCase.Sum(ctx, &userDTO)
	if err != nil {
		s.log.Errorf("service - Service - UserBalanceAccrual - s.useCase.UserBalanceAccrual: %v", err)
		return &spec.UserBalanceAccrualReply{Success: false}, errors.Unwrap(err)
	}
	return &spec.UserBalanceAccrualReply{Success: true}, nil
}

func (s Service) GetBalance(ctx context.Context, req *spec.GetBalanceRequest) (*spec.GetBalanceReply, error) {
	userDTO := entity.User{Id: req.Id}
	err := s.useCase.GetBalance(ctx, &userDTO)
	if err != nil {
		s.log.Errorf("service - Service - GetBalance - s.useCase.GetBalance: %v", err)
		return &spec.GetBalanceReply{Success: false}, errors.Unwrap(err)
	}
	return &spec.GetBalanceReply{Id: userDTO.Id, Balance: userDTO.Balance, Success: true}, nil
}
