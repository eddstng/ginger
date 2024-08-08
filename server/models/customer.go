package models

type Customer struct {
	ID           *int    `json:"id"`
	Name         *string `json:"name"`
	Phone        *string `json:"phone"`
	UnitNumber   *string `json:"unit_number"`
	StreetNumber *string `json:"street_number"`
	StreetName   *string `json:"street_name"`
	BuzzerNumber *string `json:"buzzer_number"`
	Note         *string `json:"note"`
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
