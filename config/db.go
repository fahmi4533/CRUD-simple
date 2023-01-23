package config

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func DbConnection() (*sql.DB, error) {
	dbDriver := "mysql"
	dbUser := "root"
	dbpass := "12345678"
	dbName := "crud_terbaru"

	db, err := sql.Open(dbDriver, dbUser+":"+dbpass+"@/"+dbName)
	return db, err
}
