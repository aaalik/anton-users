package jwtconf

func NewJwtConf(secret, secretRefresh string, expire int) *JwtConf {
	return &JwtConf{secret, secretRefresh, expire}
}
