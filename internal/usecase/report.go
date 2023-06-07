package usecase

import (
	"context"
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

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

	reportMap, err := u.repo.CreateMonthReport(ctx, tx, year, month)
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
	csvfile, err := os.Create("./report.csv")
	if err != nil {
		return fmt.Errorf("create: %w", err)
	}
	cswWriter := csv.NewWriter(csvfile)

	for key, value := range data {
		str1 := "название услуги"
		str2 := strconv.Itoa(int(key))
		str3 := "общая сумма выручки за отчетный период"
		str4 := strconv.FormatFloat(float64(value), 'f', 2, 64)

		var res []string
		res = append(res, str1, str2, str3, str4)
		err = cswWriter.Write(res)
		if err != nil {
			return fmt.Errorf("createReportCSV - cswWriter.Write: %w", err)
		}
	}
	cswWriter.Flush()

	err = csvfile.Close()
	if err != nil {
		return fmt.Errorf("createReportCSV - csvfile.Close: %w", err)
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
