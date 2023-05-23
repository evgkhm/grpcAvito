package service

import (
	"context"
	"grpcAvito/internal/entity"
	"grpcAvito/proto"
)

func (s Service) Dereservation(ctx context.Context, req *proto.DereservationReq) (*proto.DereservationReply, error) {
	userReservation := entity.UserReservation{Id: req.Id, IdService: req.IdService, IdOrder: req.IdOrder, Cost: req.Cost}
	err := s.useCase.Dereservation(ctx, userReservation)
	if err != nil {
		s.log.Errorf("revenue user: %v", err)
		return &proto.DereservationReply{Success: false}, err
	}
	return &proto.DereservationReply{Success: true}, nil
}

func (s Service) Reservation(ctx context.Context, req *proto.ReservationReq) (*proto.ReservationReply, error) {
	userReservation := entity.UserReservation{Id: req.Id, IdService: req.IdService, IdOrder: req.IdOrder, Cost: req.Cost}
	err := s.useCase.Reservation(ctx, userReservation)
	if err != nil {
		s.log.Errorf("reservation user: %v", err)
		return &proto.ReservationReply{Success: false}, err
	}
	return &proto.ReservationReply{Success: true}, nil
}
