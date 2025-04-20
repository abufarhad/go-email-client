package persistence

import (
	"email-client/internal/domain/model"
	"github.com/emersion/go-imap"
	"github.com/emersion/go-imap/client"
	"github.com/jhillyerd/enmime"
	"log"
	"net/smtp"
	"os"
	"strconv"
)

type ImapSmtpStore struct {
	imapHost string
	imapPort string
	smtpHost string
	smtpPort string
	username string
	password string
}

func NewImapSmtpStore(imapHost, imapPort, smtpHost, smtpPort, username, password string) *ImapSmtpStore {
	return &ImapSmtpStore{
		imapHost: imapHost,
		imapPort: imapPort,
		smtpHost: smtpHost,
		smtpPort: smtpPort,
		username: username,
		password: password,
	}
}

func (s *ImapSmtpStore) SaveEmail(email model.Email) {
	auth := smtp.PlainAuth("", s.username, s.password, s.smtpHost)
	fullAddr := s.smtpHost + ":" + s.smtpPort

	msg := []byte("To: " + email.To + "\r\n" +
		"Subject: " + email.Subject + "\r\n" +
		"\r\n" + email.Body + "\r\n")

	err := smtp.SendMail(fullAddr, auth, s.username, []string{email.To}, msg)
	if err != nil {
		log.Printf("Send failed: %v", err)
	} else {
		log.Println("Email sent successfully")
	}
}

func (s *ImapSmtpStore) ListEmails() []model.Email {
	var emails []model.Email

	c, err := client.DialTLS(s.imapHost+":"+s.imapPort, nil)
	if err != nil {
		log.Println("IMAP dial error:", err)
		return emails
	}
	defer c.Logout()

	if err := c.Login(s.username, s.password); err != nil {
		log.Println("IMAP login error:", err)
		return emails
	}

	mbox, err := c.Select("INBOX", false)
	if err != nil {
		log.Println("Inbox select error:", err)
		return emails
	}

	EmailFetchNumber := os.Getenv("NUMBER_OF_EMAIL_TO_FETCH")
	numberOfEmailToFetch, err := strconv.Atoi(EmailFetchNumber)
	if err != nil || numberOfEmailToFetch < 1 {
		log.Printf("Invalid NUMBER_OF_EMAIL_TO_FETCH (%s), defaulting to 5", EmailFetchNumber)
		numberOfEmailToFetch = 5
	}

	log.Printf("Fetching last %d emails", numberOfEmailToFetch)

	from := uint32(1)
	if mbox.Messages > uint32(numberOfEmailToFetch) {
		from = mbox.Messages - uint32(numberOfEmailToFetch) + 1
	}

	seqSet := new(imap.SeqSet)
	seqSet.AddRange(from, mbox.Messages)

	section := &imap.BodySectionName{}
	items := []imap.FetchItem{imap.FetchEnvelope, section.FetchItem()}
	messages := make(chan *imap.Message, numberOfEmailToFetch)

	go func() {
		if err := c.Fetch(seqSet, items, messages); err != nil {
			log.Println("Fetch error:", err)
		}
	}()

	for msg := range messages {
		if msg == nil {
			continue
		}

		r := msg.GetBody(section)
		if r == nil {
			log.Println("Message body missing")
			continue
		}

		env, err := enmime.ReadEnvelope(r)
		body := ""
		if err != nil {
			log.Println("Failed to parse MIME body:", err)
		} else {
			body = env.Text
			if body == "" {
				body = "(No text body found)"
			}
		}

		emails = append(emails, model.Email{
			ID:      msg.Envelope.MessageId,
			From:    msg.Envelope.From[0].Address(),
			To:      msg.Envelope.To[0].Address(),
			Subject: msg.Envelope.Subject,
			Body:    body,
		})
	}

	log.Printf("Fetched %d emails", len(emails))
	return emails
}

func (s *ImapSmtpStore) GetEmail(id string) *model.Email {
	for _, email := range s.ListEmails() {
		if email.ID == id {
			return &email
		}
	}
	return nil
}

func (s *ImapSmtpStore) DeleteEmail(id string) {
	log.Printf("Delete not implemented for IMAP store. Email ID: %s", id)
}
