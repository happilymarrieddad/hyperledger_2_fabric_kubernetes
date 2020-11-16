package types

type Organization struct {
	IDIdentity
	StampedEntity
	Name       string `json:"name" xorm:"name"`
	IsCustomer bool   `json:"is_customer" xorm:"is_customer"`
}

func (o *Organization) TableName() string {
	return "organizations"
}
