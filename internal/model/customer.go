package model

import "time"

type Customer struct {
	ID         int64     `db:"id"`
	LegalName  string    `db:"legal_name"`
	Address    string    `db:"address"`
	Phone      string    `db:"phone"`
	Birthplace string    `db:"birthplace"`
	Birthdate  time.Time `db:"birthdate"`
	NIK        string    `db:"nik"`
	Occupation string    `db:"occupation"`
	KTPUrl     string    `db:"ktp_url"`
}

type Account struct {
	ID          int64   `db:"id"`
	CustomerID  int64   `db:"customer_id"`
	AccountNo   string  `db:"account_no"`
	AccountType int64   `db:"account_type"`
	CardNumber  string  `db:"card_number"`
	CVV         string  `db:"cvv"`
	PIN         string  `db:"pin"`
	Balance     float32 `db:"balance"`
}
