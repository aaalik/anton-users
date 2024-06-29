package service

type RequestLogin struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type ResponseLogin struct {
	Token                 string `json:"token"`
	ExpiresAt             int64  `json:"expires_at"`
	RefreshToken          string `json:"refresh_token"`
	RefreshTokenExpiresAt int64  `json:"refresh_token_expires_at"`
}

type RequestRefreshToken struct {
	RefreshToken string `json:"refresh_token"`
}

type ResponseRefreshToken struct {
	Token     string `json:"token"`
	ExpiresAt int64  `json:"expires_at"`
}

type RequestRegister struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Name     string `json:"name"`
}

type ResponseRegister struct {
	Username string `json:"username"`
	Name     string `json:"name"`
}
