package models_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "api/models"
	"api/types"
)

var _ = Describe("OrdersModel", func() {

	var OrdersModel Orders
	var ord *types.Order
	var ord2 *types.Order
	var org *types.Organization
	var org2 *types.Organization

	setup := func() {
		orgModel := NewOrganizations(db)

		org = &types.Organization{
			Name:       "Some Organization",
			IsCustomer: true,
		}
		err := orgModel.Create(org)
		Expect(err).To(BeNil())

		org2 = &types.Organization{
			Name:       "Some Organization 2",
			IsCustomer: true,
		}
		err = orgModel.Create(org2)
		Expect(err).To(BeNil())

		OrdersModel = NewOrders(db)

		ord = &types.Order{
			Number:        "1234",
			Organizations: []types.Organization{*org},
		}

		ord2 = &types.Order{
			Number:        "4321",
			Organizations: []types.Organization{*org, *org2},
		}
	}

	create := func() {
		err := OrdersModel.Create(ord)
		Expect(err).To(BeNil())
		err = OrdersModel.Create(ord2)
		Expect(err).To(BeNil())
	}

	BeforeEach(func() {
		cleanMainDatabaseTables()
		setup()
	})

	Context("Create", func() {
		It("should fail to create an order with an existing number", func() {
			create()

			ord = &types.Order{
				Number:        "1234",
				Organizations: []types.Organization{*org},
			}

			err := OrdersModel.Create(ord)
			Expect(err).NotTo(BeNil())
		})

		It("should successfully create an order", create)
	})

	Context("Get", func() {
		BeforeEach(create)

		It("should be able to find the first order", func() {
			existingOrder, err := OrdersModel.Get(ord.ID)
			Expect(err).To(BeNil())
			Expect(existingOrder.Number).To(Equal(ord.Number))
			Expect(existingOrder.Organizations).NotTo(BeEmpty())
			Expect(existingOrder.Organizations[0].Name).To(Equal(org.Name))
		})

		It("should be able to find the second order", func() {
			existingOrder, err := OrdersModel.Get(ord2.ID)
			Expect(err).To(BeNil())
			Expect(existingOrder.Number).To(Equal(ord2.Number))
		})

	})

	// Context("Find", func() {
	// 	BeforeEach(create)

	// 	It("should return 2 products with associated data", func() {
	// 		products, err := OrdersModel.Find()
	// 		Expect(err).To(BeNil())
	// 		Expect(len(products)).To(Equal(2))
	// 		Expect(products[0].Organizations).NotTo(BeEmpty())
	// 		Expect(products[1].Organizations).NotTo(BeEmpty())
	// 		Expect(products[0].Organizations[0].Name).To(Equal(org.Name))
	// 		Expect(products[1].Organizations[0].Name).To(Equal(org.Name))
	// 	})

	// })

	// Context("Update", func() {
	// 	BeforeEach(create)

	// 	It("should update a product", func() {
	// 		pro.Name = "SOME NEW VAL"
	// 		err := OrdersModel.Update(pro)
	// 		Expect(err).To(BeNil())
	// 		existingOrder, err := OrdersModel.Get(pro.ID)
	// 		Expect(err).To(BeNil())
	// 		Expect(existingOrder.Name).To(Equal(pro.Name))
	// 	})

	// })

	// Context("Delete", func() {
	// 	BeforeEach(create)

	// 	It("should update a product", func() {
	// 		err := OrdersModel.Delete(pro.ID)
	// 		Expect(err).To(BeNil())
	// 		products, err := OrdersModel.Find()
	// 		Expect(err).To(BeNil())
	// 		Expect(len(products)).To(Equal(1))
	// 	})

	// })

})
