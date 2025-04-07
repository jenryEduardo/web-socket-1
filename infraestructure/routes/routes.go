package routes

import (
    "encoding/json"
    "net/http"
    "websocket/infraestructure/controllers"
    "websocket/domain"
)

// Definir la ruta que recibirá el idPedido
func InitializeRoutes(wsController *controllers.WebSocketController) {
    http.HandleFunc("/ws", wsController.HandleWebSocket)

    // Ruta para recibir el idPedido desde la API y enviarlo a WebSocket
    http.HandleFunc("/enviarPedido/", func(w http.ResponseWriter, r *http.Request) {
        if r.Method != http.MethodPost {
            http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
            return
        }

        var pedido domain.Pedido
        err := json.NewDecoder(r.Body).Decode(&pedido)
        if err != nil {
            http.Error(w, "Error al leer el cuerpo de la solicitud", http.StatusBadRequest)
            return
        }

        // Llamar al servicio WebSocket para enviar el pedido
        if err := wsController.SendPedido(pedido); err != nil {
            http.Error(w, "Error al enviar el pedido", http.StatusInternalServerError)
            return
        }

        w.WriteHeader(http.StatusOK)
        w.Write([]byte("Pedido enviado exitosamente"))
    })
}
