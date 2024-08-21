package models

import "time"

type Order struct {
	ID             *int       `json:"id,omitempty"`
	Subtotal       *float64   `json:"subtotal,omitempty"`
	Total          *float64   `json:"total,omitempty"`
	GST            *float64   `json:"gst,omitempty"`
	PST            *float64   `json:"pst,omitempty"`
	Discount       *float64   `json:"discount,omitempty"`
	Timestamp      *time.Time `json:"timestamp,omitempty"`
	Void           *bool      `json:"void,omitempty"`
	Paid           *bool      `json:"paid,omitempty"`
	Customizations *string    `json:"customizations,omitempty"`
	Category       *string    `json:"category,omitempty"`
	CustomerID     *int       `json:"customer_id,omitempty"`
}

func NewDefaultOrder() *Order {
	return &Order{
		ID:             new(int),
		Subtotal:       new(float64),
		Total:          new(float64),
		GST:            new(float64),
		PST:            new(float64),
		Discount:       new(float64),
		Timestamp:      new(time.Time),
		Void:           new(bool),
		Paid:           new(bool),
		Customizations: new(string),
		Category:       new(string),
	}
}

func NewDefaultOrderWithNil() *Order {
	return &Order{}
}
