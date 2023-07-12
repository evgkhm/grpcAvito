package report

import (
	"context"
	"encoding/csv"
	"fmt"
	"github.com/sirupsen/logrus"
	"grpcAvito/internal/repository/postgres"
	"grpcAvito/internal/usecase/transactions"
	"os"
	"strconv"
)

type UseCase struct {
	userRepo   postgres.UserRepository
	orderRepo  postgres.OrderRepository
	reportRepo postgres.ReportRepository
	log        *logrus.Logger
	txService  *transactions.TransactionServiceImpl
}

func NewReportUseCase(userRepo postgres.UserRepository, orderRepo postgres.OrderRepository, reportRepo postgres.ReportRepository,
	log *logrus.Logger, txService *transactions.TransactionServiceImpl) *UseCase {
	return &UseCase{
		userRepo:   userRepo,
		orderRepo:  orderRepo,
		reportRepo: reportRepo,
		log:        log,
		txService:  txService,
	}
}

func (u UseCase) CreateMonthReport(ctx context.Context, year, month uint32) error {
	tx, err := u.txService.NewTransaction()
	if err != nil {
		errRollback := u.txService.Rollback(tx)
		if errRollback != nil {
			return fmt.Errorf("usecase - UseCase - CreateMonthReport - u.txService.NewTransaction - u.txService.Rollback: %w", err)
		}
		return err
	}

	err = checkYear(year)
	if err != nil {
		return fmt.Errorf("usecase - UseCase - CreateMonthReport - checkYear: %w", err)
	}

	err = checkMonth(month)
	if err != nil {
		return fmt.Errorf("usecase - UseCase - CreateMonthReport - checkMonth: %w", err)
	}

	reportMap, err := u.reportRepo.CreateMonthReport(ctx, tx, year, month)
	if err != nil {
		errRollback := u.txService.Rollback(tx)
		if errRollback != nil {
			return fmt.Errorf("usecase - UseCase - CreateMonthReport - u.repo.CreateMonthReport - u.txService.Rollback: %w", err)
		}
		return fmt.Errorf("usecase - UseCase - CreateMonthReport - u.repo.CreateMonthReport: %w", err)
	}

	// Создание csv файла
	err = createReportCSV(reportMap)
	if err != nil {
		errRollback := u.txService.Rollback(tx)
		if errRollback != nil {
			return fmt.Errorf("usecase - UseCase - CreateMonthReport - createReportCSV - u.txService.Rollback: %w", err)
		}
		return fmt.Errorf("usecase - UseCase - CreateMonthReport - createReportCSV: %w", err)
	}

	return u.txService.Commit(tx)
}

// createReportCSV функция создает csv файл отсчета из мап файла.
func createReportCSV(data map[uint32]float32) error {
	csvFile, err := os.Create("./report.csv")
	if err != nil {
		return fmt.Errorf("create: %w", err)
	}
	cswWriter := csv.NewWriter(csvFile)

	for key, value := range data {
		const nameService = "название услуги"
		const commonSumProfit = "общая сумма выручки за отчетный период"

		month := strconv.Itoa(int(key))

		sumProfit := strconv.FormatFloat(float64(value), 'f', 2, 64)

		var res []string
		res = append(res, nameService, month, commonSumProfit, sumProfit)
		err = cswWriter.Write(res)
		if err != nil {
			return fmt.Errorf("createReportCSV - cswWriter.Write: %w", err)
		}
	}
	cswWriter.Flush()

	err = csvFile.Close()
	if err != nil {
		return fmt.Errorf("createReportCSV - csvFile.Close: %w", err)
	}

	return nil
}

func checkYear(year uint32) error {
	if year < 1975 || year > 2030 {
		return ErrWrongYear
	}
	return nil
}

func checkMonth(month uint32) error {
	if month > 12 {
		return ErrWrongMonth
	}
	return nil
}
