package domain

import (
	"time"
)

type Request struct {
	ID   int
	Date_request time.Time
	Id_user int
	Id_status int
}