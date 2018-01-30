package libstorage_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/bilbercode/encryption-server/libstorage"
)

var _ = Describe("LocalStorage", func() {

	Describe("Store()", func() {
		It("should store the value against the provided key", func() {
			storage := NewLocalStorage()
			err := storage.Store([]byte("test value"), "test key")
			Expect(err).ToNot(HaveOccurred())
		})
	})

	Describe("Has()", func() {
		It("should return true if the value exists", func() {
			storage := NewLocalStorage()
			storage.Store([]byte("test value"), "test key")

			has, _ := storage.Has("test key")
			Expect(has).To(BeTrue())
		})

		It("should return false if the value does not exist", func() {
			storage := NewLocalStorage()

			has, _ := storage.Has("test key")
			Expect(has).To(BeFalse())
		})
	})

	Describe("Retrieve()", func() {
		It("should return the value when specified the correct key", func() {
			storage := NewLocalStorage()
			storage.Store([]byte("test value"), "test key")

			value, _ := storage.Retrieve("test key")
			Expect(string(value)).To(Equal("test value"))
		})

		It("should return an error if the key does not exist", func() {
			storage := NewLocalStorage()

			value, err := storage.Retrieve("test key")
			Expect(value).To(BeNil())
			Expect(err).To(HaveOccurred())
		})
	})


	Describe("NewLocalStorage", func() {
		It("should return a new local storage entity", func() {
			storage := NewLocalStorage()
			Expect(storage).To(BeAssignableToTypeOf(&LocalStorage{}))
		})
	})
})
