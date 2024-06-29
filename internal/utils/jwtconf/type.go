package jwtconf

type JwtConf struct {
	secret        string
	secretRefresh string
	expires       int
}
