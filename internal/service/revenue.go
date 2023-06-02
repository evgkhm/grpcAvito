package service

import (
	"context"
	"github.com/pkg/errors"
	"grpcAvito/internal/entity"
	"grpcAvito/proto"
)

func (s Service) Revenue(ctx context.Context, req *proto.RevenueReq) (*proto.RevenueReply, error) {
	userReservation := entity.UserRevenue{Id: req.Id, IdService: req.IdService, IdOrder: req.IdOrder, Cost: req.Cost}
	err := s.useCase.Revenue(ctx, &userReservation)
	if err != nil {
		s.log.Errorf("service: Revenue: %v", err)
		return &proto.RevenueReply{Success: false}, errors.Unwrap(err)
	}
	return &proto.RevenueReply{Success: true}, nil
}
