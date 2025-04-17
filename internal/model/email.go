package model

import "time"

type Email struct {
	ID        string
	From      string
	To        string
	Subject   string
	Body      string
	Timestamp time.Time
	Read      bool
}
