package service

import (
	"context"
	"github.com/sirupsen/logrus"
	"grpcAvito/internal/usecase"
	"grpcAvito/proto"
)

type ServerServer interface {
	Create(context.Context, *proto.CreateReq) (*proto.CreateReply, error)
}

type Service struct {
	proto.UnimplementedServerServer
	useCase *usecase.UseCase
	log     *logrus.Logger
}

func (s Service) Create(ctx context.Context, req *proto.CreateReq) (*proto.CreateReply, error) {
	s.log.Printf("Received: %v", req.Id)
	err := s.useCase.Create(ctx, req.Id, req.Balance)
	if err != nil {
		s.log.Errorf("creating user %v", err)
		//return nil, err
		return &proto.CreateReply{Id: "not ok "}, err
	}
	return &proto.CreateReply{Id: "ok "}, nil
	//status.New(codes.OK, "OK")
	//return &proto.CreateReply{Id: "1 "}, nil
}

func NewService(useCase *usecase.UseCase, log *logrus.Logger) *Service {

	//usersService := NewUserService(repo, txService, log)
	return &Service{
		UnimplementedServerServer: proto.UnimplementedServerServer{},
		useCase:                   useCase,
		log:                       log,
	}
}
