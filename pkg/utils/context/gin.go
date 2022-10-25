package context

import (
	"context"

	"github.com/gin-gonic/gin"
)

const ctxKey = "ctxKey"

func GetReqCtx(ctx *gin.Context) context.Context {
	rCtx, exists := ctx.Get(ctxKey)
	if !exists {
		return context.Background()
	}
	return rCtx.(context.Context)
}

func SetReqCtx(ctx context.Context, rCtx *gin.Context) {
	rCtx.Set(ctxKey, ctx)
}
