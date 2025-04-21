package main

import (
	"email-client/utils"
	"encoding/json"
	"github.com/creack/pty"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

type resizeMsg struct {
	Resize bool `json:"resize"`
	Cols   int  `json:"cols"`
	Rows   int  `json:"rows"`
}

func wsHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("WebSocket upgrade error:", err)
		return
	}
	defer conn.Close()

	binaryPath := os.Getenv("EMAIL_CLIENT_BINARY_PATH")
	cmd := exec.Command(binaryPath)
	cmd.Env = append(os.Environ(), "TERM=xterm-256color")

	ptmx, err := pty.Start(cmd)
	if err != nil {
		log.Printf("🚨 Error starting TUI: %v\n", err)
		conn.WriteMessage(websocket.TextMessage, []byte("🚨 Error starting TUI: "+err.Error()))
		return
	}
	defer func() {
		ptmx.Close()
		cmd.Wait() // ensure cleanup
		log.Println("📴 TUI process exited")
	}()

	// Read from PTY and send to WebSocket
	go func() {
		buf := make([]byte, 1024)
		for {
			n, err := ptmx.Read(buf)
			if err != nil {
				log.Println("PTY read error:", err)
				break
			}
			if err := conn.WriteMessage(websocket.TextMessage, buf[:n]); err != nil {
				log.Println("WebSocket write error:", err)
				break
			}
		}
	}()

	// Read from WebSocket and write to PTY
	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			log.Println("WebSocket read error:", err)
			break
		}

		var resize resizeMsg
		if json.Unmarshal(msg, &resize) == nil && resize.Resize {
			_ = pty.Setsize(ptmx, &pty.Winsize{Cols: uint16(resize.Cols), Rows: uint16(resize.Rows)})
			continue
		}

		if _, err := ptmx.Write(msg); err != nil {
			log.Println("PTY write error:", err)
			break
		}
	}
}

func main() {
	utils.LoadEnv()
	fs := http.FileServer(http.Dir(filepath.Join("web", "static")))
	http.Handle("/", fs)
	http.HandleFunc("/ws", wsHandler)

	port := os.Getenv("WEB_APP_PORT")
	log.Printf("🚀 Server listening on http://0.0.0.0:%s", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
