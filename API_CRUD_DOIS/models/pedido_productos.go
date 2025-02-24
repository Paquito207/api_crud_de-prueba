package models

type PedidoProductos struct {
	PedidoId   *Pedidos   `orm:"column(pedido_id);rel(fk)"`
	ProductoId *Productos `orm:"column(producto_id);rel(fk)"`
	Cantidad   int        `orm:"column(cantidad)"`
}
