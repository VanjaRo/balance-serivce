package transport

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/VanjaRo/balance-serivce/pkg/errors"
	"github.com/VanjaRo/balance-serivce/pkg/services/users"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type mockService struct {
	GetResult users.User
	GetErr    error

	GetAllResult []users.User
	GetAllErr    error

	GetBalanceResult float64
	GetBalanceError  error

	CreateResult string
	CreateErr    error
}

func (s *mockService) Get(ctx context.Context, id string) (users.User, error) {
	return s.GetResult, s.GetErr
}

func (s *mockService) GetAll(ctx context.Context, limit, offset int) ([]users.User, error) {
	return s.GetAllResult, s.GetAllErr
}

func (s *mockService) Create(ctx context.Context, u users.User) (string, error) {
	return s.CreateResult, s.CreateErr
}

func (s *mockService) GetBalance(ctx context.Context, id string) (float64, error) {
	return s.GetBalanceResult, s.GetBalanceError
}

func TestHandlerGet(t *testing.T) {
	id := uuid.New().String()
	balance := 100.0
	tests := map[string]struct {
		mockService users.Service
		uri         string
		response    interface{}
		status      int
	}{
		"Happy path": {
			mockService: &mockService{
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
			mockService: &mockService{
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
			mockService: &mockService{
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
			newHandler(router, test.mockService)

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
		mockService users.Service
		uri         string
		response    interface{}
		status      int
	}{
		"Happy path": {
			mockService: &mockService{
				GetBalanceResult: balance,
				GetBalanceError:  nil,
			},
			uri:      fmt.Sprintf("/users/%s/balance", id),
			response: balance,
			status:   http.StatusOK,
		},
		"Not found": {
			mockService: &mockService{
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
			mockService: &mockService{
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
			newHandler(router, test.mockService)

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
