package jwtgen

func NewJwtUtils(key string) *JwtUtils {
	return &JwtUtils{signingKey: []byte(key)}
}
