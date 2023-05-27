package postgres

import (
	"context"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"grpcAvito/internal/entity"
	"time"
)

type ReportRepositoryImpl struct {
	db  *sqlx.DB
	log *logrus.Logger
}

func NewReportRepository(db *sqlx.DB, log *logrus.Logger) *ReportRepositoryImpl {
	return &ReportRepositoryImpl{
		db:  db,
		log: log,
	}
}

type Report struct {
	userData entity.UserRevenue
	Year     int `json:"year"`
	Month    int `json:"month"`
}

func (r ReportRepositoryImpl) GetReport(ctx context.Context, tx *sqlx.Tx, year uint32, month uint32) (map[uint32]float32, error) {
	report := Report{}
	reportMap := make(map[uint32]float32)
	query := `SELECT * FROM revenue 
		WHERE EXTRACT(year FROM curr_date)=$1 
		AND EXTRACT(month FROM curr_date)=$2`
	rows, err := tx.QueryxContext(ctx, query, year, month)
	if err != nil {
		return nil, ErrGetYearMonth
	}
	defer func(rows *sqlx.Rows) {
		err = rows.Close()
		if err != nil {
			return
		}
	}(rows)

	//Добавление в хэш-таблицу данных с БД
	for rows.Next() {
		var stamp time.Time
		err = rows.Scan(&report.userData.Id, &report.userData.IdService, &report.userData.IdOrder, &report.userData.Cost, &stamp)
		report.Year = stamp.Year()
		report.Month = int(stamp.Month())
		reportMap[report.userData.IdService] += report.userData.Cost
	}
	return reportMap, nil
}
