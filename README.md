# 📬 Terminal Email Client

A fast and minimal terminal-based email client built in Go. Inspired by tools like `mutt`, `aerc`, and `himalaya`.

This project demonstrates clean architecture, form-based input, and async operations using Go + TUI (`tview`).

---

## ✨ Features

- 📥 View Inbox
- 📝 Compose plain text emails
- 🗑️ Delete selected emails
- ESC/back navigation from all views
- ✅ Email validation and success messages
- 🐳 Docker support
- 🧼 Clean, modular code structure

---

## 🧪 Project Structure

```
go-email-client/
├── cmd/
│   └── main.go                       # Entry point for the terminal app
├── internal/
│   ├── domain/
│   │   ├── model/
│   │   │   └── email.go              # Email entity/model
│   │   └── service/
│   │       └── email_service.go      # Business logic for email operations
│   ├── infra/
│   │   └── logger/
│   │       └── logger.go             # Logging setup
│   └── interface/
│       ├── controller/
│       │   └── handler.go            # Application handlers (e.g. for email logic)
│       ├── persistence/
│       │   └── file_store.go         # Local storage for emails (JSON file)
│       └── ui/
│           └── app.go                # TUI (terminal UI) with tview
├── web/
│   ├── static/
│   │   └── index.html                # Web frontend (xterm.js)
│   └── main.go                       # WebSocket + PTY server for browser UI
├── emails.json                       # Local email data
├── Dockerfile                        # Multi-stage Docker setup
├── Makefile                          # Build, run, dockerize the app
├── go.mod                            # Go modules metadata
├── go.sum                            # Go modules checksum
└── README.md                         # Project overview (you're here)
```

---

## 🛠️ Requirements

- Go 1.21+
- Docker (optional)
- `make` (for easier commands)

---

## 🚀 Getting Started

### 🔧 Local Build (No Docker)

🔹 Run Terminal-Only App

```bash
go run cmd/main.go
```

🔹 Run Web-Based Terminal UI

```bash
go build -o email-client ./cmd
go run cmd/main.go
```
🔹 Or Use the Makefile (Recommended)

```bash
make build
make run
```
Then open the app in your browser:

```bash
http://localhost:8080/
```

### 🐳 Docker Workflow

```bash
make docker-build   # Builds Docker image with terminal + WebSocket server
make docker-run     # Runs the app in a Docker container on port 8080
make clean          # Removes built binaries
```
Server will be available at:
http://localhost:8080/
---

## 🧑‍💻 Controls

- `i` → 📥 Open Inbox
- `c` → 📝 Compose Email
- `d` → 🗑️ Delete Email
- `ESC` → Go back to previous screen
- `Tab` / `Shift+Tab` → Navigate form fields

---

## 📂 Data

Emails are stored in-memory or via simple file store. Modify persistence logic in:
`internal/interface/persistence/file_store.go`

---

## 📄 License

MIT License. Do whatever you want. Just don’t send spam. 😄