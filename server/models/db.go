package models

import (
	"log"
	"os"
	"strconv"
)

type ItemInput struct {
	ID                 *int     `json:"id"`
	MenuID             *int     `json:"menu_id,omitempty"`
	CategoryID         *int     `json:"category_id,omitempty"`
	Price              *float64 `json:"price,omitempty"`
	NameEng            *string  `json:"name_eng,omitempty"`
	NameOth            *string  `json:"name_oth,omitempty"`
	Special            *bool    `json:"special,omitempty"`
	Alcohol            *bool    `json:"alcohol,omitempty"`
	Custom             *bool    `json:"custom,omitempty"`
	Variant            *string  `json:"variant,omitempty"`
	VariantDefault     *bool    `json:"variant_default,omitempty"`
	VariantPriceCharge *float64 `json:"variant_price_charge,omitempty"`
}

type Item struct {
	ID                 int     `json:"id"`
	MenuID             int     `json:"menu_id"`
	CategoryID         int     `json:"category_id"`
	Price              float64 `json:"price"`
	NameEng            string  `json:"name_eng"`
	NameOth            string  `json:"name_oth"`
	Special            bool    `json:"special"`
	Alcohol            bool    `json:"alcohol"`
	Custom             bool    `json:"custom"`
	Variant            string  `json:"variant"`
	VariantDefault     bool    `json:"variant_default"`
	VariantPriceCharge float64 `json:"variant_price_charge"`
}

func GetDefaultCategoryID() int {
	var defaultCategoryID = os.Getenv("DEFAULT_ITEM_CATEGORY_ID")
	if defaultCategoryID == "" {
		defaultCategoryID = "1"
	}
	defaultCategoryIDInt, err := strconv.Atoi(defaultCategoryID)
	if err != nil {
		log.Fatalf("Failed to convert DEFAULT_ITEM_CATEGORY_ID to integer: %v", err)
	}
	return defaultCategoryIDInt
}

func NewDefaultItem() *Item {
	return &Item{
		ID:                 *new(int),
		MenuID:             *new(int),
		CategoryID:         GetDefaultCategoryID(),
		Price:              *new(float64),
		NameEng:            *new(string),
		NameOth:            *new(string),
		Special:            *new(bool),
		Alcohol:            *new(bool),
		Custom:             *new(bool),
		Variant:            *new(string),
		VariantDefault:     *new(bool),
		VariantPriceCharge: *new(float64),
	}
}

func NewDefaultItemInputWithNil() *ItemInput {
	return &ItemInput{}
}

func NewDefaultItemInput() *ItemInput {
	var defaultCategoryID = GetDefaultCategoryID()
	return &ItemInput{
		ID:                 new(int),
		MenuID:             new(int),
		CategoryID:         &defaultCategoryID,
		Price:              new(float64),
		NameEng:            new(string),
		NameOth:            new(string),
		Special:            new(bool),
		Alcohol:            new(bool),
		Custom:             new(bool),
		Variant:            new(string),
		VariantDefault:     new(bool),
		VariantPriceCharge: new(float64),
	}
}
