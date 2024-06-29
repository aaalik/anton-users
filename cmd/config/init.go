package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"

	cons "github.com/aaalik/anton-users/internal/constant"
)

func NewConfig() Config {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	stopTimeout, err := strconv.Atoi(os.Getenv(cons.ConfigStopTimeout))
	if err != nil {
		panic(err)
	}

	jwtExpire, err := strconv.Atoi(os.Getenv(cons.ConfigJwtExpire))
	if err != nil {
		panic(err)
	}

	return Config{
		Host: Host{
			Address: os.Getenv(cons.ConfigHostAddress),
		},
		StopTimeout: stopTimeout,
		SQLDB: SQLDB{
			Driver: os.Getenv(cons.ConfigSQLDBDriver),
			Write: SQLDBConfig{
				Host: os.Getenv(cons.ConfigSQLDBWriteHost),
				Port: os.Getenv(cons.ConfigSQLDBWritePort),
				User: os.Getenv(cons.ConfigSQLDBWriteUser),
				Pass: os.Getenv(cons.ConfigSQLDBWritePass),
				Name: os.Getenv(cons.ConfigSQLDBWriteName),
			},
			Read: SQLDBConfig{
				Host: os.Getenv(cons.ConfigSQLDBReadHost),
				Port: os.Getenv(cons.ConfigSQLDBReadPort),
				User: os.Getenv(cons.ConfigSQLDBReadUser),
				Pass: os.Getenv(cons.ConfigSQLDBReadPass),
				Name: os.Getenv(cons.ConfigSQLDBReadName),
			},
		},
		JwtConf: JwtConf{
			Secret:        os.Getenv(cons.ConfigJwtSecret),
			SecretRefresh: os.Getenv(cons.ConfigJwtSecretRefresh),
			Expire:        jwtExpire,
		},
	}
}
