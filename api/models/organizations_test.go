package models_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "api/models"
	"api/types"
)

var _ = Describe("OrganizationsModel", func() {

	var OrganizationsModel Organizations
	var org *types.Organization
	var org2 *types.Organization

	setup := func() {
		OrganizationsModel = NewOrganizations(db)

		org = &types.Organization{
			Name: "Organization 1",
		}

		org2 = &types.Organization{
			Name: "Organization 2",
		}
	}

	create := func() {
		err := OrganizationsModel.Create(org)
		Expect(err).To(BeNil())
		err = OrganizationsModel.Create(org2)
		Expect(err).To(BeNil())
	}

	BeforeEach(func() {
		cleanMainDatabaseTables()
		setup()
	})

	Context("Create", func() {
		It("should successfully create an organization", create)
	})

	Context("Get", func() {
		BeforeEach(create)

		It("should be able to find the first organization", func() {
			existingOrganization, err := OrganizationsModel.Get(org.ID)
			Expect(err).To(BeNil())
			Expect(existingOrganization.Name).To(Equal(org.Name))
		})

		It("should be able to find the second organization", func() {
			existingOrganization, err := OrganizationsModel.Get(org2.ID)
			Expect(err).To(BeNil())
			Expect(existingOrganization.Name).To(Equal(org2.Name))
		})

	})

	Context("Find", func() {
		BeforeEach(create)

		It("should return 2 organizations with associated data", func() {
			organizations, err := OrganizationsModel.Find()
			Expect(err).To(BeNil())
			Expect(len(organizations)).To(Equal(2))
		})

	})

	Context("Update", func() {
		BeforeEach(create)

		It("should update a organization", func() {
			org.Name = "SOME NEW VAL"
			err := OrganizationsModel.Update(org)
			Expect(err).To(BeNil())
			existingOrganization, err := OrganizationsModel.Get(org.ID)
			Expect(err).To(BeNil())
			Expect(existingOrganization.Name).To(Equal(org.Name))
		})

	})

	Context("Delete", func() {
		BeforeEach(create)

		It("should update a organization", func() {
			err := OrganizationsModel.Delete(org.ID)
			Expect(err).To(BeNil())
			organizations, err := OrganizationsModel.Find()
			Expect(err).To(BeNil())
			Expect(len(organizations)).To(Equal(1))
		})

	})

})
