package util

import (
	"fmt"
	"sync"

	"github.com/Cray-HPE/yapl/model"
	"github.com/gorilla/websocket"
)

var ws *websocket.Conn
var mu sync.Mutex

func ChangeStatus(genericYaml *model.GenericYAML, newStatus string) {
	genericYaml.Metadata.Status = newStatus
	if ws != nil {
		out := []byte(fmt.Sprintf(`{"id": "%s","status": "%s"}`, genericYaml.Metadata.Id, genericYaml.Metadata.Status))
		mu.Lock()
		defer mu.Unlock()
		ws.WriteMessage(websocket.TextMessage, out)
	}
}
func SetWebsocket(in_ws *websocket.Conn) {
	ws = in_ws
}
func SendMessage(msg string) {
	mu.Lock()
	defer mu.Unlock()
	ws.WriteMessage(websocket.TextMessage, []byte(msg))
}
