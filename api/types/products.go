package types

type Product struct {
	IDIdentity
	StampedEntity
	Name            string  `json:"name" xorm:"name"`
	OrganizationIDs []int64 `json:"organization_ids" xorm:"organization_ids"`
}

func (p *Product) TableName() string {
	return "products"
}
