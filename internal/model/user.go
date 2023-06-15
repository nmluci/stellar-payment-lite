package model

import "time"

type User struct {
	UserID     int64      `db:"id"`
	CustomerID int64      `db:"customer_id"`
	RoleID     int64      `db:"role_id"`
	Username   string     `db:"username"`
	Password   string     `db:"password"`
	CreatedAt  time.Time  `db:"created_at"`
	UpdatedAt  time.Time  `db:"updated_at"`
	DeletedAt  *time.Time `db:"deleted_at"`
}
