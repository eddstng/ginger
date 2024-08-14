package models

type Customer struct {
	ID           *int    `json:"id,omitempty"`
	Name         *string `json:"name,omitempty"`
	Phone        *string `json:"phone,omitempty"`
	UnitNumber   *string `json:"unit_number,omitempty"`
	StreetNumber *string `json:"street_number,omitempty"`
	StreetName   *string `json:"street_name,omitempty"`
	BuzzerNumber *string `json:"buzzer_number,omitempty"`
	Note         *string `json:"note,omitempty"`
}

func NewDefaultCustomer() *Customer {
	return &Customer{
		ID:           new(int),
		Name:         new(string),
		Phone:        new(string),
		UnitNumber:   new(string),
		StreetNumber: new(string),
		StreetName:   new(string),
		BuzzerNumber: new(string),
		Note:         new(string),
	}
}

func NewDefaultCustomerWithNil() *Customer {
	return &Customer{}
}
