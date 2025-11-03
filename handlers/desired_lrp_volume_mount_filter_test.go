package handlers_test

import (
	"net/http"
	"net/http/httptest"

	"code.cloudfoundry.org/bbs/models"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("DesiredLRP VolumeMountDriver Filter", func() {
	var (
		responseRecorder *httptest.ResponseRecorder
		requestBody      interface{}
	)

	BeforeEach(func() {
		responseRecorder = httptest.NewRecorder()
		desiredLRP1 := models.DesiredLRP{}
		desiredLRP2 := models.DesiredLRP{}
		fakeDesiredLRPDB.DesiredLRPsReturns([]*models.DesiredLRP{&desiredLRP1, &desiredLRP2}, nil)
	})

	Describe("DesiredLRPs_r2 with VolumeMountDriver filter", func() {
		JustBeforeEach(func() {
			request := newTestRequest(requestBody)
			request.Header.Set("X-Request-Id", "test-request-id")
			handler.DesiredLRPs_r2(logger, responseRecorder, request)
		})

		Context("when filtering by volume mount driver", func() {
			BeforeEach(func() {
				requestBody = &models.DesiredLRPsRequest{VolumeMountDriver: "test-driver"}
			})

			It("passes the volume mount driver filter to the DB", func() {
				Expect(fakeDesiredLRPDB.DesiredLRPsCallCount()).To(Equal(1))
				_, _, filter := fakeDesiredLRPDB.DesiredLRPsArgsForCall(0)
				Expect(filter.VolumeMountDriver).To(Equal("test-driver"))
			})

			It("returns a successful response", func() {
				Expect(responseRecorder.Code).To(Equal(http.StatusOK))
			})
		})

		Context("when filtering by multiple filters including volume mount driver", func() {
			BeforeEach(func() {
				requestBody = &models.DesiredLRPsRequest{
					Domain:            "test-domain",
					ProcessGuids:      []string{"guid-1", "guid-2"},
					VolumeMountDriver: "test-driver",
				}
			})

			It("passes all filters to the DB", func() {
				Expect(fakeDesiredLRPDB.DesiredLRPsCallCount()).To(Equal(1))
				_, _, filter := fakeDesiredLRPDB.DesiredLRPsArgsForCall(0)
				Expect(filter.Domain).To(Equal("test-domain"))
				Expect(filter.ProcessGuids).To(Equal([]string{"guid-1", "guid-2"}))
				Expect(filter.VolumeMountDriver).To(Equal("test-driver"))
			})
		})
	})

	Describe("DesiredLRPs with VolumeMountDriver filter", func() {
		JustBeforeEach(func() {
			request := newTestRequest(requestBody)
			request.Header.Set("X-Request-Id", "test-request-id")
			handler.DesiredLRPs(logger, responseRecorder, request)
		})

		Context("when filtering by volume mount driver", func() {
			BeforeEach(func() {
				requestBody = &models.DesiredLRPsRequest{VolumeMountDriver: "test-driver"}
			})

			It("passes the volume mount driver filter to the DB", func() {
				Expect(fakeDesiredLRPDB.DesiredLRPsCallCount()).To(Equal(1))
				_, _, filter := fakeDesiredLRPDB.DesiredLRPsArgsForCall(0)
				Expect(filter.VolumeMountDriver).To(Equal("test-driver"))
			})

			It("returns a successful response", func() {
				Expect(responseRecorder.Code).To(Equal(http.StatusOK))
			})
		})
	})

	Describe("DesiredLRPSchedulingInfos with VolumeMountDriver filter", func() {
		BeforeEach(func() {
			fakeDesiredLRPDB.DesiredLRPSchedulingInfosReturns([]*models.DesiredLRPSchedulingInfo{}, nil)
		})

		JustBeforeEach(func() {
			request := newTestRequest(requestBody)
			request.Header.Set("X-Request-Id", "test-request-id")
			handler.DesiredLRPSchedulingInfos(logger, responseRecorder, request)
		})

		Context("when filtering by volume mount driver", func() {
			BeforeEach(func() {
				requestBody = &models.DesiredLRPsRequest{VolumeMountDriver: "test-driver"}
			})

			It("passes the volume mount driver filter to the DB", func() {
				Expect(fakeDesiredLRPDB.DesiredLRPSchedulingInfosCallCount()).To(Equal(1))
				_, _, filter := fakeDesiredLRPDB.DesiredLRPSchedulingInfosArgsForCall(0)
				Expect(filter.VolumeMountDriver).To(Equal("test-driver"))
			})
		})
	})

	Describe("DesiredLRPRoutingInfos with VolumeMountDriver filter", func() {
		BeforeEach(func() {
			fakeDesiredLRPDB.DesiredLRPRoutingInfosReturns([]*models.DesiredLRP{}, nil)
		})

		JustBeforeEach(func() {
			request := newTestRequest(requestBody)
			handler.DesiredLRPRoutingInfos(logger, responseRecorder, request)
		})

		Context("when filtering by volume mount driver", func() {
			BeforeEach(func() {
				requestBody = &models.DesiredLRPsRequest{VolumeMountDriver: "test-driver"}
			})

			It("passes the volume mount driver filter to the DB", func() {
				Expect(fakeDesiredLRPDB.DesiredLRPRoutingInfosCallCount()).To(Equal(1))
				_, _, filter := fakeDesiredLRPDB.DesiredLRPRoutingInfosArgsForCall(0)
				Expect(filter.VolumeMountDriver).To(Equal("test-driver"))
			})
		})
	})
})