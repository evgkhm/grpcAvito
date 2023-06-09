package service

import (
	"context"
	"github.com/pkg/errors"
	"grpcAvito/internal/entity"
	"grpcAvito/internal/service/spec"
)

func (s Service) UserOrderRevenue(ctx context.Context, req *spec.UserOrderRevenueRequest) (*spec.UserOrderRevenueResponse, error) {
	userReservation := entity.UserRevenue{ID: req.Id, ServiceID: req.ServiceId, OrderID: req.OrderId, Cost: req.Cost}
	err := s.useCase.UserOrderRevenue(ctx, &userReservation)
	if err != nil {
		s.log.Errorf("service - Service - UserOrderRevenue - s.useCase.UserOrderRevenue %v", err)
		return &spec.UserOrderRevenueResponse{Success: false}, errors.Unwrap(err)
	}
	return &spec.UserOrderRevenueResponse{Success: true}, nil
}
