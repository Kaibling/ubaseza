package webserver

import (
	"embed"
	"fmt"
	"io/fs"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

//go:embed static
var embededFiles embed.FS
var upgrader = websocket.Upgrader{}

type WebServer struct {
	c chan string
}

func NewWebServer(c chan string) *WebServer {
	return &WebServer{c: c}
}
func (s *WebServer) Configure() {
	http.HandleFunc("/wsa", func(w http.ResponseWriter, r *http.Request) {
		// Upgrade upgrades the HTTP server connection to the WebSocket protocol.
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Print("upgrade failed: ", err)
			return
		}
		defer conn.Close()

		// Continuosly read and write message
		for {
			message := <-s.c
			err := conn.WriteMessage(1, []byte(message))
			if err != nil {
				log.Println("write failed:", err)
				break
			}
		}
	})
	http.Handle("/", http.FileServer(getFileSystem()))
}

func (s *WebServer) Start() {
	fmt.Println("ready....")
	http.ListenAndServe(":8080", nil)
}

func getFileSystem() http.FileSystem {

	log.Print("using embed mode")
	fsys, err := fs.Sub(embededFiles, "static")
	if err != nil {
		panic(err)
	}

	return http.FS(fsys)
}
