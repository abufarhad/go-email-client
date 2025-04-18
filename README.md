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

## 🛠️ Requirements

- Go 1.21+
- Docker (optional)
- `make` (for easier commands)

---

## 🚀 Getting Started

### 🔧 Local Build (No Docker)

Manually:

```bash
go run cmd/main.go
```

### 🐳 Docker Workflow

```bash
make build       # Builds Docker image
make run         # Runs the email client inside a container
make clean       # Removes the Docker image
```

---

## 🧑‍💻 Controls

- `i` → 📥 Open Inbox
- `c` → 📝 Compose Email
- `d` → 🗑️ Delete Email
- `ESC` → Go back to previous screen
- `Tab` / `Shift+Tab` → Navigate form fields

---

## 🧪 Project Structure

```
├── cmd/main.go             # Entry point
├── internal/
│   ├── domain/             # Business logic
│   ├── interface/
│   │   ├── controller/     # Handlers
│   │   ├── persistence/    # File-based store
│   │   └── ui/             # TUI components
│   └── model/              # Email model
├── Dockerfile
├── Makefile
└── README.md
```

---

## 📂 Data

Emails are stored in-memory or via simple file store. Modify persistence logic in:
`internal/interface/persistence/file_store.go`

---

## 📄 License

MIT License. Do whatever you want. Just don’t send spam. 😄