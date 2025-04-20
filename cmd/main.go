package main

import (
	"email-client/internal/domain/service"
	"email-client/internal/infra/logger"
	"email-client/internal/interface/controller"
	"email-client/internal/interface/persistence"
	"email-client/internal/interface/ui"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	logger.InitLogger()
	log.Println("🚀 Starting Email Client")

	// Load environment variables from .env
	err := godotenv.Load()
	if err != nil {
		log.Println("⚠️  No .env file found, falling back to system env")
	}

	var emailService *service.EmailService
	useReal := os.Getenv("USE_REAL_EMAIL") == "true"

	log.Printf("🧠 USE_REAL_EMAIL = %v", useReal)

	if useReal {
		log.Println("📡 Using real IMAP/SMTP email backend")

		email := os.Getenv("EMAIL_USER")
		pass := os.Getenv("EMAIL_PASS")

		if email == "" || pass == "" {
			log.Fatal("❌ Missing EMAIL_USER or EMAIL_PASS environment variables")
		}

		log.Printf("📧 Email configured for user: %s", email)

		imapStore := persistence.NewImapSmtpStore(
			os.Getenv("EMAIL_IMAP_HOST"),
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
