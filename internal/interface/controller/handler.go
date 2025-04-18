package controller

import (
	"email-client/internal/domain/model"
	"email-client/internal/domain/service"
)

type Handler struct {
	emailService *service.EmailService
}

func NewHandler(s *service.EmailService) *Handler {
	return &Handler{emailService: s}
}

func (h *Handler) GetInbox() []model.Email {
	return h.emailService.Inbox()
}

func (h *Handler) GetEmail(id string) *model.Email {
	return h.emailService.ReadEmail(id)
}

func (h *Handler) Send(email model.Email) {
	h.emailService.Send(email)
}

func (h *Handler) Delete(id string) {
	h.emailService.Delete(id)
}
