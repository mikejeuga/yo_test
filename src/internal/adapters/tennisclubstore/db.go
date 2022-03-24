package tennisclubstore

import (
	"database/sql"
)

type DB struct {
	config DBConfig
	clubDB *sql.DB
}

func createDB(config DBConfig, connectionString string) (db *sql.DB, dbName string, err error) {
	return nil, "", nil
}
