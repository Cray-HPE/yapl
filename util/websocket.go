package util

import (
	"sync"

	"github.com/Cray-HPE/yapl/model"
	"github.com/gorilla/websocket"
	"gopkg.in/yaml.v2"
)

var ws *websocket.Conn
var mu sync.Mutex

func ChangeStatus(genericYaml *model.GenericYAML, newStatus string) {
	genericYaml.Metadata.Status = newStatus
	if ws != nil {
		out, _ := yaml.Marshal(genericYaml)
		mu.Lock()
		defer mu.Unlock()
		ws.WriteMessage(websocket.TextMessage, out)
	}
}
func SetWebsocket(in_ws *websocket.Conn) {
	ws = in_ws
}
