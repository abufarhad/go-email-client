package persistence

import (
	"email-client/internal/domain/model"
	"github.com/emersion/go-imap"
	"github.com/emersion/go-imap/client"
	"github.com/jhillyerd/enmime"
	"log"
	"net/smtp"
)

type ImapSmtpStore struct {
	imapHost string
	smtpHost string
	smtpPort string
	username string
	password string
}

func NewImapSmtpStore(imapHost, smtpHost, smtpPort, username, password string) *ImapSmtpStore {
	return &ImapSmtpStore{
		imapHost: imapHost,
		smtpHost: smtpHost,
		smtpPort: smtpPort,
		username: username,
		password: password,
	}
}

// SaveEmail implements EmailRepository (sends email)
func (s *ImapSmtpStore) SaveEmail(email model.Email) {
	auth := smtp.PlainAuth("", s.username, s.password, s.smtpHost)
	fullAddr := s.smtpHost + ":" + s.smtpPort

	msg := []byte("To: " + email.To + "\r\n" +
		"Subject: " + email.Subject + "\r\n" +
		"\r\n" + email.Body + "\r\n")

	//log.Printf("📤 Sending email to %s via %s as %s", email.To, fullAddr, s.username)
	err := smtp.SendMail(fullAddr, auth, s.username, []string{email.To}, msg)
	if err != nil {
		log.Printf("❌ Send failed: %v", err)
	} else {
		log.Println("✅ Email sent!")
	}
}

// ListEmails fetches the inbox and returns full plain-text emails
func (s *ImapSmtpStore) ListEmails() []model.Email {
	var emails []model.Email

	c, err := client.DialTLS(s.imapHost+":993", nil)
	if err != nil {
		log.Println("❌ IMAP dial error:", err)
		return emails
	}
	defer c.Logout()

	if err := c.Login(s.username, s.password); err != nil {
		log.Println("❌ IMAP login error:", err)
		return emails
	}

	mbox, err := c.Select("INBOX", false)
	if err != nil {
		log.Println("❌ Inbox select error:", err)
		return emails
	}

	// Fetch last 10 emails
	from := uint32(1)
	if mbox.Messages > 10 {
		from = mbox.Messages - 9
	}
	seqSet := new(imap.SeqSet)
	seqSet.AddRange(from, mbox.Messages)

	section := &imap.BodySectionName{}
	items := []imap.FetchItem{imap.FetchEnvelope, section.FetchItem()}
	messages := make(chan *imap.Message, 10)

	go func() {
		if err := c.Fetch(seqSet, items, messages); err != nil {
			log.Println("❌ Fetch error:", err)
		}
	}()

	for msg := range messages {
		if msg == nil {
			continue
		}

		r := msg.GetBody(section)
		if r == nil {
			log.Println("⚠️ No message body")
			continue
		}

		env, err := enmime.ReadEnvelope(r)
		body := ""
		if err != nil {
			log.Println("⚠️ Failed to parse MIME body:", err)
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

	log.Printf("📥 Fetched %d emails", len(emails))
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
	log.Printf("🗑️ Delete not implemented for IMAP (skipped ID: %s)", id)
}

//package persistence
//
//import (
//	"email-client/internal/domain/model"
//	"log"
//	"net/smtp"
//
//	"github.com/emersion/go-imap"
//	"github.com/emersion/go-imap/client"
//)
//
//type ImapSmtpStore struct {
//	imapHost string
//	imapPort string
//	smtpHost string
//	smtpPort string
//	username string
//	password string
//}
//
//func NewImapSmtpStore(imapHost, imapPort, smtpHost, smtpPort, username, password string) *ImapSmtpStore {
//	return &ImapSmtpStore{
//		imapHost: imapHost,
//		imapPort: imapPort,
//		smtpHost: smtpHost,
//		smtpPort: smtpPort,
//		username: username,
//		password: password,
//	}
//}
//
//// ListEmails implements EmailRepository
//func (s *ImapSmtpStore) ListEmails() []model.Email {
//	var emails []model.Email
//
//	c, err := client.DialTLS(s.imapHost+":993", nil)
//	if err != nil {
//		log.Println("❌ IMAP dial error:", err)
//		return emails
//	}
//	defer c.Logout()
//
//	if err := c.Login(s.username, s.password); err != nil {
//		log.Println("❌ IMAP login error:", err)
//		return emails
//	}
//
//	mbox, err := c.Select("INBOX", false)
//	if err != nil {
//		log.Println("❌ Inbox select error:", err)
//		return emails
//	}
//
//	from := uint32(1)
//	if mbox.Messages > 10 {
//		from = mbox.Messages - 9
//	}
//	seqSet := new(imap.SeqSet)
//	seqSet.AddRange(from, mbox.Messages)
//
//	messages := make(chan *imap.Message, 10)
//	err = c.Fetch(seqSet, []imap.FetchItem{imap.FetchEnvelope}, messages)
//	if err != nil {
//		log.Println("❌ Fetch error:", err)
//		return emails
//	}
//
//	for msg := range messages {
//		emails = append(emails, model.Email{
//			ID:      msg.Envelope.MessageId,
//			From:    msg.Envelope.From[0].Address(),
//			To:      msg.Envelope.To[0].Address(),
//			Subject: msg.Envelope.Subject,
//			Body:    "", // Body fetching not implemented here
//		})
//	}
//
//	return emails
//}
//
//// GetEmail implements EmailRepository (simplified: scans inbox)
//func (s *ImapSmtpStore) GetEmail(id string) *model.Email {
//	for _, email := range s.ListEmails() {
//		if email.ID == id {
//			return &email
//		}
//	}
//	return nil
//}
//
//// SaveEmail implements EmailRepository (sends email)
//func (s *ImapSmtpStore) SaveEmail(email model.Email) {
//	auth := smtp.PlainAuth("", s.username, s.password, s.smtpHost)
//
//	fullAddr := s.smtpHost + ":" + s.smtpPort
//
//	msg := []byte("To: " + email.To + "\r\n" +
//		"Subject: " + email.Subject + "\r\n" +
//		"\r\n" + email.Body + "\r\n")
//
//	err := smtp.SendMail(fullAddr, auth, s.username, []string{email.To}, msg)
//
//	if err != nil {
//		log.Printf("❌ Send failed: %v", err)
//	} else {
//		log.Println("✅ Email sent!")
//	}
//}
//
//// DeleteEmail implements EmailRepository (noop for now)
//func (s *ImapSmtpStore) DeleteEmail(id string) {
//	log.Printf("🗑️  Delete not implemented for IMAP store (skipped ID: %s)", id)
//}
