package models_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "api/models"
	"api/types"
)

var _ = Describe("OrderItemsModel", func() {

	var OrderItemsModel OrderItems
	var ord *types.Order
	var ordItm *types.OrderItem
	var ordItm2 *types.OrderItem
	var org *types.Organization
	var pro *types.Product
	var pro2 *types.Product

	setup := func() {
		orgModel := NewOrganizations(db)

		org = &types.Organization{
			Name:       "Some Organization",
			IsCustomer: true,
		}
		err := orgModel.Create(org)
		Expect(err).To(BeNil())

		proModel := NewProducts(db)

		pro = &types.Product{
			Name:          "Some Product",
			Organizations: []types.Organization{*org},
		}
		err = proModel.Create(pro)
		Expect(err).To(BeNil())

		pro2 = &types.Product{
			Name:          "Some Product 2",
			Organizations: []types.Organization{*org},
		}
		err = proModel.Create(pro2)
		Expect(err).To(BeNil())

		OrdersModel := NewOrders(db)

		ord = &types.Order{
			Number:        "6234",
			Organizations: []types.Organization{*org},
		}

		err = OrdersModel.Create(ord)
		Expect(err).To(BeNil())

		OrderItemsModel = NewOrderItems(db)

		ordItm = &types.OrderItem{
			Quantity:  50,
			OrderID:   ord.ID,
			ProductID: pro.ID,
		}

		ordItm2 = &types.OrderItem{
			Quantity:  25,
			OrderID:   ord.ID,
			ProductID: pro2.ID,
		}
	}

	create := func() {
		err := OrderItemsModel.Create(ordItm)
		Expect(err).To(BeNil())
		err = OrderItemsModel.Create(ordItm2)
		Expect(err).To(BeNil())
	}

	BeforeEach(func() {
		cleanMainDatabaseTables()
		setup()
	})

	Context("Create", func() {
		It("should successfully create an order item", create)
	})

	Context("Get", func() {
		BeforeEach(create)

		It("should be able to find the first order item", func() {
			existingOrderItem, err := OrderItemsModel.Get(ordItm.ID)
			Expect(err).To(BeNil())
			Expect(existingOrderItem.Quantity).To(Equal(ordItm.Quantity))
			Expect(existingOrderItem.OrderID).To(Equal(ord.ID))
			// Expect(existingOrderItem.Product.Name).To(Equal(pro.Name))
		})

		It("should be able to find the second order item", func() {
			existingOrderItem, err := OrderItemsModel.Get(ordItm2.ID)
			Expect(err).To(BeNil())
			Expect(existingOrderItem.Quantity).To(Equal(ordItm2.Quantity))
			Expect(existingOrderItem.OrderID).To(Equal(ord.ID))
			// Expect(existingOrderItem.Product.Name).To(Equal(pro2.Name))
		})

	})

	// Context("Find", func() {
	// 	BeforeEach(create)

	// 	It("should return 2 products with associated data", func() {
	// 		products, err := OrderItemsModel.Find()
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
	// 		err := OrderItemsModel.Update(pro)
	// 		Expect(err).To(BeNil())
	// 		existingOrderItem, err := OrderItemsModel.Get(pro.ID)
	// 		Expect(err).To(BeNil())
	// 		Expect(existingOrderItem.Name).To(Equal(pro.Name))
	// 	})

	// })

	// Context("Delete", func() {
	// 	BeforeEach(create)

	// 	It("should update a product", func() {
	// 		err := OrderItemsModel.Delete(pro.ID)
	// 		Expect(err).To(BeNil())
	// 		products, err := OrderItemsModel.Find()
	// 		Expect(err).To(BeNil())
	// 		Expect(len(products)).To(Equal(1))
	// 	})

	// })

})
