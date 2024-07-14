package cmdb

import (
	"github.com/gin-gonic/gin"

	"github.com/zhaohaihang/k8s-manage/dto"
	"github.com/zhaohaihang/k8s-manage/middleware"
	v1 "github.com/zhaohaihang/k8s-manage/pkg/core/kubemanage/v1"
	"github.com/zhaohaihang/k8s-manage/pkg/globalError"
)

// @Description  创建秘钥
// @ID           /api/cmdb/createSecret
// @Tags         Secret
// @Summary   创建主机组
// @Accept    application/json
// @Produce   application/json
// @Param     data  body      dto.CMDBSecretCreateInput 	true 	"秘钥信息"
// @Security  ApiKeyAuth
// @Success   200   {object}  middleware.Response{msg=string}  "创建秘钥"
// @Router    /api/cmdb/createSecret [post]
func (c *cmdbController) CreateSecret(ctx *gin.Context) {
	params := &dto.CMDBSecretCreateInput{}
	if err := params.BindingValidParams(ctx); err != nil {
		v1.Log.ErrorWithCode(globalError.ParamBindError, err)
		middleware.ResponseError(ctx, globalError.NewGlobalError(globalError.ParamBindError, err))
		return
	}
	if err := v1.CoreV1.CMDB().Secret().CreateSecret(ctx, params); err != nil {
		v1.Log.ErrorWithCode(globalError.CreateError, err)
		middleware.ResponseError(ctx, globalError.NewGlobalError(globalError.CreateError, err))
		return
	}
	middleware.ResponseSuccess(ctx, "")
}

// @Description  更新秘钥
// @ID           /api/cmdb/updateSecret
// @Tags         Secret
// @Summary   更新秘钥
// @Accept    application/json
// @Produce   application/json
// @Param     data  body      dto.CMDBSecretUpdateInput 	true 	"秘钥信息"
// @Security  ApiKeyAuth
// @Success   200   {object}  middleware.Response{msg=string}  "更新秘钥"
// @Router    /api/cmdb/updateSecret [put]
func (c *cmdbController) UpdateSecret(ctx *gin.Context) {
	params := &dto.CMDBSecretUpdateInput{}
	if err := params.BindingValidParams(ctx); err != nil {
		v1.Log.ErrorWithCode(globalError.ParamBindError, err)
		middleware.ResponseError(ctx, globalError.NewGlobalError(globalError.ParamBindError, err))
		return
	}
	if err := v1.CoreV1.CMDB().Secret().UpdateSecret(ctx, params); err != nil {
		v1.Log.ErrorWithCode(globalError.UpdateError, err)
		middleware.ResponseError(ctx, globalError.NewGlobalError(globalError.UpdateError, err))
		return
	}
	middleware.ResponseSuccess(ctx, "")
}

// @Description  秘钥分页
// @ID       /api/cmdb/pageSecret
// @Tags      Secret
// @Summary   秘钥分页
// @Security  ApiKeyAuth
// @Produce   application/json
// @Success   200   {object}  middleware.Response{data=dto.PageCMDBSecretOut,msg=string}  "秘钥分页"
// @Router    /api/cmdb/pageSecret [get]
func (c *cmdbController) PageSecret(ctx *gin.Context) {
	params := &dto.PageListCMDBSecretInput{}
	if err := params.BindingValidParams(ctx); err != nil {
		v1.Log.ErrorWithCode(globalError.ParamBindError, err)
		middleware.ResponseError(ctx, globalError.NewGlobalError(globalError.ParamBindError, err))
		return
	}
	data, err := v1.CoreV1.CMDB().Secret().PageSecret(ctx, params)
	if err != nil {
		v1.Log.ErrorWithCode(globalError.GetError, err)
		middleware.ResponseError(ctx, globalError.NewGlobalError(globalError.GetError, err))
		return
	}
	middleware.ResponseSuccess(ctx, data)
}

// @Description  获取秘钥列表
// @ID       /api/cmdb/getHostsList
// @Tags      Secret
// @Summary   获取秘钥列表
// @Security  ApiKeyAuth
// @Produce   application/json
// @Success   200   {object}  middleware.Response{data=[]model.CMDBSecret,msg=string}  "获取秘钥列表"
// @Router    /api/cmdb/getHostsList [get]
func (c *cmdbController) GetSecretList(ctx *gin.Context) {
	data, err := v1.CoreV1.CMDB().Secret().GetSecretList(ctx)
	if err != nil {
		v1.Log.ErrorWithCode(globalError.GetError, err)
		middleware.ResponseError(ctx, globalError.NewGlobalError(globalError.GetError, err))
		return
	}
	middleware.ResponseSuccess(ctx, data)
}

// @Description  删除秘钥
// @ID           /api/cmdb/:instanceID/deleteSecret
// @Tags         Secret
// @Summary   删除秘钥
// @Produce   application/json
// @Param     data  body       dto.Empty    true  "空"
// @Security  ApiKeyAuth
// @Success   200   {object}  middleware.Response{msg=string}  "删除秘钥"
// @Router    /api/cmdb/:instanceID/deleteSecret [delete]
func (c *cmdbController) DeleteSecret(ctx *gin.Context) {
	instanceid := ctx.Param("instanceID")
	if err := v1.CoreV1.CMDB().Secret().DeleteSecret(ctx, instanceid); err != nil {
		v1.Log.ErrorWithCode(globalError.DeleteError, err)
		middleware.ResponseError(ctx, globalError.NewGlobalError(globalError.DeleteError, err))
		return
	}
	middleware.ResponseSuccess(ctx, "")
}

// @Description  批量删除秘钥
// @ID           /api/cmdb/deleteSecrets
// @Tags         Secret
// @Summary   批量删除秘钥
// @Produce   application/json
// @Param     data  body       dto.Empty    true  "空"
// @Security  ApiKeyAuth
// @Success   200   {object}  middleware.Response{msg=string}  "批量删除秘钥"
// @Router    /api/cmdb/deleteSecrets [delete]
func (c *cmdbController) DeleteSecrets(ctx *gin.Context) {
	params := &dto.InstancesReq{}
	if err := params.BindingValidParams(ctx); err != nil {
		v1.Log.ErrorWithCode(globalError.ParamBindError, err)
		middleware.ResponseError(ctx, globalError.NewGlobalError(globalError.ParamBindError, err))
		return
	}
	if err := v1.CoreV1.CMDB().Secret().DeleteSecrets(ctx, params.Ids); err != nil {
		v1.Log.ErrorWithCode(globalError.DeleteError, err)
		middleware.ResponseError(ctx, globalError.NewGlobalError(globalError.DeleteError, err))
		return
	}
	middleware.ResponseSuccess(ctx, "")
}
