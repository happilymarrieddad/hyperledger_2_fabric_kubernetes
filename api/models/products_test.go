package models_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "api/models"
	"api/types"
)

var _ = Describe("ProductsModel", func() {

	var ProductsModel Products
	var pro *types.Product
	var pro2 *types.Product
	var org *types.Organization

	setup := func() {
		orgModel := NewOrganizations(db)

		org = &types.Organization{
			Name:       "Some Organization",
			IsCustomer: true,
		}
		err := orgModel.Create(org)
		Expect(err).To(BeNil())

		ProductsModel = NewProducts(db)

		pro = &types.Product{
			Name:          "Product 1",
			Organizations: []types.Organization{*org},
		}

		pro2 = &types.Product{
			Name:          "Product 2",
			Organizations: []types.Organization{*org},
		}
	}

	create := func() {
		err := ProductsModel.Create(pro)
		Expect(err).To(BeNil())
		err = ProductsModel.Create(pro2)
		Expect(err).To(BeNil())
	}

	BeforeEach(func() {
		cleanMainDatabaseTables()
		setup()
	})

	Context("Create", func() {
		It("should successfully create a place", create)
	})

	Context("Get", func() {
		BeforeEach(create)

		It("should be able to find the first product", func() {
			existingProduct, err := ProductsModel.Get(pro.ID)
			Expect(err).To(BeNil())
			Expect(existingProduct.Name).To(Equal(pro.Name))
			Expect(existingProduct.Organizations).NotTo(BeEmpty())
			Expect(existingProduct.Organizations[0].Name).To(Equal(org.Name))
		})

		It("should be able to find the second product", func() {
			existingProduct, err := ProductsModel.Get(pro2.ID)
			Expect(err).To(BeNil())
			Expect(existingProduct.Name).To(Equal(pro2.Name))
		})

	})

	Context("Find", func() {
		BeforeEach(create)

		It("should return 2 products with associated data", func() {
			products, err := ProductsModel.Find()
			Expect(err).To(BeNil())
			Expect(len(products)).To(Equal(2))
			Expect(products[0].Organizations).NotTo(BeEmpty())
			Expect(products[1].Organizations).NotTo(BeEmpty())
			Expect(products[0].Organizations[0].Name).To(Equal(org.Name))
			Expect(products[1].Organizations[0].Name).To(Equal(org.Name))
		})

	})

	Context("Update", func() {
		BeforeEach(create)

		It("should update a product", func() {
			pro.Name = "SOME NEW VAL"
			err := ProductsModel.Update(pro)
			Expect(err).To(BeNil())
			existingProduct, err := ProductsModel.Get(pro.ID)
			Expect(err).To(BeNil())
			Expect(existingProduct.Name).To(Equal(pro.Name))
		})

	})

	Context("Delete", func() {
		BeforeEach(create)

		It("should update a product", func() {
			err := ProductsModel.Delete(pro.ID)
			Expect(err).To(BeNil())
			products, err := ProductsModel.Find()
			Expect(err).To(BeNil())
			Expect(len(products)).To(Equal(1))
		})

	})

})
