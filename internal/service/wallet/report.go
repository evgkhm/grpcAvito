package wallet

import (
	"context"
	"github.com/pkg/errors"
	"grpcAvito/internal/service/wallet/spec"
)

func (s Service) CreateMonthReport(ctx context.Context, req *spec.CreateMonthReportRequest) (*spec.CreateMonthReportResponse, error) {
	year, month := req.Year, req.Month
	err := s.useCase.CreateMonthReport(ctx, year, month)
	if err != nil {
		s.log.Errorf("service - Service - s.useCase.CreateMonthReport: %v", err)
		return &spec.CreateMonthReportResponse{Success: false}, errors.Unwrap(err)
	}
	return &spec.CreateMonthReportResponse{Success: true}, nil

}
