package postgres

import (
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
	"grpcAvito/internal/entity"
	"time"
)

type Report struct {
	userData entity.UserRevenue
	Year     int `json:"year"`
	Month    int `json:"month"`
}

func (r Repo) CreateMonthReport(ctx context.Context, tx *sqlx.Tx, year, month uint32) (map[uint32]float32, error) {
	report := Report{}
	reportMap := make(map[uint32]float32)
	query := `SELECT * FROM revenue 
		WHERE EXTRACT(year FROM curr_date)=$1 
		AND EXTRACT(month FROM curr_date)=$2`
	rows, err := tx.QueryxContext(ctx, query, year, month)
	if err != nil {
		return nil, fmt.Errorf("postgres - ReportRepositoryImpl - CreateMonthReport - tx.QueryxContext: %w", ErrGetYearMonth)
	}
	defer func(rows *sqlx.Rows) {
		err = rows.Close()
		if err != nil {
			return
		}
	}(rows)

	// Добавление в хэш-таблицу данных с БД
	for rows.Next() {
		var stamp time.Time
		err = rows.Scan(&report.userData.ID, &report.userData.IDService, &report.userData.IDOrder, &report.userData.Cost, &stamp)
		if err != nil {
			return nil, fmt.Errorf("postgres - ReportRepositoryImpl - CreateMonthReport - rows.Scan: %w", err)
		}
		report.Year = stamp.Year()
		report.Month = int(stamp.Month())
		reportMap[report.userData.IDService] += report.userData.Cost
	}
	return reportMap, nil
}
