package db

import (
	// IsDuplicateErr return true if is duplicate error

	"github.com/go-sql-driver/mysql"
	"github.com/jackc/pgconn"
	"github.com/lib/pq"
)

// IsDuplicateErr ...
func IsDuplicateErr(err error) bool {
	if err == nil {
		return false
	}
	mysqlErr, ok := err.(*mysql.MySQLError)
	if ok {
		if mysqlErr.Number == 1062 {
			// solve the duplicate key error.
			return true
		}
	}

	pgErr, ok := err.(*pq.Error)
	if ok {
		if pgErr.Code == "23505" {
			return true
		}
	}

	connPGErr, ok := err.(*pgconn.PgError)
	if ok {
		if connPGErr.Code == "23505" {
			return true
		}
	}

	return false
}
