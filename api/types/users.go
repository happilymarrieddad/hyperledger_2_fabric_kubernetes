package types

import "time"

type User struct {
	ID              int64      `json:"id" xorm:"'id' pk autoincr"`
	FirstName       string     `json:"first_name" xorm:"first_name"`
	LastName        string     `json:"last_name" xorm:"last_name"`
	Email           string     `json:"email" xorm:"email"`
	Password        string     `json:"password" xorm:"password"`
	OrganizationIDs []int64    `json:"organization_ids" xorm:"organization_ids"`
	CreatedAt       *time.Time `json:"created_at" xorm:"created_at"`
	UpdatedAt       *time.Time `json:"updated_at" xorm:"updated_at"`
}

func (u *User) TableName() string {
	return "users"
}
