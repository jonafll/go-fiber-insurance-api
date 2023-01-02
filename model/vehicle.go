package model

type Vehicle struct {
	Id           int64  `json:"id" validate:"required"`
	Manufacturer string `json:"manufacturer" validate:"required"`
	Model        string `json:"model" validate:"required"`
	Vin          string `json:"vin" validate:"required,min=17,max=17"`
	Performance  int32  `json:"performance" validate:"max=999"`
}
