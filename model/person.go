package model

type Person struct {
	FirstName   string  `json:"firstName" validate:"required"`
	LastName    string  `json:"lastName" validate:"required"`
	DateOfBirth string  `json:"dateOfBirth" validate:"required,datetime=2006-01-02"`
	Email       string  `json:"email" validate:"required,email"`
	Phone       string  `json:"phone" validate:"e164"`
	Address     Address `json:"address" validate:"required"`
}
