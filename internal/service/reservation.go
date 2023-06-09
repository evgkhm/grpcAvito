package service

import (
	"context"
	"github.com/pkg/errors"
	"grpcAvito/internal/entity"
	"grpcAvito/internal/service/spec"
)

func (s Service) UserOrderDeleteReservation(ctx context.Context, req *spec.UserOrderDeleteReservationRequest) (*spec.UserOrderDeleteReservationReply, error) {
	userReservation := entity.UserReservation{ID: req.Id, ServiceID: req.ServiceId, OrderID: req.OrderId, Cost: req.Cost}
	err := s.useCase.UserOrderDeleteReservation(ctx, &userReservation)
	if err != nil {
		s.log.Errorf("service - Service - UserOrderDeleteReservation - s.useCase.UserOrderDeleteReservation: %v", err)
		return &spec.UserOrderDeleteReservationReply{Success: false}, errors.Unwrap(err)
	}
	return &spec.UserOrderDeleteReservationReply{Success: true}, nil
}

func (s Service) UserOrderReservation(ctx context.Context, req *spec.UserOrderReservationRequest) (*spec.UserOrderReservationReply, error) {
	userReservation := entity.UserReservation{ID: req.Id, ServiceID: req.ServiceId, OrderID: req.OrderId, Cost: req.Cost}
	err := s.useCase.UserOrderReservation(ctx, &userReservation)
	if err != nil {
		s.log.Errorf("service - Service - UserOrderReservation - s.useCase.UserOrderReservation: %v", err)
		return &spec.UserOrderReservationReply{Success: false}, errors.Unwrap(err)
	}
	return &spec.UserOrderReservationReply{Success: true}, nil
}
