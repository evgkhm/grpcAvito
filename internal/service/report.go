package service

import (
	"context"
	"github.com/pkg/errors"
	"grpcAvito/internal/service/spec"
)

func (s Service) CreateMonthReport(ctx context.Context, req *spec.CreateMonthReportRequest) (*spec.CreateMonthReportReply, error) {
	Year, Month := req.Year, req.Month
	err := s.useCase.Report(ctx, Year, Month)
	if err != nil {
		s.log.Errorf("service - Service - s.useCase.CreateMonthReport: %v", err)
		return &spec.CreateMonthReportReply{Success: false}, errors.Unwrap(err)
	}
	return &spec.CreateMonthReportReply{Success: true}, nil

}
