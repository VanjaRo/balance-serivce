package transport

import (
	"net/http"

	"github.com/VanjaRo/balance-serivce/pkg/errors"
	"github.com/VanjaRo/balance-serivce/pkg/services/users"
	"github.com/VanjaRo/balance-serivce/pkg/services/users/store"
	"github.com/VanjaRo/balance-serivce/pkg/utils/context"
	"github.com/VanjaRo/balance-serivce/pkg/utils/log"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type handler struct {
	UserService users.Service
}

func ActivateHandlers(router *gin.Engine, db *gorm.DB) {
	userService := users.NewUserService(store.NewUserRepo(db))
	newHandler(router, userService)
}

func newHandler(router *gin.Engine, userService users.Service) {
	handler := handler{
		UserService: userService,
	}
	router.GET("/users", handler.GetAllUsers)
	router.GET("/users/:id", handler.GetUser)
	router.GET("/users/:id/balance", handler.GetUserBalance)
}

func (h *handler) GetUser(rCtx *gin.Context) {
	ctx := context.GetReqCtx(rCtx)

	log.Info(ctx, "retrieving user id=%s", rCtx.Param("id"))
	user, err := h.UserService.Get(ctx, rCtx.Param("id"))
	if err != nil {
		status, appErr := handleError(err)
		rCtx.IndentedJSON(status, appErr)
		return
	}

	rCtx.IndentedJSON(http.StatusOK, user)
}

func (h *handler) GetAllUsers(rCtx *gin.Context) {
	var q struct {
		Limit  int `form:"limit,default=25"`
		Offset int `form:"offset,default=0"`
	}
	ctx := context.GetReqCtx(rCtx)

	if err := rCtx.BindQuery(&q); err != nil {
		rCtx.IndentedJSON(http.StatusBadRequest, errors.NewAppError(errors.BadRequest, errors.Desctiptions[errors.BadRequest], ""))
		return
	}

	usrs, err := h.UserService.GetAll(ctx, q.Limit, q.Offset)
	if err != nil {
		status, appErr := handleError(err)
		rCtx.IndentedJSON(status, appErr)
		return
	}
	rCtx.IndentedJSON(http.StatusOK, users.Users{Users: usrs})
}

func (h *handler) GetUserBalance(rCtx *gin.Context) {
	ctx := context.GetReqCtx(rCtx)

	balance, err := h.UserService.GetBalance(ctx, rCtx.Param("id"))
	if err != nil {
		status, appErr := handleError(err)
		rCtx.IndentedJSON(status, appErr)
		return
	}

	rCtx.IndentedJSON(http.StatusOK, balance)
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
