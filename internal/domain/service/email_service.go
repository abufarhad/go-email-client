package service

import "email-client/internal/domain/model"

//go:generate mockgen -destination=../../mock/email_repo.go -package=mock email-client/internal/domain/service EmailRepository

type EmailRepository interface {
	ListEmails() []model.Email
	GetEmail(id string) *model.Email
	SaveEmail(email model.Email)
	DeleteEmail(id string)
}

type EmailService struct {
	repo EmailRepository
}

func NewEmailService(r EmailRepository) *EmailService {
	return &EmailService{repo: r}
}

func (s *EmailService) Inbox() []model.Email {
	return s.repo.ListEmails()
}

func (s *EmailService) ReadEmail(id string) *model.Email {
	return s.repo.GetEmail(id)
}

func (s *EmailService) Send(email model.Email) {
	s.repo.SaveEmail(email)
}

func (s *EmailService) Delete(id string) {
	s.repo.DeleteEmail(id)
}
