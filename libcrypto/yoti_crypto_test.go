package libcrypto_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"bytes"
	"crypto/rand"
	. "github.com/bilbercode/encryption-server/libcrypto"
)

var _ = Describe("YotiCrypto", func() {

	Describe("NewYotiCrypto()", func() {

		Context("with a valid rand source", func() {
			It("should return a new YotiCrypto", func() {
				crypto, err := NewYotiCrypto(rand.Reader)
				Expect(err).ToNot(HaveOccurred())
				Expect(crypto).To(BeAssignableToTypeOf(&YotiCrypto{}))
			})
		})

		Context("with an invalid rand source", func() {
			It("should return an error", func() {
				crypto, err := NewYotiCrypto(bytes.NewBuffer(nil))
				Expect(err).To(HaveOccurred())
				Expect(crypto).To(BeNil())
			})
		})
	})

	Describe("Encrypt()", func() {
		Context("configured with a proper rand source", func() {
			It("should return a populated EncryptionResult", func() {
				crypto, err := NewYotiCrypto(rand.Reader)
				Expect(err).ToNot(HaveOccurred())
				res, err := crypto.Encrypt([]byte("test"))
				Expect(err).ToNot(HaveOccurred())
				Expect(res).To(BeAssignableToTypeOf(EncryptionResult{}))
				Expect(res.Key).ToNot(BeNil())
				Expect(res.Key).ToNot(HaveLen(0))
				Expect(res.Data).ToNot(BeNil())
				Expect(res.Data).ToNot(HaveLen(0))
			})
		})

		Context("configured with a weak rand source", func() {
			It("should return an error", func() {
				crypto, err := NewYotiCrypto(bytes.NewBuffer([]byte{12, 1}))
				Expect(err).ToNot(HaveOccurred())
				_, err = crypto.Encrypt([]byte("test"))
				Expect(err).To(HaveOccurred())
			})
		})
	})

	Describe("Decrypt()", func() {
		Context("passed a valid encrypted byte slice and key", func() {
			It("should return the decrypted byte slice", func() {
				crypto, err := NewYotiCrypto(rand.Reader)
				Expect(err).ToNot(HaveOccurred())
				eRes, err := crypto.Encrypt([]byte("test"))
				Expect(err).ToNot(HaveOccurred())

				dRes, err := crypto.Decrypt(eRes.Data, eRes.Key)
				Expect(err).ToNot(HaveOccurred())
				Expect(string(dRes)).To(Equal("test"))
			})
		})

		Context("passed an invalid byte slice and key", func() {
			Context("data to short", func() {
				It("should return an error", func() {
					crypto, err := NewYotiCrypto(rand.Reader)
					Expect(err).ToNot(HaveOccurred())
					eRes, err := crypto.Encrypt([]byte("test"))
					Expect(err).ToNot(HaveOccurred())
					_, err = crypto.Decrypt([]byte("this should fail"), eRes.Key)
					Expect(err).To(HaveOccurred())
				})
			})

			Context("data wrong block size", func() {
				It("should return an error", func() {
					crypto, err := NewYotiCrypto(rand.Reader)
					Expect(err).ToNot(HaveOccurred())
					eRes, err := crypto.Encrypt([]byte("test"))
					Expect(err).ToNot(HaveOccurred())
					_, err = crypto.Decrypt([]byte("this should fail, this should fail, this should fail"), eRes.Key)
					Expect(err).To(HaveOccurred())
				})
			})

			Context("data just wrong", func() {
				It("should return an error", func() {
					crypto, err := NewYotiCrypto(rand.Reader)
					Expect(err).ToNot(HaveOccurred())
					eRes, err := crypto.Encrypt([]byte("test"))
					Expect(err).ToNot(HaveOccurred())
					_, err = crypto.Decrypt([]byte("this should fail, this should :("), eRes.Key)
					Expect(err).To(HaveOccurred())
				})
			})

		})

		Context("passed an invalid key and byte slice", func() {
			It("should return an error", func() {
				crypto, err := NewYotiCrypto(rand.Reader)
				Expect(err).ToNot(HaveOccurred())
				eRes, err := crypto.Encrypt([]byte("test"))
				Expect(err).ToNot(HaveOccurred())
				_, err = crypto.Decrypt(eRes.Data, []byte("foo bat"))
				Expect(err).To(HaveOccurred())
			})
		})
	})

	Describe("HashIdWithKey", func() {


		It("should not match the original composite", func() {
			crypto, err := NewYotiCrypto(rand.Reader)
			Expect(err).ToNot(HaveOccurred())
			id := "this is a test ID"
			key := []byte("everybody should really worry about the squirrels")
			oc := crypto.HashIdWithKey(id, key)
			comb := append([]byte(id), key...)
			Expect(oc).ToNot(Equal(comb))

		})
	})

})
