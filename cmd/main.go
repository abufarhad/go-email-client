package main

import (
	"email-client/internal/domain/service"
	"email-client/internal/infra/logger"
	"email-client/internal/interface/controller"
	"email-client/internal/interface/persistence"
	"email-client/internal/interface/ui"
	"email-client/utils"
	"log"
	"os"
)

func main() {
	logger.InitLogger()
	log.Println("🚀 Starting Email Client")

	utils.LoadEnv()

	var emailService *service.EmailService
	useReal := os.Getenv("USE_REAL_EMAIL")

	log.Printf("🧠 USE_REAL_EMAIL = %v", useReal)

	if useReal == "true" {
		log.Println("📡 Using real IMAP/SMTP email backend")

		email := os.Getenv("EMAIL_USER")
		pass := os.Getenv("EMAIL_PASS")

		if email == "" || pass == "" {
			log.Fatal("❌ Missing EMAIL_USER or EMAIL_PASS environment variables")
		}

		log.Printf("📧 Email configured for user: %s", email)

		imapStore := persistence.NewImapSmtpStore(
			os.Getenv("EMAIL_IMAP_HOST"),
			os.Getenv("EMAIL_IMAP_PORT"),
			os.Getenv("EMAIL_SMTP_HOST"),
			os.Getenv("EMAIL_SMTP_PORT"),
			email,
			pass,
		)
		emailService = service.NewEmailService(imapStore)
	} else {
		log.Println("📁 Using local file-based email store: emails.json")
		fileStore := persistence.NewFileStore("emails.json")
		emailService = service.NewEmailService(fileStore)
	}

	handler := controller.NewHandler(emailService)

	log.Println("🖥️  Launching TUI...")
	ui.StartApp(handler)

	log.Println("👋 Email client exited")
}
