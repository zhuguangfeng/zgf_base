package websocket

import (
	"github.com/gorilla/websocket"
	"net/http"
	"testing"
	"time"
)

func TestServer(t *testing.T) {

	upgrader := &websocket.Upgrader{}
	http.HandleFunc("/ws", func(writer http.ResponseWriter, request *http.Request) {
		conn, err := upgrader.Upgrade(writer, request, nil)
		if err != nil {
			writer.Write([]byte("升级ws失败"))
			return
		}

		go func() {
			for {
				//typ是websocket里的消息类型
				typ, msg, err := conn.ReadMessage()

				if err != nil {

				}
				switch typ {
				case websocket.CloseMessage:
					conn.Close()
				default:
					t.Log(typ, string(msg))
				}
			}
		}()

		go func() {
			ticker := time.NewTicker(time.Second * 3)
			for now := range ticker.C {
				err := conn.WriteMessage(websocket.TextMessage, []byte("hello "+now.String()))
				if err != nil {
					return
				}
			}
		}()

	})
	http.ListenAndServe(":8081", nil)
}
