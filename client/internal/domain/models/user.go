package models

type User struct {
	Id        uint    `json:"id"`
	FirstName string  `json:"first_name"`
	City      string  `json:"city"`
	Height    float32 `json:"height"`
	Married   bool    `json:"married"`
	Phone     string  `json:"phone"`
}
