package dto

type AuthLoginPayload struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
type AuthRefreshTokenPayload struct {
	RT string `json:"refresh_token"`
}

type AuthResponse struct {
	UserID       int64  `json:"user_id"`
	RoleID       int64  `json:"role_id"`
	Username     string `json:"username"`
	Name         string `json:"name"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type AuthlessLoginPayload struct {
	AuthCode string `json:"auth_code"`
}
