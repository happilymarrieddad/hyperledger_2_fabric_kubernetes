package models

import (
	"api/types"

	"gorm.io/gorm"
)

type Organizations interface {
	Find() ([]*types.Organization, error)
	Get(uint) (*types.Organization, error)
	Create(*types.Organization) error
	Update(*types.Organization) error
	Delete(uint) error
}

func NewOrganizations(db *gorm.DB) Organizations {
	return &organizations{db}
}

type organizations struct {
	db *gorm.DB
}

func (p *organizations) Find() (res []*types.Organization, err error) {
	tx := p.db.Find(&res)
	err = tx.Error

	return
}

func (p *organizations) Get(id uint) (organization *types.Organization, err error) {
	organization = new(types.Organization)

	tx := p.db.First(organization, id)
	err = tx.Error

	return
}

func (p *organizations) Create(newOrganization *types.Organization) (err error) {
	tx := p.db.Create(newOrganization)

	return tx.Error
}

func (p *organizations) Update(updateOrganization *types.Organization) error {
	tx := p.db.Save(updateOrganization)

	return tx.Error
}

func (p *organizations) Delete(id uint) (err error) {
	tx := p.db.Delete(&types.Organization{}, id)

	return tx.Error
}
