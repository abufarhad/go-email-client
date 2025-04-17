package controller

import (
	"email-client/internal/backend"
	"email-client/internal/model"
)

type Handler struct {
	store *backend.Store
}

func NewHandler(store *backend.Store) *Handler {
	return &Handler{store: store}
}

func (h *Handler) GetInbox() []model.Email {
	return h.store.ListEmails()
}

func (h *Handler) GetEmail(id string) *model.Email {
	return h.store.GetEmail(id)
}

func (h *Handler) Send(email model.Email) {
	h.store.SendEmail(email)
}
