package model

import (
	"sync"
	"time"
)

type Email struct {
	ID        string
	From      string
	To        string
	Subject   string
	Body      string
	Timestamp time.Time
	Read      bool
}

type FileStore struct {
	path   string
	emails []Email
	mu     sync.Mutex
}
