package entity

type UserRevenue struct {
	ID        uint32  `json:"id"`
	IDService uint32  `json:"id_service"`
	IDOrder   uint32  `json:"id_order"`
	Cost      float32 `json:"cost"`
}
