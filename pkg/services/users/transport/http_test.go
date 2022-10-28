package transport

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/VanjaRo/balance-serivce/pkg/errors"
	servmocs "github.com/VanjaRo/balance-serivce/pkg/mocks/services"
	"github.com/VanjaRo/balance-serivce/pkg/services/users"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHandlerGet(t *testing.T) {
	id := uuid.New().String()
	balance := 100.0
	tests := map[string]struct {
		mockUsersService users.Service
		uri              string
		response         interface{}
		status           int
	}{
		"Happy path": {
			mockUsersService: &servmocs.MockUsersService{
				GetResult: users.User{Id: id, Balance: balance},
				GetErr:    nil,
			},
			uri: fmt.Sprintf("/users/%s", id),
			response: users.User{
				Id:      id,
				Balance: balance,
			},
			status: http.StatusOK,
		},
		"Not found": {
			mockUsersService: &servmocs.MockUsersService{
				GetResult: users.User{},
				GetErr:    users.ErrUserNotFound,
			},
			uri: fmt.Sprintf("/users/%s", id),
			response: errors.AppError{
				Code:        errors.NotFound,
				Description: users.ErrUserNotFound.Error(),
				Field:       "id",
			},
			status: http.StatusNotFound,
		},
		"Server error": {
			mockUsersService: &servmocs.MockUsersService{
				GetResult: users.User{},
				GetErr:    fmt.Errorf("internal"),
			},
			uri: fmt.Sprintf("/users/%s", id),
			response: errors.AppError{
				Code:        errors.InternalServerError,
				Description: "internal",
				Field:       "unknown",
			},
			status: http.StatusInternalServerError,
		},
	}

	for testName, test := range tests {
		t.Run(testName, func(t *testing.T) {
			response := httptest.NewRecorder()
			router := gin.New()
			newUsersHandler(router, test.mockUsersService)

			req, err := http.NewRequest(http.MethodGet, test.uri, nil)
			require.NoError(t, err)
			req.Header.Add("Content-Type", "application/json")

			router.ServeHTTP(response, req)

			assert.Equal(t, test.status, response.Code)

			if test.status == http.StatusOK {
				var ur users.User
				if err := json.Unmarshal(response.Body.Bytes(), &ur); err != nil {
					assert.Fail(t, "failed to unmarshal", response.Body.String(), err)
				}
				assert.Equal(t, test.response, ur)
			} else {
				var err errors.AppError
				if err := json.Unmarshal(response.Body.Bytes(), &err); err != nil {
					assert.Fail(t, "failed to unmarshal", response.Body.String(), err)
				}
				assert.Equal(t, test.response, err)
			}
		})
	}
}

func TestHandlerGetBalance(t *testing.T) {
	id := uuid.New().String()
	balance := 100.0

	tests := map[string]struct {
		mockUsersService users.Service
		uri              string
		response         interface{}
		status           int
	}{
		"Happy path": {
			mockUsersService: &servmocs.MockUsersService{
				GetBalanceResult: balance,
				GetBalanceError:  nil,
			},
			uri:      fmt.Sprintf("/users/%s/balance", id),
			response: balance,
			status:   http.StatusOK,
		},
		"Not found": {
			mockUsersService: &servmocs.MockUsersService{
				GetBalanceResult: 0,
				GetBalanceError:  users.ErrUserNotFound,
			},
			uri: fmt.Sprintf("/users/%s/balance", id),
			response: errors.AppError{
				Code:        errors.NotFound,
				Description: users.ErrUserNotFound.Error(),
				Field:       "id",
			},
			status: http.StatusNotFound,
		},
		"Server error": {
			mockUsersService: &servmocs.MockUsersService{
				GetBalanceResult: 0,
				GetBalanceError:  fmt.Errorf("internal"),
			},
			uri: fmt.Sprintf("/users/%s/balance", id),
			response: errors.AppError{
				Code:        errors.InternalServerError,
				Description: "internal",
				Field:       "unknown",
			},
			status: http.StatusInternalServerError,
		},
	}

	for testName, test := range tests {
		t.Run(testName, func(t *testing.T) {
			response := httptest.NewRecorder()
			router := gin.New()
			newUsersHandler(router, test.mockUsersService)

			req, err := http.NewRequest(http.MethodGet, test.uri, nil)
			require.NoError(t, err)
			req.Header.Add("Content-Type", "application/json")

			router.ServeHTTP(response, req)

			assert.Equal(t, test.status, response.Code)

			if test.status == http.StatusOK {
				var userBalance float64
				if err := json.Unmarshal(response.Body.Bytes(), &userBalance); err != nil {
					assert.Fail(t, "failed to unmarshal", response.Body.String(), err)
				}
				assert.Equal(t, test.response, userBalance)
			} else {
				var err errors.AppError
				if err := json.Unmarshal(response.Body.Bytes(), &err); err != nil {
					assert.Fail(t, "failed to unmarshal", response.Body.String(), err)
				}
				assert.Equal(t, test.response, err)
			}
		})
	}
}
