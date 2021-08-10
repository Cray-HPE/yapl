package util

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/pterm/pterm"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,

	// We'll need to check the origin of our connection
	// this will allow us to make requests from our React
	// development server to here.
	// For now, we'll do no checking and just allow any connection
	CheckOrigin: func(r *http.Request) bool { return true },
}
var cfg *Config

// define a reader which will listen for
// new messages being sent to our WebSocket
// endpoint
func reader(conn *websocket.Conn) {
	for {
		// read in a message
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		// print out that message for clarity
		fmt.Println(string(p))

		if err := conn.WriteMessage(messageType, p); err != nil {
			log.Println(err)
			return
		}

	}
}

// define our WebSocket endpoint
func serveWs(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Host)

	// upgrade this connection to a WebSocket
	// connection
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}
	SetWebsocket(ws)
	// listen indefinitely for new messages coming
	// through on our WebSocket connection
	reader(ws)
}

var status string = "Not started"

func setupRoutes() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Simple Server")
	})

	// /ws endpoint
	http.HandleFunc("/ws", serveWs)

	// /start endpoint
	http.HandleFunc("/start", func(w http.ResponseWriter, r *http.Request) {
		if status == "running" {
			fmt.Fprintf(w, `{"Status":"another pipeline is running"}`)
			return
		}
		status = "running"
		SendMessage(fmt.Sprintf(`{"id": "0","status": "%s"}`, status))
		fmt.Fprintf(w, `{"Status":"Started"}`)
		ClearCache()
		err := ExecutePipeline(cfg)
		if nil != err {
			status = "Failed"
			SendMessage(fmt.Sprintf(`{"id": "0","status": "%s"}`, status))
			return
		}
		status = "Completed"
		SendMessage(fmt.Sprintf(`{"id": "0","status": "%s"}`, status))
	})

	// /status endpoint
	http.HandleFunc("/status", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, `{"Status":"%s"}`, status)
	})
}

func Serve(inCfg *Config) error {
	pterm.Info.Println("YAPL server mode")
	cfg = inCfg
	setupRoutes()
	return http.ListenAndServe(":8080", nil)
}
