package system

import (
	"github.com/gin-gonic/gin"

	"github.com/zhaohaihang/k8s-manage/middleware"
	"github.com/zhaohaihang/k8s-manage/pkg/core/kubemanage/v1"
	"github.com/zhaohaihang/k8s-manage/pkg/globalError"
)

func (s *systemController) GetSystemState(ctx *gin.Context) {
	data, err := v1.CoreV1.System().SystemService().GetSystemState(ctx)
	if err != nil {
		v1.Log.ErrorWithCode(globalError.GetError, err)
		middleware.ResponseError(ctx, globalError.NewGlobalError(globalError.GetError, err))
		return
	}
	middleware.ResponseSuccess(ctx, data)
}
