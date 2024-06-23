package config

type Config struct {
	Host        Host
	StopTimeout int
	SQLDB       SQLDB
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
