package types

type IDIdentity struct {
	ID string `json:"id" xorm:"'id' pk auto"`
}
