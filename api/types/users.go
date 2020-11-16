package types

type User struct {
	IDIdentity
	StampedEntity
	FirstName       string  `json:"first_name" xorm:"first_name"`
	LastName        string  `json:"last_name" xorm:"last_name"`
	Email           string  `json:"email" xorm:"email"`
	Password        string  `json:"password" xorm:"password"`
	OrganizationIDs []int64 `json:"organization_ids" xorm:"organization_ids"`
}

func (u *User) TableName() string {
	return "users"
}
