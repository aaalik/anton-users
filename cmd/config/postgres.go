package config

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/jmoiron/sqlx/reflectx"
	_ "github.com/lib/pq"
)

func (cf Config) NewPostgres() (reader, writer *sqlx.DB) {
	reader, err := sqlx.Connect(
		cf.SQLDB.Driver,
		fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
			cf.SQLDB.Read.Host,
			cf.SQLDB.Read.Port,
			cf.SQLDB.Read.User,
			cf.SQLDB.Read.Pass,
			cf.SQLDB.Read.Name,
		),
	)
	if err != nil {
		panic(err)
	}

	err = reader.Ping()
	if err != nil {
		panic(err)
	}

	reader.Mapper = reflectx.NewMapper("json")
	reader.SetMaxIdleConns(0)

	writer, err = sqlx.Connect(
		cf.SQLDB.Driver,
		fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
			cf.SQLDB.Write.Host,
			cf.SQLDB.Write.Port,
			cf.SQLDB.Write.User,
			cf.SQLDB.Write.Pass,
			cf.SQLDB.Write.Name,
		),
	)
	if err != nil {
		panic(err)
	}

	err = writer.Ping()
	if err != nil {
		panic(err)
	}

	writer.Mapper = reflectx.NewMapper("json")
	writer.SetMaxIdleConns(0)
	return
}
