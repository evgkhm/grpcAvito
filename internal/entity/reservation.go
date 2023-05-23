package entity

type UserReservation struct {
	Id        uint32  `json:"id"`
	IdService uint32  `json:"id_service"`
	IdOrder   uint32  `json:"id_order"`
	Cost      float32 `json:"cost"`
}
