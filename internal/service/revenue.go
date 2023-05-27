package service

import (
	"context"
	"grpcAvito/internal/entity"
	"grpcAvito/proto"
)

func (s Service) Revenue(ctx context.Context, req *proto.RevenueReq) (*proto.RevenueReply, error) {
	userReservation := entity.UserRevenue{Id: req.Id, IdService: req.IdService, IdOrder: req.IdOrder, Cost: req.Cost}
	err := s.useCase.Revenue(ctx, &userReservation)
	if err != nil {
		s.log.Errorf("revenue user: %v", err)
		return &proto.RevenueReply{Success: false}, err
	}
	return &proto.RevenueReply{Success: true}, nil
}
