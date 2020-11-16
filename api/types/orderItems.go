package types

type OrderItem struct {
	IDIdentity
	StampedEntity
	Quantity       int     `json:"quantity" xorm:"quantity"`
	ProductID      int64   `json:"product_id" xorm:"product_id"`
	OrderID        int64   `json:"order_id" xorm:"order_id"`
	OrganizationID []int64 `json:"organization_id" xorm:"organization_id"`
	Product        Product `json:"product" xorm:"product"`
}

func (o *OrderItem) TableName() string {
	return "order_items"
}
