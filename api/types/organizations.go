package types

import "gorm.io/gorm"

type Organization struct {
	gorm.Model
	Name       string `json:"name"`
	IsCustomer bool   `json:"is_customer"`
}

func (o *Organization) TableName() string {
	return "organizations"
}
