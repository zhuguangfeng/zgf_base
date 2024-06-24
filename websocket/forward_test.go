package websocket

import (
	"github.com/ecodeclub/ekit/syncx"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"testing"
)

type Hub struct {
	conns syncx.Map[string, *websocket.Conn]
}

func (h *Hub) AddConn(name string, conn *websocket.Conn) {
	h.conns.Store(name, conn)
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
				log.Println("来自客户端", name, typ, string(msg))
				h.conns.Range(func(key string, value *websocket.Conn) bool {
					if key == name {
						return true
					}
					log.Println("转发给", key)
					err := value.WriteMessage(typ, msg)
					if err != nil {
						log.Println(err)
					}
					return true
				})
			}
		}
	}()
}

func TestHub(t *testing.T) {
	upgrader := &websocket.Upgrader{}
	hub := &Hub{conns: syncx.Map[string, *websocket.Conn]{}}
	http.HandleFunc("/ws", func(writer http.ResponseWriter, request *http.Request) {
		c, err := upgrader.Upgrade(writer, request, nil)
		if err != nil {
			writer.Write([]byte("升级 ws失败"))
			return
		}
		name := request.URL.Query().Get("name")
		hub.AddConn(name, c)
	})
	http.ListenAndServe(":8081", nil)
}
