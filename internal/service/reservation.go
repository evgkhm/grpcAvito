package service

import (
	"context"
	"github.com/pkg/errors"
	"grpcAvito/internal/entity"
	"grpcAvito/proto"
)

func (s Service) DeleteReservation(ctx context.Context, req *proto.DeleteReservationReq) (*proto.DeleteReservationReply, error) {
	userReservation := entity.UserReservation{Id: req.Id, IdService: req.IdService, IdOrder: req.IdOrder, Cost: req.Cost}
	err := s.useCase.DeleteReservation(ctx, &userReservation)
	if err != nil {
		s.log.Errorf("service: DeleteReservation: %v", err)
		return &proto.DeleteReservationReply{Success: false}, errors.Unwrap(err)
	}
	return &proto.DeleteReservationReply{Success: true}, nil
}

func (s Service) Reservation(ctx context.Context, req *proto.ReservationReq) (*proto.ReservationReply, error) {
	userReservation := entity.UserReservation{Id: req.Id, IdService: req.IdService, IdOrder: req.IdOrder, Cost: req.Cost}
	err := s.useCase.Reservation(ctx, &userReservation)
	if err != nil {
		s.log.Errorf("service: Reservation: %v", err)
		return &proto.ReservationReply{Success: false}, errors.Unwrap(err)
	}
	return &proto.ReservationReply{Success: true}, nil
}
