package persistence

import (
	"email-client/internal/domain/model"
	"os"
	"testing"
)

func createTempFileStore(t *testing.T) (*FileStore, func()) {
	t.Helper()

	tmpFile, err := os.CreateTemp("", "emails_test_*.json")
	if err != nil {
		t.Fatalf("failed to create temp file: %v", err)
	}

	// Clean up function to delete the file after test
	cleanup := func() {
		os.Remove(tmpFile.Name())
	}

	store := NewFileStore(tmpFile.Name())
	return store, cleanup
}

func TestSaveAndListEmails(t *testing.T) {
	store, cleanup := createTempFileStore(t)
	defer cleanup()

	email := model.Email{
		From:    "a@b.com",
		To:      "c@d.com",
		Subject: "Test Subject",
		Body:    "Test Body",
	}

	store.SaveEmail(email)

	emails := store.ListEmails()
	if len(emails) == 0 {
		t.Fatal("expected at least 1 email, got 0")
	}

	if emails[0].Subject != "Test Subject" {
		t.Errorf("expected subject 'Test Subject', got '%s'", emails[0].Subject)
	}
}

func TestGetEmailMarksAsRead(t *testing.T) {
	store, cleanup := createTempFileStore(t)
	defer cleanup()

	email := model.Email{
		From:    "x@y.com",
		To:      "z@a.com",
		Subject: "Mark Read Test",
		Body:    "Body",
	}

	store.SaveEmail(email)
	saved := store.ListEmails()[0]

	fetched := store.GetEmail(saved.ID)
	if fetched == nil {
		t.Fatal("expected to fetch saved email, got nil")
	}
	if !fetched.Read {
		t.Error("expected email to be marked as read")
	}
}

func TestDeleteEmail(t *testing.T) {
	store, cleanup := createTempFileStore(t)
	defer cleanup()

	email := model.Email{
		From:    "foo@bar.com",
		To:      "baz@qux.com",
		Subject: "Delete Me",
		Body:    "This will be deleted",
	}

	store.SaveEmail(email)
	saved := store.ListEmails()[0]

	store.DeleteEmail(saved.ID)

	emails := store.ListEmails()
	for _, e := range emails {
		if e.ID == saved.ID {
			t.Error("email was not deleted")
		}
	}
}
