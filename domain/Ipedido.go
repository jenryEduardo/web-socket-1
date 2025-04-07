package domain

// PedidoPort define el puerto que será utilizado por la capa de aplicación.
type PedidoPort interface {
    EnviarPedido(pedido Pedido) error
}