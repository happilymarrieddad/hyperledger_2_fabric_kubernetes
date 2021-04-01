package types

import "gorm.io/gorm"

type OrderItem struct {
	gorm.Model
	Quantity  int      `json:"quantity"`
	Product   *Product `json:"product" gorm:"ForeignKey:ProductID"`
	ProductID uint     `json:"product_id"`
	OrderID   uint     `json:"order_id"`
}

func (o *OrderItem) TableName() string {
	return "order_items"
}
