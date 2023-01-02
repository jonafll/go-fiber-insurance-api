package model

type Tariff struct {
	InsuranceStart string  `json:"insuranceStart" validate:"required,datetime=2006-01-02,futuredate"`
	Coverage       string  `json:"coverage" validate:"required,oneof=liability comprehensive partial"`
	Vehicle        Vehicle `json:"vehicle" validate:"required"`
	Person         Person  `json:"person" validate:"required"`
}

type TariffAmount struct {
	Total         float64 `json:"total"`
	Partial       float64 `json:"partial"`
	Comprehensive float64 `json:"comprehensive"`
	Liability     float64 `json:"liability"`
	Currency      string  `json:"currency" validate:"required,format=iso4217"`
}
