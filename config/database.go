package config

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/tanabebe/scraping-dermatology/domain"
)

// Databaseとの接続を行う
func ConnectDb(cfg domain.DbConfig) (Db *sql.DB, err error) {
	dsn := fmt.Sprintf("host=/cloudsql/%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.InstanceConnectionName,
		cfg.DatabaseUser,
		cfg.Password,
		cfg.DatabaseName)

	Db, err = sql.Open(cfg.DriverName, dsn)

	if err != nil {
		return nil, err
	}

	return Db, nil
}
