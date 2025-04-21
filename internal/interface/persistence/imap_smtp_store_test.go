package persistence

import (
	"email-client/internal/domain/model"
	"testing"
)

// --- FakeStore for mocking ListEmails and GetEmail ---

type fakeStore struct {
	ImapSmtpStore
	mockEmails []model.Email
}

func (f *fakeStore) ListEmails() []model.Email {
	return f.mockEmails
}

func (f *fakeStore) GetEmail(id string) *model.Email {
	for _, email := range f.ListEmails() {
		if email.ID == id {
			return &email
		}
	}
	return nil
}

// --- GetEmail Tests ---

func TestGetEmail(t *testing.T) {
	email := model.Email{ID: "123", Subject: "Mock Email"}
	store := &fakeStore{
		mockEmails: []model.Email{email},
	}

	result := store.GetEmail("123")
	if result == nil {
		t.Fatal("expected to find email, got nil")
	}
	if result.Subject != "Mock Email" {
		t.Errorf("unexpected subject: got %q, want %q", result.Subject, "Mock Email")
	}
}

func TestGetEmail_NotFound(t *testing.T) {
	store := &fakeStore{mockEmails: []model.Email{}}
	result := store.GetEmail("not-found")
	if result != nil {
		t.Errorf("expected nil for missing email, got %+v", result)
	}
}
