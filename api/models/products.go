package models

import (
	"api/types"

	"gorm.io/gorm"
)

type Products interface {
	Find() ([]*types.Product, error)
	Get(uint) (*types.Product, error)
	Create(*types.Product) error
	Update(*types.Product) error
	Delete(uint) error
}

func NewProducts(db *gorm.DB) Products {
	return &products{db}
}

type products struct {
	db *gorm.DB
}

func (p *products) Find() (res []*types.Product, err error) {
	tx := p.db.Preload("Organizations").Find(&res)
	err = tx.Error

	return
}

func (p *products) Get(id uint) (product *types.Product, err error) {
	product = new(types.Product)

	tx := p.db.
		Preload("Organizations").
		First(product, id)
	err = tx.Error

	return
}

func (p *products) Create(newProduct *types.Product) (err error) {
	tx := p.db.Create(newProduct)

	return tx.Error
}

func (p *products) Update(updateProduct *types.Product) error {
	tx := p.db.Save(updateProduct)

	return tx.Error
}

func (p *products) Delete(id uint) (err error) {
	tx := p.db.Delete(&types.Product{}, id)

	return tx.Error
}
