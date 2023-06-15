package indto

type UserRole struct {
	UserID     int64  `db:"user_id"`
	RoleID     int64  `db:"role_id"`
	CustomerID int64  `db:"customer_id"`
	Username   string `db:"username"`
	Name       string `db:"name"`
}

type UserDetail struct {
	UserID     int64  `db:"user_id"`
	CustomerID int64  `db:"customer_id"`
	Username   string `db:"username"`
	LegalName  string `db:"legal_name"`
	RoleID     int64  `db:"role_id"`
	Password   string `db:"password"`
}
