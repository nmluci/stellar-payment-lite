package dto

type CustomerQueryParams struct {
	CustomerID int64 `param:"customerID"`
}

type CustomerPayload struct {
	LegalName  string `json:"legal_name"`
	Address    string `json:"address"`
	Phone      string `json:"phone"`
	Birthplace string `json:"birthplace"`
	Birthdate  string `json:"birthdate"`
	NIK        string `json:"nik"`
	Occupation string `json:"occupation"`
	KTPUrl     string `json:"ktp_url"`
}

type CustomerResponse struct {
	CustomerID int64  `json:"customer_id"`
	LegalName  string `json:"legal_name"`
	Address    string `json:"address"`
	Phone      string `json:"phone"`
	Birthplace string `json:"birthplace"`
	Birthdate  string `json:"birthdate"`
	NIK        string `json:"nik"`
	Occupation string `json:"occupation"`
	KTPUrl     string `json:"ktp_url"`
}
