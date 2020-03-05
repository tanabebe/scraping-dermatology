package domain

import "time"

type WaitReservation struct {
	Id              int       `json:id`
	IsReservation   bool      `json:is_reservation,omitempty`
	ReservationDate time.Time `json:reservation_date,omitempty`
}
