package service

import (
	"context"
	"github.com/pkg/errors"
	"grpcAvito/internal/entity"
	"grpcAvito/internal/service/spec"
)

func (s Service) UserOrderRevenue(ctx context.Context, req *spec.UserOrderRevenueRequest) (*spec.UserOrderRevenueReply, error) {
	userReservation := entity.UserRevenue{Id: req.Id, IdService: req.IdService, IdOrder: req.IdOrder, Cost: req.Cost}
	err := s.useCase.Revenue(ctx, &userReservation)
	if err != nil {
		s.log.Errorf("service - Service - UserOrderRevenue - s.useCase.UserOrderRevenue %v", err)
		return &spec.UserOrderRevenueReply{Success: false}, errors.Unwrap(err)
	}
	return &spec.UserOrderRevenueReply{Success: true}, nil
}
