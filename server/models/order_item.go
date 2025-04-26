package models

type OrderItem struct {
	ID       *int     `json:"id"`
	ItemID   *int     `json:"item_id"`
	OrderID  *int     `json:"order_id"`
	Quantity *int     `json:"quantity"`
	Price    *float64 `json:"price"`
}

func NewOrderItem() *OrderItem {
	return &OrderItem{
		ID:       new(int),
		ItemID:   new(int),
		OrderID:  new(int),
		Quantity: new(int),
		Price:    new(float64),
	}
}

func NewOrderItemWithNil() *OrderItem {
	return &OrderItem{}
}
