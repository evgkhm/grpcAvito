package service

import (
	"context"
	"github.com/sirupsen/logrus"
	"grpcAvito/internal/usecase"
	"grpcAvito/proto"
)

type Service struct {
	proto.UnimplementedServerServer
	useCase *usecase.UseCase
	log     *logrus.Logger
}

func New(useCase *usecase.UseCase, log *logrus.Logger) *Service {
	return &Service{
		UnimplementedServerServer: proto.UnimplementedServerServer{},
		useCase:                   useCase,
		log:                       log,
	}
}

type ServerServer interface {
	Create(context.Context, *proto.CreateReq) (*proto.CreateReply, error)
	Sum(context.Context, *proto.SumReq) (*proto.SumReply, error)
	Revenue(ctx context.Context, req *proto.RevenueReq) (*proto.RevenueReply, error)
	Dereservation(ctx context.Context, req *proto.DereservationReq) (*proto.DereservationReply, error)
	Reservation(ctx context.Context, req *proto.ReservationReq) (*proto.ReservationReply, error)
	Report(ctx context.Context, req *proto.ReportReq) (*proto.ReportReply, error)
	GetBalance(ctx context.Context, req *proto.BalanceReq) (*proto.BalanceReply, error)
}
