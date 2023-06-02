package service

import (
	"context"
	"github.com/pkg/errors"
	"grpcAvito/internal/entity"
	"grpcAvito/proto"
)

func (s Service) Create(ctx context.Context, req *proto.CreateReq) (*proto.CreateReply, error) {
	userDTO := entity.User{Id: req.Id, Balance: req.Balance}
	err := s.useCase.Create(ctx, &userDTO)
	if err != nil {
		s.log.Errorf("service: Create: %v", err)
		return &proto.CreateReply{Success: false}, errors.Unwrap(err)
	}
	return &proto.CreateReply{Success: true}, nil
}

func (s Service) Sum(ctx context.Context, req *proto.SumReq) (*proto.SumReply, error) {
	userDTO := entity.User{Id: req.Id, Balance: req.Balance}
	err := s.useCase.Sum(ctx, &userDTO)
	if err != nil {
		s.log.Errorf("service: Sum: %v", err)
		return &proto.SumReply{Success: false}, errors.Unwrap(err)
	}
	return &proto.SumReply{Success: true}, nil
}

func (s Service) GetBalance(ctx context.Context, req *proto.BalanceReq) (*proto.BalanceReply, error) {
	userDTO := entity.User{Id: req.Id}
	err := s.useCase.GetBalance(ctx, &userDTO)
	if err != nil {
		s.log.Errorf("service: GetBalance: %v", err)
		return &proto.BalanceReply{Success: false}, errors.Unwrap(err)
	}
	return &proto.BalanceReply{Id: userDTO.Id, Balance: userDTO.Balance, Success: true}, nil
}
