package dbConn

import "github.com/jmoiron/sqlx"

type IDBConnector interface {
	GetConnect() *sqlx.DB
}
