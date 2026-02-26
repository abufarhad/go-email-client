# 📬 Terminal Email Client [ GoMail ]

A fast and minimal terminal-based email client built in Go. Inspired by tools like `mutt`, `aerc`, and `himalaya`.

This project demonstrates clean architecture, form-based input, SMTP/IMAP support, and async operations using Go + TUI (`tview`).

[🚀 Live Demo](https://go-email-client-production.up.railway.app)

---
## 📹 Video Preview

[![Watch the video](https://img.icons8.com/ios-filled/50/000000/video.png)](https://github.com/user-attachments/assets/112b6a42-9c0e-43f9-b99c-b4ea4b587b8e)

### 🎥 Watch the Full Demo

[🚀 YouTube Video](https://youtu.be/e9KuynLewJQ)

---

## ✨ Features

- 📥 View Inbox (via file store or real IMAP)
- 📝 Compose & send real emails via SMTP
- 🚗 Dual-mode backend: file or real email provider (Gmail, Outlook, etc.)
- 🗑️ Delete email (file-based only)
- ❌ ESC/back navigation from all views
- ✅ Email validation and success messages
- 📃 Logs written to `logs.txt` (resets each run)
- 🐳 Docker support
- 🧼 Clean, modular code structure

---

## 🔌 SMTP/IMAP Integration

Enable real email capabilities with Gmail, Outlook, Fastmail, etc. (via App Passwords or standard login).

### Setup `.env`:

```env
USE_REAL_EMAIL=true 
EMAIL_IMAP_HOST=imap.gmail.com
EMAIL_SMTP_HOST=smtp.gmail.com
EMAIL_SMTP_PORT=587
EMAIL_IMAP_PORT=993
EMAIL_USER=your@email.com
EMAIL_PASS=your-app-password

NUMBER_OF_EMAIL_TO_FETCH=5
```

> ⚠️ Use App Password for Gmail (NOT your real password!)

---

## 📊 Project Structure

```
go-email-client/
├── cmd/
│   └── main.go                       # Application entry point (CLI)
├── internal/
│   ├── domain/
│   │   ├── model/
│   │   │   └── email.go              # Core domain model for emails
│   │   └── service/
│   │       └── email_service.go      # Business logic for email operations
│   ├── infra/
│   │   └── logger/
│   │       └── logger.go             # Centralized logging configuration
│   ├── interface/
│   │   ├── controller/
│   │   │   └── handler.go            # Application-layer request handling
│   │   ├── persistence/
│   │   │   ├── file_store.go         # File-based storage implementation
│   │   │   └── imap_smtp_store.go    # IMAP/SMTP backend implementation
│   │   └── ui/
│   │       └── app.go                # Terminal UI (TUI) using tview
│   └── utils/
│       └── utils.go                  # Shared utility functions
├── web/
│   ├── static/
│   │   └── index.html                # Web-based terminal (xterm.js)
│   └── main.go                       # WebSocket and PTY server
├── .env.example                      # Sample environment configuration
├── .gitignore                        # Git ignored files and directories
├── Dockerfile                        # Docker multi-stage build config
├── emails.json                       # Sample local email data store
├── go.mod                            # Go module definition
├── logs.txt                          # Log output file
├── Makefile                          # Build and development commands
└── README.md                         # Project documentation
```

---

## 🛠️ Requirements

- Go 1.21+
- Docker (optional)
- `make` (for simplified workflows)

---

## 🚀 Getting Started

### 🔧 Local Build (No Docker)

🔹 Run Terminal-Only App
```bash
go run cmd/main.go              # terminal-only UI
```

🔹 Run Web-Based Terminal UI
```bash
go build -o email-client ./cmd  # build CLI binary
go run ./web                    # run web interface
```

🔹 Or Use the Makefile (Recommended)
```bash
make build && make run
```

Visit: [http://localhost:8080](http://localhost:8080)

### 🐳 Docker Workflow

```bash
make docker-build   # builds WebSocket + CLI
make docker-run     # launches container on port 8080
make clean          # removes built binaries
```

Visit: [http://localhost:8080](http://localhost:8080)

---

## 👨‍💻 Controls

- `i` → 📥 Open Inbox
- `c` → 📝 Compose Email
- `d` → 🗑️ Delete Email
- `ESC` → Back
- `Tab` / `Shift+Tab` → Move between form fields

---

## 📂 Data Layer

- Fake local store: `file_store.go`
- Real backend (SMTP/IMAP): `imap_smtp_store.go`

Switch between them via `USE_REAL_EMAIL` in `.env`
