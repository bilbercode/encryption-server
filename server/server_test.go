package server_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/bilbercode/encryption-server/server"
	"github.com/bilbercode/encryption-server/server/handler"
	"github.com/bilbercode/yoti-storage-server/server"
	"github.com/golang/mock/gomock"
	"net/http"
	"net/http/httptest"
	"github.com/golang/protobuf/proto"
	"github.com/bilbercode/encryption-server/protobuf"
	"bytes"
	"io/ioutil"
)

var _ = Describe("Server", func() {

	var ts *httptest.Server
	var crypto *MockCrypto
	var cryptoController *gomock.Controller
	var storage *MockStorage
	var storageController *gomock.Controller

	BeforeSuite(func() {
		cryptoController = gomock.NewController(GinkgoT())
		storageController = gomock.NewController(GinkgoT())
		crypto = NewMockCrypto(cryptoController)
		storage = NewMockStorage(storageController)
		router, _ := NewRouter([]Route{
			{
				Name:    "store",
				Path:    "/v1/store",
				Method:  http.MethodPost,
				Handler: handler.NewStoreHandler(storage, crypto).HandleHttp,
			},
			{
				Name:    "retrieve",
				Path:    "/v1/retrieve",
				Method:  http.MethodPost,
				Handler: handler.NewRetrieveHandler(storage, crypto).HandleHttp,
			},
			{
				Name:    "ping",
				Path:    "/",
				Method:  http.MethodGet,
				Handler: handler.Ping,
			},
		})

		ts = httptest.NewServer(router)
	})

	AfterSuite(func() {
		server.Close()
	})


	AfterEach(func() {
		cryptoController.Finish()
		storageController.Finish()
	})

	Describe("GET /", func() {
		It("", func() {
			res, err := http.Get(ts.URL)
			Expect(err).ToNot(HaveOccurred())
			if err != nil {
				defer res.Body.Close()
			}
			Expect(res.StatusCode).To(Equal(200))
		})
	})

	Describe("POST /v1/retrieve", func() {
		It("should respond with the plaintext data", func() {

			crypto.EXPECT().HashIdWithKey("test", []byte("test")).
				Times(1).
				Return([]byte("1234567890"))

			storage.EXPECT().Retrieve("1234567890").
				Times(1).
				Return([]byte("some encrypted data"), nil)

			crypto.EXPECT().Decrypt([]byte("some encrypted data"), []byte("test")).
				Times(1).
				Return([]byte("test plaintext"), nil)

			reqPayload, err := proto.Marshal(&protobuf.RetrieveRequest{
				ID: "test",
				Key: []byte("test"),
			})

			if err != nil {
				panic(err)
			}


			res, err := http.Post(
				ts.URL + "/v1/retrieve",
				"application/octet-stream",
				bytes.NewBuffer(reqPayload),
			)
			Expect(err).To(BeNil())

			defer res.Body.Close()

			b, err := ioutil.ReadAll(res.Body)
			Expect(err).To(BeNil())

			var retrieveResponse = new(protobuf.RetrieveResponse)

			err = proto.Unmarshal(b, retrieveResponse)
			Expect(err).ToNot(HaveOccurred())

			Expect(retrieveResponse.Data).To(Equal([]byte("test plaintext")))

		})
	})

})
