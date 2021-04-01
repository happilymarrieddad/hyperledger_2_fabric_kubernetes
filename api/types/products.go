package types

import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name          string         `json:"name"`
	Organizations []Organization `json:"organizations" gorm:"many2many:product_organizations;"`
}

func (p *Product) TableName() string {
	return "products"
}
