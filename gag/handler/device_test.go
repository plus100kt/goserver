package handler

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/plus100kt/goserver/gag/model"
	"github.com/plus100kt/goserver/gag/model/mocks"
	"github.com/stretchr/testify/mock"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestDevice(t *testing.T) {
	// Setup
	gin.SetMode(gin.TestMode)

	t.Run("Success", func(t *testing.T) {
		uuid := "11"

		mockDeviceResp := &model.Device{
			UUID:          uuid,
			RsaPublicKey:  "1234",
			RsaPrivateKey: "5678",
		}

		mockDeviceService := new(mocks.MockDeviceService)
		mockDeviceService.On("Register", mock.AnythingOfType("*gin.Context"), uuid).Return(mockDeviceResp, nil)

		// a response recorder for getting written http response
		rr := httptest.NewRecorder()

		// use a middleware to set context for test
		router := gin.Default()
		router.Use(func(c *gin.Context) {
			c.Set("Device", &model.Device{
				UUID:          uuid,
				RsaPrivateKey: "",
				RsaPublicKey:  "",
			},
			)
		})

		NewHandler(&Config{
			R:             router,
			DeviceService: mockDeviceService,
		})

		request, err := http.NewRequest(http.MethodPost, "v1/device/register", nil)
		assert.NoError(t, err)

		router.ServeHTTP(rr, request)

		respBody, err := json.Marshal(gin.H{
			"Device": mockDeviceResp,
		})
		assert.NoError(t, err)

		assert.Equal(t, 200, rr.Code)
		assert.Equal(t, respBody, rr.Body.Bytes())
		mockDeviceService.AssertExpectations(t) // assert that DeviceService.Get was called
	})

	t.Run("NoContextDevice", func(t *testing.T) {
		mockDeviceService := new(mocks.MockDeviceService)
		mockDeviceService.On("Register", mock.Anything, mock.Anything).Return(nil, nil)

		// a response recorder for getting written http response
		rr := httptest.NewRecorder()

		// do not append Device to context
		router := gin.Default()
		NewHandler(&Config{
			R:             router,
			DeviceService: mockDeviceService,
		})

		request, err := http.NewRequest(http.MethodPost, "v1/device/register", nil)
		assert.NoError(t, err)

		router.ServeHTTP(rr, request)

		assert.Equal(t, 500, rr.Code)
		mockDeviceService.AssertNotCalled(t, "Post", mock.Anything)
	})
}
