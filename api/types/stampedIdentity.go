package types

import "time"

type StampedEntity struct {
	CreatedAt *time.Time `json:"created_at" xorm:"created_at"`
	UpdatedAt *time.Time `json:"updated_at" xorm:"updated_at"`
}
