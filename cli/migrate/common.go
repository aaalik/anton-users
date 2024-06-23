package migrate

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/joho/godotenv/autoload"
	_ "github.com/lib/pq"

	cons "github.com/aaalik/anton-users/internal/constant"
)

func setupMigrateDB() (*migrate.Migrate, error) {
	db, err := sql.Open(
		os.Getenv(cons.ConfigSQLDBDriver),
		fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
			os.Getenv(cons.ConfigSQLDBWriteHost),
			os.Getenv(cons.ConfigSQLDBWritePort),
			os.Getenv(cons.ConfigSQLDBWriteUser),
			os.Getenv(cons.ConfigSQLDBWritePass),
			os.Getenv(cons.ConfigSQLDBWriteName),
		),
	)
	if err != nil {
		return nil, err
	}

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		return nil, err
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://migrations",
		os.Getenv(cons.ConfigSQLDBWriteName), driver)
	if err != nil {
		return nil, err
	}

	return m, nil
}
