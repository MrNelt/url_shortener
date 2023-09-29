package db

import "github.com/jmoiron/sqlx"

type IDB interface {
	GetConnect() *sqlx.DB
	CloseConnect() error
}
