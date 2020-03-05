package repository

import (
	"database/sql"
	"log"

	"github.com/tanabebe/scraping-dermatology/domain"
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

func CreateReservation(Db *sql.DB, reservation *domain.WaitReservation) error {
	statement := "INSERT INTO wait_reservation (is_reservation, reservation_date) VALUES ($1, $2) returning id"
	stmt, err := Db.Prepare(statement)
	defer stmt.Close()
	if err != nil {
		return err
	}
	err = stmt.QueryRow(reservation.IsReservation, reservation.ReservationDate).Scan(&reservation.Id)
	if err != nil {
		return err
	}
	return nil
}
