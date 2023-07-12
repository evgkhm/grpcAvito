package wallet

import (
	"context"
	"github.com/pkg/errors"
	"grpcAvito/internal/entity/user"
	"grpcAvito/internal/service/wallet/spec"
)

func (s Service) UserOrderDeleteReservation(ctx context.Context, req *spec.UserOrderDeleteReservationRequest) (*spec.UserOrderDeleteReservationResponse, error) {
	userReservation := user.Reservation{ID: req.Id, ServiceID: req.ServiceId, OrderID: req.OrderId, Cost: req.Cost}
	err := s.useCase.UserOrderDeleteReservation(ctx, &userReservation)
	if err != nil {
		s.log.Errorf("service - Service - UserOrderDeleteReservation - s.useCase.UserOrderDeleteReservation: %v", err)
		return &spec.UserOrderDeleteReservationResponse{Success: false}, errors.Unwrap(err)
	}
	return &spec.UserOrderDeleteReservationResponse{Success: true}, nil
}

func (s Service) UserOrderReservation(ctx context.Context, req *spec.UserOrderReservationRequest) (*spec.UserOrderReservationResponse, error) {
	userReservation := user.Reservation{ID: req.Id, ServiceID: req.ServiceId, OrderID: req.OrderId, Cost: req.Cost}
	err := s.useCase.UserOrderReservation(ctx, &userReservation)
	if err != nil {
		s.log.Errorf("service - Service - UserOrderReservation - s.useCase.UserOrderReservation: %v", err)
		return &spec.UserOrderReservationResponse{Success: false}, errors.Unwrap(err)
	}
	return &spec.UserOrderReservationResponse{Success: true}, nil
}

func (s Service) UserOrderRevenue(ctx context.Context, req *spec.UserOrderRevenueRequest) (*spec.UserOrderRevenueResponse, error) {
	userRevenue := user.Revenue{ID: req.Id, ServiceID: req.ServiceId, OrderID: req.OrderId, Cost: req.Cost}
	err := s.useCase.UserOrderRevenue(ctx, &userRevenue)
	if err != nil {
		s.log.Errorf("service - Service - UserOrderRevenue - s.useCase.UserOrderRevenue %v", err)
		return &spec.UserOrderRevenueResponse{Success: false}, errors.Unwrap(err)
	}
	return &spec.UserOrderRevenueResponse{Success: true}, nil
}
