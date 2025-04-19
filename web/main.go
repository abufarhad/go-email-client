package main

import (
	"encoding/json"
	"github.com/creack/pty"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
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

	cmd := exec.Command("./email-client")

	ptmx, err := pty.Start(cmd)
	if err != nil {
		log.Fatal(err)
	}
	defer ptmx.Close()

	go func() {
		buf := make([]byte, 1024)
		for {
			n, err := ptmx.Read(buf)
			if err != nil {
				return
			}
			conn.WriteMessage(websocket.TextMessage, buf[:n])
		}
	}()

	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			return
		}
		var resize resizeMsg
		if json.Unmarshal(msg, &resize) == nil && resize.Resize {
			pty.Setsize(ptmx, &pty.Winsize{Cols: uint16(resize.Cols), Rows: uint16(resize.Rows)})
			continue
		}
		ptmx.Write(msg)
	}
}

func main() {
	fs := http.FileServer(http.Dir(filepath.Join("web", "static")))
	http.Handle("/", fs)
	http.HandleFunc("/ws", wsHandler)

	log.Println("Listening on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
