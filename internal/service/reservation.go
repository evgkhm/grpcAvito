package service

import (
	"context"
	"grpcAvito/internal/entity"
	"grpcAvito/proto"
)

func (s Service) Reservation(ctx context.Context, req *proto.ReservationReq) (*proto.ReservationReply, error) {
	//s.log.Printf("Received: %v", req.Id)
	userReservation := entity.UserReservation{Id: req.Id, IdService: req.IdService, IdOrder: req.IdOrder, Cost: req.Cost}
	err := s.useCase.Reservation(ctx, userReservation)
	if err != nil {
		s.log.Errorf("reservation user %v", err)
		return &proto.ReservationReply{Success: false}, err
	}
	return &proto.ReservationReply{Success: true}, nil
}
