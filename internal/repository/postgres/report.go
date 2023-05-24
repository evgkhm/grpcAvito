package postgres

import (
	"context"
	"database/sql"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"grpcAvito/internal/entity"
	"time"
)

func (r ReservationRepositoryImpl) GetReport(ctx context.Context, tx *sqlx.Tx, year uint32, month uint32) (map[uint32]float32, error) {
	userRevenue := entity.UserRevenue{}
	reportMap := make(map[uint32]float32)
	rows, err := tx.Query("SELECT * FROM revenue WHERE EXTRACT(year FROM curr_date) = $1 AND EXTRACT(month FROM curr_date) = $2", year, month)
	if err != nil {
		errGetYearMonth := errors.New("getting year or month")
		return nil, errGetYearMonth
	}
	defer func(rows *sql.Rows) {
		err = rows.Close()
		if err != nil {
			return
		}
	}(rows)

	//Добавление в хэш-таблицу данных с БД
	for rows.Next() {
		var stamp time.Time
		err = rows.Scan(&userRevenue.Id, &userRevenue.IdService, &userRevenue.IdOrder, &userRevenue.Cost, &stamp)
		//dataDB.Year = stamp.Year()
		//dataDB.Month = int(stamp.Month())
		reportMap[userRevenue.IdService] += userRevenue.Cost
	}
	return reportMap, nil
}
