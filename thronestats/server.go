package thronestats

import (
	"fmt"
	"log"
	"net/http"
	"github.com/nu7hatch/gouuid"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}
var connections = map[uuid.UUID]ConnectionWrapper{}

type ConnectionWrapper struct {
	uuid uuid.UUID
	conn *websocket.Conn
}

func (c *ConnectionWrapper) Close() {
	Unsubscribe(c.uuid)
	delete(connections, c.uuid)
}

func (c *ConnectionWrapper) Send(data []byte) {
	c.conn.WriteMessage(websocket.TextMessage, data)
}

func SendToConnection(uuid uuid.UUID, data []byte) {
	conn, ok := connections[uuid]

	if ok {
		conn.Send(data)
	}
}

func websocketHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)

	if err != nil {
		log.Print("upgrade:", err)
		return
	}

	defer conn.Close()

	uuidPtr, err := uuid.NewV4()
	uuid := *uuidPtr

	log.Printf("Connection %s from %s", uuid.String(), r.RemoteAddr)

	if err != nil {
		log.Fatalf("error:", err)
	}

	wrapper := ConnectionWrapper{
		uuid,
		conn,
	}

	connections[uuid] = wrapper

	SendHello(uuid)
	SendGlobalStats(uuid)

	for {
		_, message, err := conn.ReadMessage()

		if err != nil {
			if err.Error() == "websocket: close 1001 " {
				log.Printf("Client %s disconnected", uuid.String())
			} else {
				log.Printf("Client %s error, %s", err.Error())
			}
			wrapper.Close()
			break
		}

		HandleMessage(uuid, message)
	}
}

func RunServer(settings ServerSettings) {
	address := fmt.Sprintf("%s:%d", settings.ListenAddress, settings.ListenPort)

	log.Printf("Listening to %s", address)
	log.Printf("Serving static files from %s", settings.WwwPath)

	http.HandleFunc("/data", websocketHandler)
	http.Handle("/", http.FileServer(http.Dir(settings.WwwPath)))

	log.Fatal(http.ListenAndServe(address, nil))
}