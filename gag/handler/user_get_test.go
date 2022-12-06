package handler

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/plus100kt/goserver/gag/model"
	"github.com/plus100kt/goserver/gag/model/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/gin-gonic/gin"
)

func TestUserGet(t *testing.T) {
	// Setup
	gin.SetMode(gin.TestMode)

	t.Run("Success", func(t *testing.T) {
		uid := "1234"

		nowtime := time.Now()
		mockUserResp := &model.User{
			ID:            "1234",
			UUID:          "1234",
			RsaPrivateKey: "1234",
			AesPassword:   "1234",
			Name:          "1234",
			Email:         "1234",
			ImageURL:      "1234",
			CreatedAt:     nowtime,
			UpdatedAt:     nowtime,
		}
		mockUserService := new(mocks.MockUserService)
		mockUserService.On("Get", mock.AnythingOfType("*gin.Context"), uid).Return(mockUserResp, nil)

		// a response recorder for getting written http response
		rr := httptest.NewRecorder()

		// use a middleware to set context for test
		// the only claims we care about in this test
		// is the UID
		router := gin.Default()
		router.Use(func(c *gin.Context) {
			c.Set("user", &model.User{
				ID: uid,
			},
			)
		})

		NewHandler(&Config{
			R:           router,
			UserService: mockUserService,
		})

		request, err := http.NewRequest(http.MethodGet, "/v1/user/", nil)
		assert.NoError(t, err)

		router.ServeHTTP(rr, request)

		respBody, err := json.Marshal(gin.H{
			"user": mockUserResp,
		})
		assert.NoError(t, err)

		assert.Equal(t, 200, rr.Code)
		assert.Equal(t, respBody, rr.Body.Bytes())
		mockUserService.AssertExpectations(t) // assert that UserService.Get was called
	})
}
