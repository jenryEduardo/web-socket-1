package adapters

import (
    "log"
    "github.com/gorilla/websocket"
    "websocket/domain"
)

type WebSocketServer struct {
    Clients         map[*websocket.Conn]bool
    Register        chan *websocket.Conn
    Unregister      chan *websocket.Conn
    Broadcast       chan domain.Pedido
    BroadcastSensor chan domain.SensorStatus
}

func NewWebSocketServer() *WebSocketServer {
    return &WebSocketServer{
        Clients:         make(map[*websocket.Conn]bool),
        Register:        make(chan *websocket.Conn),
        Unregister:      make(chan *websocket.Conn),
        Broadcast:       make(chan domain.Pedido),
        BroadcastSensor: make(chan domain.SensorStatus),
    }
}

func (s *WebSocketServer) Run() {
    for {
        select {
        case conn := <-s.Register:
            s.Clients[conn] = true

        case conn := <-s.Unregister:
            delete(s.Clients, conn)

        case pedido := <-s.Broadcast:
            for client := range s.Clients {
                err := client.WriteJSON(map[string]interface{}{
                    "type": "pedido",
                    "data": pedido,
                })
                if err != nil {
                    log.Println("Error al enviar pedido:", err)
                    client.Close()
                    delete(s.Clients, client)
                }
            }

        case sensor := <-s.BroadcastSensor:
            for client := range s.Clients {
                err := client.WriteJSON(map[string]interface{}{
                    "type": "sensor",
                    "data": sensor,
                })
                if err != nil {
                    log.Println("Error al enviar sensor:", err)
                    client.Close()
                    delete(s.Clients, client)
                }
            }
        }
    }
}
