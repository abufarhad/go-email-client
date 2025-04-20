# 📬 Terminal Email Client

A fast and minimal terminal-based email client built in Go. Inspired by tools like `mutt`, `aerc`, and `himalaya`.

This project demonstrates clean architecture, form-based input, SMTP/IMAP support, and async operations using Go + TUI (`tview`).

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
│   └── main.go                       # Entry point for terminal app
├── internal/
│   ├── domain/
│   │   ├── model/
│   │   │   └── email.go              # Email entity/model
│   │   └── service/
│   │       └── email_service.go      # Business logic
│   ├── infra/
│   │   └── logger/
│   │       └── logger.go             # Logger setup
│   └── interface/
│       ├── controller/
│       │   └── handler.go            # Application layer
│       ├── persistence/
│       │   ├── file_store.go         # File-based backend
│       │   └── imap_smtp_store.go    # Real IMAP/SMTP backend
│       └── ui/
│           └── app.go                # TUI (tview)
├── web/
│   ├── static/
│   │   └── index.html                # Web terminal via xterm.js
│   └── main.go                       # WebSocket/PTY server
├── emails.json                       # Local file email DB
├── logs.txt                          # Log output
├── Dockerfile                        # Multi-stage Docker build
├── Makefile                          # CLI helpers
├── .env                              # Config vars
└── README.md                         # You're here
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

---

## 📄 License

MIT License. Use it, share it, build on it — but don’t send spam 😏
