package user

import (
	"github.com/gin-gonic/gin"

	"github.com/zhaohaihang/k8s-manage/dto"
	"github.com/zhaohaihang/k8s-manage/middleware"
	"github.com/zhaohaihang/k8s-manage/pkg/core/kubemanage/v1"
	"github.com/zhaohaihang/k8s-manage/pkg/globalError"
	"github.com/zhaohaihang/k8s-manage/pkg/utils"
)

func (u *userController) GetDepartmentTree(ctx *gin.Context) {
	data, err := v1.CoreV1.System().Department().GetDepartmentTree(ctx)
	if err != nil {
		v1.Log.ErrorWithCode(globalError.GetError, err)
		middleware.ResponseError(ctx, globalError.NewGlobalError(globalError.GetError, err))
		return
	}
	middleware.ResponseSuccess(ctx, data)
}

func (u *userController) GetDepartmentUsers(ctx *gin.Context) {
	did, err := utils.ParseUint(ctx.Param("id"))
	if err != nil {
		v1.Log.ErrorWithCode(globalError.ParamBindError, err)
		middleware.ResponseError(ctx, globalError.NewGlobalError(globalError.ParamBindError, err))
	}
	data, err := v1.CoreV1.System().Department().GetDeptUsers(ctx, did)
	if err != nil {
		v1.Log.ErrorWithCode(globalError.GetError, err)
		middleware.ResponseError(ctx, globalError.NewGlobalError(globalError.GetError, err))
		return
	}
	middleware.ResponseSuccess(ctx, data)
}

func (u *userController) GetDeptByPage(ctx *gin.Context) {
	params := &dto.PageListDeptInput{}
	if err := params.BindingValidParams(ctx); err != nil {
		v1.Log.ErrorWithCode(globalError.ParamBindError, err)
		middleware.ResponseError(ctx, globalError.NewGlobalError(globalError.ParamBindError, err))
		return
	}
	data, err := v1.CoreV1.System().Department().PageList(ctx, *params)
	if err != nil {
		v1.Log.ErrorWithCode(globalError.GetError, err)
		middleware.ResponseError(ctx, globalError.NewGlobalError(globalError.GetError, err))
		return
	}
	middleware.ResponseSuccess(ctx, data)
}
