package service

import (
	"email-client/internal/domain/model"
	"email-client/internal/domain/service/mocks"
	"reflect"
	"testing"
)

func TestInbox(t *testing.T) {
	mock := &mocks.EmailRepositoryMock{
		Emails: []model.Email{
			{ID: "1", Subject: "Test"},
		},
	}
	svc := NewEmailService(mock)

	got := svc.Inbox()
	if !reflect.DeepEqual(got, mock.Emails) {
		t.Errorf("Inbox() = %v, want %v", got, mock.Emails)
	}
}

func TestReadEmail(t *testing.T) {
	mock := &mocks.EmailRepositoryMock{
		Emails: []model.Email{{ID: "abc123", Subject: "Read me"}},
	}
	service := NewEmailService(mock)

	got := service.ReadEmail("abc123")
	if got == nil {
		t.Fatal("ReadEmail() returned nil, expected an email")
	}
	if got.ID != "abc123" || got.Subject != "Read me" {
		t.Errorf("ReadEmail() returned unexpected result: %+v", got)
	}
}

func TestSendEmail(t *testing.T) {
	mock := &mocks.EmailRepositoryMock{}
	service := NewEmailService(mock)

	email := model.Email{ID: "xyz789", Subject: "Send this"}
	service.Send(email)

	if len(mock.Emails) != 1 {
		t.Fatalf("expected 1 email after Send(), got %d", len(mock.Emails))
	}

	got := mock.Emails[0]
	if got.ID != "xyz789" || got.Subject != "Send this" {
		t.Errorf("Send() stored incorrect email: %+v", got)
	}
}

func TestDeleteEmail(t *testing.T) {
	mock := &mocks.EmailRepositoryMock{
		Emails: []model.Email{
			{ID: "1"}, {ID: "2"}, {ID: "3"},
		},
	}
	service := NewEmailService(mock)

	service.Delete("2")

	if len(mock.Emails) != 2 {
		t.Fatalf("expected 2 emails after Delete(), got %d", len(mock.Emails))
	}

	for _, e := range mock.Emails {
		if e.ID == "2" {
			t.Errorf("Delete() failed to remove email with ID '2'")
		}
	}
}
