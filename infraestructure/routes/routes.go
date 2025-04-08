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
    http.HandleFunc("/enviarPedido", func(w http.ResponseWriter, r *http.Request) {
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

        if err := wsController.SendPedido(pedido); err != nil {
            http.Error(w, "Error al enviar el pedido", http.StatusInternalServerError)
            return
        }

        w.WriteHeader(http.StatusOK)
        w.Write([]byte("Pedido enviado exitosamente"))
    })

    // Rutas para sensores
    http.HandleFunc("/statusSensor1", func(w http.ResponseWriter, r *http.Request) {
        handleSensorStatus(w, r, wsController, "sensor1")
    })

    http.HandleFunc("/statusSensor2", func(w http.ResponseWriter, r *http.Request) {
        handleSensorStatus(w, r, wsController, "sensor2")
    })

    http.HandleFunc("/statusSensor3", func(w http.ResponseWriter, r *http.Request) {
        handleSensorStatus(w, r, wsController, "sensor3")
    })
}

// ✅ Esta función debe ir fuera de InitializeRoutes
func handleSensorStatus(w http.ResponseWriter, r *http.Request, wsController *controllers.WebSocketController, sensorID string) {
    if r.Method != http.MethodPost {
        http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
        return
    }

    var status domain.SensorStatus
    err := json.NewDecoder(r.Body).Decode(&status)
    if err != nil {
        http.Error(w, "Error al leer el cuerpo de la solicitud", http.StatusBadRequest)
        return
    }

    status.SensorID = sensorID

    if err := wsController.SendSensorStatus(status); err != nil {
        http.Error(w, "Error al enviar el estado del sensor", http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusOK)
    w.Write([]byte("Estado del sensor enviado correctamente"))
}
