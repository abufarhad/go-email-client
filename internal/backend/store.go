package backend

import (
	"email-client/internal/model"
	"github.com/google/uuid"
	"time"
)

type Store struct {
	emails []model.Email
}

func NewStore() *Store {
	return &Store{
		emails: []model.Email{
			{
				ID:        uuid.New().String(),
				From:      "user1@example.com",
				To:        "you@example.com",
				Subject:   "Welcome!",
				Body:      "Welcome to your new email client.",
				Timestamp: time.Now(),
				Read:      false,
			},
		},
	}
}

func (s *Store) ListEmails() []model.Email {
	return s.emails
}

func (s *Store) GetEmail(id string) *model.Email {
	for _, e := range s.emails {
		if e.ID == id {
			return &e
		}
	}
	return nil
}

func (s *Store) SendEmail(email model.Email) {
	email.ID = uuid.New().String()
	email.Timestamp = time.Now()
	s.emails = append([]model.Email{email}, s.emails...)
}
