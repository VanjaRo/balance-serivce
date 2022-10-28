package transport

import (
	"fmt"
	"net/http"

	"github.com/VanjaRo/balance-serivce/pkg/errors"
	"github.com/VanjaRo/balance-serivce/pkg/services/transactions"
	transStore "github.com/VanjaRo/balance-serivce/pkg/services/transactions/store"
	"github.com/VanjaRo/balance-serivce/pkg/services/users"
	userStore "github.com/VanjaRo/balance-serivce/pkg/services/users/store"
	"github.com/VanjaRo/balance-serivce/pkg/utils/context"
	"github.com/VanjaRo/balance-serivce/pkg/utils/log"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type handler struct {
	UserService        users.Service
	TransactionService transactions.Service
}

func ActivateHandlers(router *gin.Engine, db *gorm.DB) {
	userService := users.NewUserService(userStore.NewUserRepo(db))
	transactionService := transactions.NewTransactionService(transStore.NewTransactionRepo(db))
	newTransactionHandler(router, userService, transactionService)
}

func newTransactionHandler(router *gin.Engine, userService users.Service, transactionService transactions.Service) {
	handler := handler{
		UserService:        userService,
		TransactionService: transactionService,
	}

	router.POST("/transactions/deposit", handler.Deposit)
	router.POST("/transactions/freeze", handler.Freeze)
	router.POST("/transactions/apply", handler.Apply)
}

func (h *handler) Deposit(rCtx *gin.Context) {
	var q struct {
		UserId string  `json:"user_id" binding:"required"`
		Amount float64 `json:"amount" binding:"required"`
	}
	ctx := context.GetReqCtx(rCtx)

	if err := rCtx.BindJSON(&q); err != nil {
		log.Info(ctx, "query parse error: %s", err.Error())
		rCtx.IndentedJSON(http.StatusBadRequest, errors.NewAppError(errors.BadRequest, errors.Desctiptions[errors.BadRequest], ""))
		return
	}
	fmt.Printf("q: %+v", q)
	// check if user exists
	_, err := h.UserService.Get(ctx, q.UserId)
	if err != nil {
		// if user does not exist, create one
		_, err := h.UserService.Create(ctx, q.UserId, q.Amount)
		if err != nil {
			status, appErr := handleError(err)
			rCtx.IndentedJSON(status, appErr)
			return
		}
	} else {
		// create transaction
		err = h.TransactionService.Deposit(ctx, q.UserId, q.Amount)
		if err != nil {
			status, appErr := handleError(err)
			rCtx.IndentedJSON(status, appErr)
			return
		}
		// if user exists, update balance
		err := h.UserService.UpdateUserBalance(ctx, q.UserId, q.Amount)
		if err != nil {
			status, appErr := handleError(err)
			rCtx.IndentedJSON(status, appErr)
			return
		}
	}

	rCtx.IndentedJSON(http.StatusOK, gin.H{"status": "ok"})
}

func (h *handler) Freeze(rCtx *gin.Context) {
	var q struct {
		UserId    string  `json:"user_id" binding:"required"`
		OrderId   string  `json:"order_id" binding:"required"`
		ServiceId string  `json:"service_id" binding:"required"`
		Amount    float64 `json:"amount" binding:"required"`
	}
	ctx := context.GetReqCtx(rCtx)

	if err := rCtx.BindJSON(&q); err != nil {
		log.Info(ctx, "query parse error: %s", err.Error())
		rCtx.IndentedJSON(http.StatusBadRequest, errors.NewAppError(errors.BadRequest, errors.Desctiptions[errors.BadRequest], ""))
		return
	}
	fmt.Printf("q: %+v", q)
	// check if user exists
	_, err := h.UserService.Get(ctx, q.UserId)
	if err != nil {
		// if user does not exist return error
		status, appErr := handleError(err)
		rCtx.IndentedJSON(status, appErr)
		return
	} else {
		// create transaction
		err = h.TransactionService.Freeze(ctx, q.UserId, q.OrderId, q.ServiceId, q.Amount)
		if err != nil {
			status, appErr := handleError(err)
			rCtx.IndentedJSON(status, appErr)
			return
		}
		// if user exists, update balance
		err := h.UserService.UpdateUserBalance(ctx, q.UserId, -q.Amount)
		if err != nil {
			status, appErr := handleError(err)
			rCtx.IndentedJSON(status, appErr)
			return
		}
	}

	rCtx.IndentedJSON(http.StatusOK, gin.H{"status": "ok"})
}

func (h *handler) Apply(rCtx *gin.Context) {
	var q struct {
		UserId    string  `json:"user_id" binding:"required"`
		OrderId   string  `json:"order_id" binding:"required"`
		ServiceId string  `json:"service_id" binding:"required"`
		Amount    float64 `json:"amount" binding:"required"`
	}
	ctx := context.GetReqCtx(rCtx)

	if err := rCtx.BindJSON(&q); err != nil {
		log.Info(ctx, "query parse error: %s", err.Error())
		rCtx.IndentedJSON(http.StatusBadRequest, errors.NewAppError(errors.BadRequest, errors.Desctiptions[errors.BadRequest], ""))
		return
	}
	fmt.Printf("q: %+v", q)
	// check if user exists
	_, err := h.UserService.Get(ctx, q.UserId)
	if err != nil {
		// if user does not exist return error
		status, appErr := handleError(err)
		rCtx.IndentedJSON(status, appErr)
		return
	} else {
		// create transaction
		err = h.TransactionService.Apply(ctx, q.UserId, q.OrderId, q.ServiceId, q.Amount)
		if err != nil {
			status, appErr := handleError(err)
			rCtx.IndentedJSON(status, appErr)
			return
		}
	}

	rCtx.IndentedJSON(http.StatusOK, gin.H{"status": "ok"})
}

func handleError(e error) (int, error) {
	switch e {
	case users.ErrUserNotFound:
		return http.StatusNotFound, errors.NewAppError(errors.NotFound, e.Error(), "id")
	case users.ErrUserCreate:
		return http.StatusInternalServerError, errors.NewAppError(errors.InternalServerError, "unable to create user", "")
	default:
		return http.StatusInternalServerError, errors.NewAppError(errors.InternalServerError, e.Error(), "unknown")
	}
}
