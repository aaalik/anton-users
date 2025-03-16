package jwtconf

import (
	_ "github.com/joho/godotenv/autoload"
)

func (jut *JwtConf) GetSecret() string {
	return jut.secret
}

func (jut *JwtConf) GetSecretRefresh() string {
	return jut.secretRefresh
}

func (jut *JwtConf) GetExpire() int {
	return jut.expires
}
