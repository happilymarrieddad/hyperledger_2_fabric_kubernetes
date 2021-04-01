package models

import (
	"api/types"

	"gorm.io/gorm"
)

type Orders interface {
	Get(uint) (*types.Order, error)
	Create(*types.Order) error
}

func NewOrders(db *gorm.DB) Orders {
	return &orders{db}
}

type orders struct {
	db *gorm.DB
}

func (o *orders) Get(id uint) (ord *types.Order, err error) {
	ord = new(types.Order)

	tx := o.db.
		Preload("Organizations").
		Preload("OrderItems").
		First(ord, id)
	err = tx.Error

	return
}

func (o *orders) Create(ord *types.Order) error {
	tx := o.db.Create(ord)

	return tx.Error
}
