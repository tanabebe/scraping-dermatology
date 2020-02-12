package repository

import (
	"database/sql"
	"log"
)

func CountByWaitReservation(Db *sql.DB) int {
	rows, err := Db.Query("SELECT COUNT(*) FROM wait_reservation")
	defer rows.Close()
	rows.Next()
	var count int
	err = rows.Scan(&count)
	if err != nil {
		log.Fatal(err)
	}
	return count
}
