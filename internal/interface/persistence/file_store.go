package persistence

import (
	"email-client/internal/domain/model"
	"encoding/json"
	"log"
	"os"
	"sync"
	"time"

	"github.com/google/uuid"
)

type FileStore struct {
	path   string
	emails []model.Email
	mu     sync.Mutex
}

func NewFileStore(path string) *FileStore {
	fs := &FileStore{path: path}
	fs.load()
	return fs
}

func (fs *FileStore) load() {
	if data, err := os.ReadFile(fs.path); err == nil {
		json.Unmarshal(data, &fs.emails)
		log.Println("Loaded emails from file", "count", len(fs.emails))
	} else {
		fs.emails = []model.Email{
			{
				ID: uuid.New().String(), From: "system@client.io", To: "you@example.com",
				Subject: "Welcome", Body: "Hello!", Timestamp: time.Now(), Read: false,
			},
		}
		fs.save()
	}
}

func (fs *FileStore) save() {
	fs.mu.Lock()
	defer fs.mu.Unlock()
	data, _ := json.MarshalIndent(fs.emails, "", "  ")
	os.WriteFile(fs.path, data, 0644)
}

func (fs *FileStore) ListEmails() []model.Email {
	return fs.emails
}

func (fs *FileStore) GetEmail(id string) *model.Email {
	for i := range fs.emails {
		if fs.emails[i].ID == id {
			fs.emails[i].Read = true
			fs.save()
			return &fs.emails[i]
		}
	}
	return nil
}

func (fs *FileStore) DeleteEmail(id string) {
	for i, email := range fs.emails {
		if email.ID == id {
			fs.emails = append(fs.emails[:i], fs.emails[i+1:]...)
			fs.save()
			return
		}
	}
}

func (fs *FileStore) SaveEmail(email model.Email) {
	email.ID = uuid.New().String()
	email.Timestamp = time.Now()
	fs.emails = append([]model.Email{email}, fs.emails...)
	fs.save()
	log.Println("Saved email", "subject", email.Subject, "to", email.To)
}
