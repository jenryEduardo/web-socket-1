package app

import (
    "websocket/domain"
    "websocket/infraestructure/controllers"
)

// PedidoService maneja la lógica de negocio relacionada con los pedidos.
type PedidoService struct {
    WebSocketController *controllers.WebSocketController
}

// NewPedidoService crea una nueva instancia de PedidoService.
func NewPedidoService(wsController *controllers.WebSocketController) *PedidoService {
    return &PedidoService{WebSocketController: wsController}
}

// EnviarPedido envía un pedido al cliente a través del WebSocket.
func (s *PedidoService) EnviarPedido(pedido domain.Pedido) error {
    return s.WebSocketController.SendPedido(pedido)
}
