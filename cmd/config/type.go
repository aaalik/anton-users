package config

type Config struct {
	Host        Host
	StopTimeout int
	SQLDB       SQLDB
	JwtConf     JwtConf
}

type Host struct {
	Address string
}

type HttpService struct {
	URL string
}

type SQLDB struct {
	Driver string
	Write  SQLDBConfig
	Read   SQLDBConfig
}

type SQLDBConfig struct {
	Host string
	Port string
	User string
	Pass string
	Name string
}

type JwtConf struct {
	Secret        string
	SecretRefresh string
	Expire        int
}
