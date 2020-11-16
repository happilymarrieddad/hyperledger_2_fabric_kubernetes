package types

type Order struct {
	IDIdentity
	StampedEntity
	Number         string      `json:"number" xorm:"number"`
	OrganizationID []int64     `json:"organization_id" xorm:"organization_id"`
	OrderItems     []OrderItem `json:"order_items" xorm:"order_items"`
}

func (o *Order) TableName() string {
	return "orders"
}
