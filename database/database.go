package database

import (
	"database/sql"
	"fmt"

	_ "github.com/jackc/pgx/v5/stdlib"
)

var DB *sql.DB

func DSN(host, port, user, pass, db, ssl string) string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s", user, pass, host, port, db, ssl)
}

func Init(dsn string) error {
	var err error
	DB, err = sql.Open("pgx", dsn)
	if err != nil {
		return err
	}
	return DB.Ping()
}
