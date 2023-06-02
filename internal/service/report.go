package service

import (
	"context"
	"github.com/pkg/errors"
	"grpcAvito/proto"
)

func (s Service) Report(ctx context.Context, req *proto.ReportReq) (*proto.ReportReply, error) {
	Year, Month := req.Year, req.Month
	err := s.useCase.Report(ctx, Year, Month)
	if err != nil {
		s.log.Errorf("service: Report: %v", err)
		return &proto.ReportReply{Success: false}, errors.Unwrap(err)
	}
	return &proto.ReportReply{Success: true}, nil

}
