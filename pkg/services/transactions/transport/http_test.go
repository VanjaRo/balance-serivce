package transport

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/VanjaRo/balance-serivce/pkg/errors"
	servmocks "github.com/VanjaRo/balance-serivce/pkg/mocks/services"
	"github.com/VanjaRo/balance-serivce/pkg/services/transactions"
	"github.com/VanjaRo/balance-serivce/pkg/services/users"

	"github.com/gin-gonic/gin"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHandlerDeposit(t *testing.T) {
	tests := map[string]struct {
		mockTransactionsService transactions.Service
		mockUsersService        users.Service
		uri                     string
		response                interface{}
		body                    string
		status                  int
	}{
		"User does not exist": {
			mockTransactionsService: &servmocks.MockTransactionsService{
				DepositErr: nil,
			},
			mockUsersService: &servmocks.MockUsersService{
				GetResult:    users.User{},
				GetErr:       users.ErrUserNotFound,
				CreateResult: "1",
				CreateErr:    nil,
			},
			uri: "/transactions/deposit",
			response: gin.H{
				"status": "ok",
			},
			body:   `{"user_id": "1", "amount": 100}`,
			status: http.StatusOK,
		},
		"User already exists": {
			mockTransactionsService: &servmocks.MockTransactionsService{
				DepositErr: nil,
			},
			mockUsersService: &servmocks.MockUsersService{
				GetResult: users.User{Id: "1", Balance: 100},
				GetErr:    nil,
				UpdateErr: nil,
			},
			uri: "/transactions/deposit",
			response: gin.H{
				"status": "ok",
			},
			body:   `{"user_id": "1", "amount": 100}`,
			status: http.StatusOK,
		},

		// "Server error": {
		// 	mockTransactionsService: &MockTransactionsService{
		// 		GetResult: users.User{},
		// 		GetErr:    fmt.Errorf("internal"),
		// 	},
		// 	mockUsersService: &users.mockUsersService{
		// 		GetResult: users.User{},
		// 		GetErr:    fmt.Errorf("internal"),
		// 	},
		// 	uri: fmt.Sprintf("/users/%s", id),
		// 	response: errors.AppError{
		// 		Code:        errors.InternalServerError,
		// 		Description: "internal",
		// 		Field:       "unknown",
		// 	},
		// 	status: http.StatusInternalServerError,
		// },
	}

	for testName, test := range tests {
		t.Run(testName, func(t *testing.T) {
			response := httptest.NewRecorder()
			router := gin.New()
			newTransactionHandler(router, test.mockUsersService, test.mockTransactionsService)

			req, err := http.NewRequest(http.MethodPost, test.uri, strings.NewReader(test.body))
			require.NoError(t, err)
			req.Header.Add("Content-Type", "application/json")

			router.ServeHTTP(response, req)

			assert.Equal(t, test.status, response.Code)

			if test.status == http.StatusOK {
				statusMessage := gin.H{}
				if err := json.Unmarshal(response.Body.Bytes(), &statusMessage); err != nil {
					assert.Fail(t, "failed to unmarshal", response.Body.String(), err)
				}
				assert.Equal(t, test.response, statusMessage)
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
