package user

type Revenue struct {
	ID        uint32  `json:"id"`
	ServiceID uint32  `json:"service_id"`
	OrderID   uint32  `json:"order_id"`
	Cost      float32 `json:"cost"`
}
