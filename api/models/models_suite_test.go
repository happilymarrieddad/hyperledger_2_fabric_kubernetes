package models_test

import (
	"api/types"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB
var err error

func cleanMainDatabaseTables() {
	db.Session(&gorm.Session{AllowGlobalUpdate: true}).Exec("TRUNCATE products CASCADE")
	db.Session(&gorm.Session{AllowGlobalUpdate: true}).Exec("TRUNCATE order_items CASCADE")
	db.Session(&gorm.Session{AllowGlobalUpdate: true}).Exec("TRUNCATE orders CASCADE")
	db.Session(&gorm.Session{AllowGlobalUpdate: true}).Exec("TRUNCATE organizations CASCADE")
}

var _ = BeforeSuite(func() {
	defer GinkgoRecover()

	dsn := "user=postgres password=postgres dbname=order-manager-test port=5432 host=localhost connect_timeout=180 sslmode=disable TimeZone=America/Denver"
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	Expect(err).To(BeNil())

	db.AutoMigrate(&types.Organization{})
	db.AutoMigrate(&types.Product{})
	db.AutoMigrate(&types.Order{})
	db.AutoMigrate(&types.OrderItem{})

})

func TestModels(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Models Suite")
}
