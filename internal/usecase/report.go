package usecase

import (
	"context"
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

func (u UseCase) Report(ctx context.Context, year uint32, month uint32) error {
	tx, err := u.txService.NewTransaction()
	if err != nil {
		_ = u.txService.Rollback(tx)
		return err
	}

	err = checkYear(year)
	if err != nil {
		return fmt.Errorf("usecase: Report: checkYear: %w", err)
	}

	err = checkMonth(month)
	if err != nil {
		return fmt.Errorf("usecase: Report: checkMonth: %w", err)
	}

	reportMap, err := u.repo.GetReport(ctx, tx, year, month)
	if err != nil {
		return fmt.Errorf("usecase: Report: GetReport: %w", err)
	}

	//Создание csv файла
	err = createReportCSV(reportMap)
	if err != nil {
		return fmt.Errorf("usecase: Report: createReportCSV: %w", err)
	}

	return u.txService.Commit(tx)
}

// createReportCSV функция создает csv файл отсчета из мап файла
func createReportCSV(data map[uint32]float32) error {
	csvfile, err := os.Create("./report.csv")
	if err != nil {
		return err
	}
	cswWriter := csv.NewWriter(csvfile)

	for key, value := range data {
		str1 := "название услуги"
		str2 := strconv.Itoa(int(key))
		str3 := "общая сумма выручки за отчетный период"
		str4 := strconv.FormatFloat(float64(value), 'f', 2, 64)

		var res []string
		res = append(res, str1)
		res = append(res, str2)
		res = append(res, str3)
		res = append(res, str4)
		err = cswWriter.Write(res)
		if err != nil {
			return fmt.Errorf("usecase: %w", err)
		}
	}
	cswWriter.Flush()

	err = csvfile.Close()
	if err != nil {
		return fmt.Errorf("usecase: %w", err)
	}

	return nil
}

func checkYear(year uint32) error {
	if year < 1975 || year > 2030 {
		return errWrongYear
	}
	return nil
}

func checkMonth(month uint32) error {
	if month > 12 {
		return errWrongMonth
	}
	return nil
}
