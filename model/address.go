package model

type Address struct {
	City       string `json:"city" validate:"required"`
	Street     string `json:"street" validate:"required"`
	Country    string `json:"country" validate:"required,iso3166_1_alpha2"`
	PostalCode int32  `json:"postalCode" validate:"required"`
}
