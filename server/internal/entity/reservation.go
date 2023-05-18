package entity

type UserReservation struct {
	Id        int     `json:"id"`
	IdService int     `json:"id_service"`
	IdOrder   int     `json:"id_order"`
	Cost      float64 `json:"cost"`
}
