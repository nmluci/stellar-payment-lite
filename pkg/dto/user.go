package dto

type UserQueryParams struct {
	UserID int64 `param:"UserID"`
	Limit  int64 `json:"limit"`
	Offset int64 `json:"offset"`
}

type UserPayload struct {
	Password string `json:"password"`
}

type UserRegistrationPayload struct {
	LegalName  string `json:"legal_name"`
	Address    string `json:"address"`
	Phone      string `json:"phone"`
	Birthplace string `json:"birthplace"`
	Birthdate  string `json:"birthdate"`
	NIK        string `json:"nik"`
	Occupation string `json:"occupation"`
	KTPUrl     string `json:"ktp_url"`
	Username   string `json:"username"`
	Password   string `json:"password"`
}

type UserResponse struct {
	UserID     int64  `json:"user_id"`
	CustomerID int64  `json:"customer_id"`
	Username   string `json:"username"`
	RoleID     int64  `json:"role_id"`
}
