package mocks

import (
	"email-client/internal/domain/model"
)

type EmailRepositoryMock struct {
	Emails []model.Email
}

func (m *EmailRepositoryMock) ListEmails() []model.Email {
	return m.Emails
}

func (m *EmailRepositoryMock) GetEmail(id string) *model.Email {
	for _, e := range m.Emails {
		if e.ID == id {
			return &e
		}
	}
	return nil
}

func (m *EmailRepositoryMock) SaveEmail(email model.Email) {
	m.Emails = append([]model.Email{email}, m.Emails...)
}

func (m *EmailRepositoryMock) DeleteEmail(id string) {
	for i, e := range m.Emails {
		if e.ID == id {
			m.Emails = append(m.Emails[:i], m.Emails[i+1:]...)
			break
		}
	}
}
