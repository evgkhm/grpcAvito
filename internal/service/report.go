package service

import (
	"context"
	"grpcAvito/proto"
)

func (s Service) Report(ctx context.Context, req *proto.ReportReq) (*proto.ReportReply, error) {
	Year, Month := req.Year, req.Month
	err := s.useCase.Report(ctx, Year, Month)
	if err != nil {
		s.log.Errorf("report: %v", err)
		return &proto.ReportReply{Success: false}, err
	}
	return &proto.ReportReply{Success: true}, nil

}
