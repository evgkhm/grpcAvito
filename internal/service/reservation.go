package service

import (
	"context"
	"grpcAvito/internal/entity"
	"grpcAvito/proto"
)

func (s Service) DeleteReservation(ctx context.Context, req *proto.DeleteReservationReq) (*proto.DeleteReservationReply, error) {
	userReservation := entity.UserReservation{Id: req.Id, IdService: req.IdService, IdOrder: req.IdOrder, Cost: req.Cost}
	err := s.useCase.DeleteReservation(ctx, &userReservation)
	if err != nil {
		s.log.Errorf("delete reservation user: %v", err)
		return &proto.DeleteReservationReply{Success: false}, err
	}
	return &proto.DeleteReservationReply{Success: true}, nil
}

func (s Service) Reservation(ctx context.Context, req *proto.ReservationReq) (*proto.ReservationReply, error) {
	userReservation := entity.UserReservation{Id: req.Id, IdService: req.IdService, IdOrder: req.IdOrder, Cost: req.Cost}
	err := s.useCase.Reservation(ctx, &userReservation)
	if err != nil {
		s.log.Errorf("reservation user: %v", err)
		return &proto.ReservationReply{Success: false}, err
	}
	return &proto.ReservationReply{Success: true}, nil
}
