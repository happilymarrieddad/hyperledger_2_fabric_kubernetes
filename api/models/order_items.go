package models

import (
	"api/types"

	"gorm.io/gorm"
)

type OrderItems interface {
	Get(uint) (*types.OrderItem, error)
	Create(*types.OrderItem) error
}

func NewOrderItems(db *gorm.DB) OrderItems {
	return &orderItems{db}
}

type orderItems struct {
	db *gorm.DB
}

func (o *orderItems) Get(id uint) (ordItem *types.OrderItem, err error) {
	ordItem = new(types.OrderItem)

	tx := o.db.First(ordItem, id)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return
}

func (o *orderItems) Create(ordItem *types.OrderItem) error {
	tx := o.db.Create(ordItem)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}
