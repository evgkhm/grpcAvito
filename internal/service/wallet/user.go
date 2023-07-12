package wallet

import (
	"context"
	"github.com/pkg/errors"
	"grpcAvito/internal/entity/user"
	"grpcAvito/internal/service/wallet/spec"
)

func (s Service) CreateUser(ctx context.Context, req *spec.CreateUserRequest) (*spec.CreateUserResponse, error) {
	userDTO := user.User{ID: req.Id, Balance: req.Balance}
	err := s.useCase.CreateUser(ctx, &userDTO)
	if err != nil {
		s.log.Errorf("service - Service - CreateUser - s.useCase.CreateUser: %v", err)
		return &spec.CreateUserResponse{Success: false}, errors.Unwrap(err)
	}
	return &spec.CreateUserResponse{Success: true}, nil
}

func (s Service) UserBalanceAccrual(ctx context.Context, req *spec.UserBalanceAccrualRequest) (*spec.UserBalanceAccrualResponse, error) {
	userDTO := user.User{ID: req.Id, Balance: req.Balance}
	err := s.useCase.UserBalanceAccrual(ctx, &userDTO)
	if err != nil {
		s.log.Errorf("service - Service - UserBalanceAccrual - s.useCase.UserBalanceAccrual: %v", err)
		return &spec.UserBalanceAccrualResponse{Success: false}, errors.Unwrap(err)
	}
	return &spec.UserBalanceAccrualResponse{Success: true}, nil
}

func (s Service) GetBalance(ctx context.Context, req *spec.GetBalanceRequest) (*spec.GetBalanceResponse, error) {
	userDTO := user.User{ID: req.Id}
	err := s.useCase.GetBalance(ctx, &userDTO)
	if err != nil {
		s.log.Errorf("service - Service - GetBalance - s.useCase.GetBalance: %v", err)
		return &spec.GetBalanceResponse{Success: false}, errors.Unwrap(err)
	}
	return &spec.GetBalanceResponse{Id: userDTO.ID, Balance: userDTO.Balance, Success: true}, nil
}
