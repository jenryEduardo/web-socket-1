package domain


// Pedido define la estructura del pedido que se enviará por WebSocket.
type Pedido struct {
    ID int `json:"idPedido"`
}
