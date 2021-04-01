package types

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	Number        string         `json:"number" gorm:"uniqueIndex"`
	Organizations []Organization `json:"organizations" gorm:"many2many:order_organizations;"`
	OrderItems    []OrderItem    `json:"order_items" gorm:"foreignKey:OrderID"`
}

func (o *Order) TableName() string {
	return "orders"
}
