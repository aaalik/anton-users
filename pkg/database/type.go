package database

import "github.com/jmoiron/sqlx"

type DB struct {
	Reader *sqlx.DB
	Writer *sqlx.DB
}
