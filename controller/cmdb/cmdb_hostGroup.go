package cmdb

import (
	"github.com/gin-gonic/gin"

	"github.com/zhaohaihang/k8s-manage/dto"
	"github.com/zhaohaihang/k8s-manage/middleware"
	v1 "github.com/zhaohaihang/k8s-manage/pkg/core/kubemanage/v1"
	"github.com/zhaohaihang/k8s-manage/pkg/globalError"
)

// @Description  创建主机组
// @ID           /api/cmdb/createHostGroup
// @Tags         HostGroup
// @Summary   创建主机组
// @Accept    application/json
// @Produce   application/json
// @Param     data  body      dto.HostGroupCreateOrUpdateInput 	true 	"主机组信息"
// @Security  ApiKeyAuth
// @Success   200   {object}  middleware.Response{msg=string}  "创建主机组"
// @Router    /api/cmdb/createHostGroup [post]
func (c *cmdbController) CreateHostGroup(ctx *gin.Context) {
	params := &dto.HostGroupCreateOrUpdateInput{}
	if err := params.BindingValidParams(ctx); err != nil {
		v1.Log.ErrorWithCode(globalError.ParamBindError, err)
		middleware.ResponseError(ctx, globalError.NewGlobalError(globalError.ParamBindError, err))
		return
	}
	err := v1.CoreV1.CMDB().HostGroup().CreateHostGroup(ctx, params)
	if err != nil {
		v1.Log.ErrorWithCode(globalError.CreateError, err)
		middleware.ResponseError(ctx, globalError.NewGlobalError(globalError.CreateError, err))
		return
	}
	middleware.ResponseSuccess(ctx, "创建成功")
}

// @Description  创建子主机组
// @ID           /api/cmdb/createSubHostGroup
// @Tags         HostGroup
// @Summary   创建子主机组
// @Accept    application/json
// @Produce   application/json
// @Param     data  body      dto.HostGroupCreateOrUpdateInput 	true 	"子主机组信息"
// @Security  ApiKeyAuth
// @Success   200   {object}  middleware.Response{msg=string}  "创建子主机组"
// @Router    /api/cmdb/createSubHostGroup [post]
func (c *cmdbController) CreateSubHostGroup(ctx *gin.Context) {
	params := &dto.HostGroupCreateOrUpdateInput{}
	if err := params.BindingValidParams(ctx); err != nil {
		v1.Log.ErrorWithCode(globalError.ParamBindError, err)
		middleware.ResponseError(ctx, globalError.NewGlobalError(globalError.ParamBindError, err))
		return
	}
	err := v1.CoreV1.CMDB().HostGroup().CreateSubHostGroup(ctx, params)
	if err != nil {
		v1.Log.ErrorWithCode(globalError.CreateError, err)
		middleware.ResponseError(ctx, globalError.NewGlobalError(globalError.CreateError, err))
		return
	}
	middleware.ResponseSuccess(ctx, "创建成功")
}

// @Description  更新主机组
// @ID           /api/cmdb/updateHostGroup
// @Tags         HostGroup
// @Summary   更新主机组
// @Accept    application/json
// @Produce   application/json
// @Param     data  body      dto.HostGroupCreateOrUpdateInput 	true 	"主机组信息"
// @Security  ApiKeyAuth
// @Success   200   {object}  middleware.Response{msg=string}  "更新主机组"
// @Router    /api/cmdb/updateHostGroup [put]
func (c *cmdbController) UpdateHostGroup(ctx *gin.Context) {
	params := &dto.HostGroupCreateOrUpdateInput{}
	if err := params.BindingValidParams(ctx); err != nil {
		v1.Log.ErrorWithCode(globalError.ParamBindError, err)
		middleware.ResponseError(ctx, globalError.NewGlobalError(globalError.ParamBindError, err))
		return
	}
	err := v1.CoreV1.CMDB().HostGroup().UpdateHostGroup(ctx, params)
	if err != nil {
		v1.Log.ErrorWithCode(globalError.UpdateError, err)
		middleware.ResponseError(ctx, globalError.NewGlobalError(globalError.UpdateError, err))
		return
	}
	middleware.ResponseSuccess(ctx, "更新成功")
}

// @Description  删除主机组
// @ID           /api/cmdb/:instanceID/deleteHostGroup
// @Tags         HostGroup
// @Summary   删除主机组
// @Produce   application/json
// @Param     data  body       dto.Empty    true  "空"
// @Security  ApiKeyAuth
// @Success   200   {object}  middleware.Response{msg=string}  "更新主机组"
// @Router    /api/cmdb/:instanceID/deleteHostGroup [delete]
func (c *cmdbController) DeleteHostGroup(ctx *gin.Context) {
	instanceID := ctx.Param("instanceID")
	if err := v1.CoreV1.CMDB().HostGroup().DeleteHostGroup(ctx, instanceID); err != nil {
		v1.Log.ErrorWithCode(globalError.DeleteError, err)
		middleware.ResponseError(ctx, globalError.NewGlobalError(globalError.DeleteError, err))
		return
	}
	middleware.ResponseSuccess(ctx, "")
}

// @Description  获取主机组树
// @ID       /api/cmdb/getHostGroupTree
// @Tags      HostGroup
// @Summary   获取主机组树
// @Security  ApiKeyAuth
// @Produce   application/json
// @Success   200   {object}  middleware.Response{data=model.CMDBHostGroup,msg=string}  "获取主机组树"
// @Router    /api/cmdb/getHostGroupTree [get]
func (c *cmdbController) GetHostGroupTree(ctx *gin.Context) {
	data, err := v1.CoreV1.CMDB().HostGroup().GetHostGroupTree(ctx)
	if err != nil {
		v1.Log.ErrorWithCode(globalError.GetError, err)
		middleware.ResponseError(ctx, globalError.NewGlobalError(globalError.GetError, err))
		return
	}
	middleware.ResponseSuccess(ctx, data)
}

// @Description  获取主机组列表
// @ID       /api/cmdb/getHostGroupList
// @Tags      HostGroup
// @Summary   获取主机组列表
// @Security  ApiKeyAuth
// @Produce   application/json
// @Success   200   {object}  middleware.Response{data=model.CMDBHostGroup,msg=string}  "获取主机组列表"
// @Router    /api/cmdb/getHostGroupList [get]
func (c *cmdbController) GetHostGroupList(ctx *gin.Context) {
	data, err := v1.CoreV1.CMDB().HostGroup().GetHostGroupList(ctx)
	if err != nil {
		v1.Log.ErrorWithCode(globalError.GetError, err)
		middleware.ResponseError(ctx, globalError.NewGlobalError(globalError.GetError, err))
		return
	}
	middleware.ResponseSuccess(ctx, data)
}
