package main

import (
    "log"
    "net/http"
    "websocket/infraestructure/controllers"
    "websocket/infraestructure/routes"
    "websocket/infraestructure/adapters"
)

func main() {
    // Configurar WebSocket Server
    wsServer := adapters.NewWebSocketServer()
    go wsServer.Run()

    // Crear el controlador WebSocket
    wsController := controllers.NewWebSocketController(wsServer)

    // Inicializar las rutas
    routes.InitializeRoutes(wsController)

    log.Println("Servidor WebSocket iniciado en :3010")
    log.Fatal(http.ListenAndServe(":3010", nil))
}
