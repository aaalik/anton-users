package database

import "github.com/jmoiron/sqlx"

func NewDB(reader, writer *sqlx.DB) *DB {
	return &DB{Reader: reader, Writer: writer}
}
