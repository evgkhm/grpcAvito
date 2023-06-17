package service

import (
	"context"
	"github.com/pkg/errors"
	"grpcAvito/internal/service/spec"
)

func (s Service) CreateMonthReport(ctx context.Context, req *spec.CreateMonthReportRequest) (*spec.CreateMonthReportResponse, error) {
	Year, Month := req.Year, req.Month
	err := s.useCase.CreateMonthReport(ctx, Year, Month)
	if err != nil {
		s.log.Errorf("service - Service - s.useCase.CreateMonthReport: %v", err)
		return &spec.CreateMonthReportResponse{Success: false}, errors.Unwrap(err)
	}
	return &spec.CreateMonthReportResponse{Success: true}, nil

}
