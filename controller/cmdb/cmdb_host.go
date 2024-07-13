package cmdb

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/zhaohaihang/k8s-manage/pkg/utils"

	"github.com/zhaohaihang/k8s-manage/dto"
	"github.com/zhaohaihang/k8s-manage/middleware"
	v1 "github.com/zhaohaihang/k8s-manage/pkg/core/kubemanage/v1"
	"github.com/zhaohaihang/k8s-manage/pkg/globalError"
)

// @Description  创建主机
// @ID           /api/cmdb/createHost
// @Tags         Host
// @Summary   创建主机
// @Accept    application/json
// @Produce   application/json
// @Param     data  body      dto.CMDBHostCreateInput 	true 	"主机信息"
// @Security  ApiKeyAuth
// @Success   200   {object}  middleware.Response{msg=string}  "创建主机"
// @Router    /api/cmdb/createHost [post]
func (c *cmdbController) CreateHost(ctx *gin.Context) {
	params := &dto.CMDBHostCreateInput{}
	if err := params.BindingValidParams(ctx); err != nil {
		v1.Log.ErrorWithCode(globalError.ParamBindError, err)
		middleware.ResponseError(ctx, globalError.NewGlobalError(globalError.ParamBindError, err))
		return
	}
	if err := v1.CoreV1.CMDB().Host().CreateHost(ctx, params); err != nil {
		v1.Log.ErrorWithCode(globalError.CreateError, err)
		middleware.ResponseError(ctx, globalError.NewGlobalError(globalError.CreateError, err))
		return
	}
	middleware.ResponseSuccess(ctx, "")
}

// @Description  更新主机
// @ID           /api/cmdb/updateHost
// @Tags         Host
// @Summary   更新主机
// @Accept    application/json
// @Produce   application/json
// @Param     data  body      dto.CMDBHostCreateInput 	true 	"主机信息"
// @Security  ApiKeyAuth
// @Success   200   {object}  middleware.Response{msg=string}  "更新主机"
// @Router    /api/cmdb/updateHost [put]
func (c *cmdbController) UpdateHost(ctx *gin.Context) {
	userUUID, err := utils.GetUserUUID(ctx)
	if err != nil {
		v1.Log.ErrorWithCode(globalError.GetError, err)
		middleware.ResponseError(ctx, globalError.NewGlobalError(globalError.GetError, err))
		return
	}
	params := &dto.CMDBHostCreateInput{}
	if err := params.BindingValidParams(ctx); err != nil {
		v1.Log.ErrorWithCode(globalError.ParamBindError, err)
		middleware.ResponseError(ctx, globalError.NewGlobalError(globalError.ParamBindError, err))
		return
	}
	if err := v1.CoreV1.CMDB().Host().UpdateHost(ctx, userUUID, params); err != nil {
		v1.Log.ErrorWithCode(globalError.UpdateError, err)
		middleware.ResponseError(ctx, globalError.NewGlobalError(globalError.UpdateError, err))
		return
	}
	middleware.ResponseSuccess(ctx, "")
}

// @Description  主机分页
// @ID       /api/cmdb/:groupID/pageHost
// @Tags      Host
// @Summary   获取主机分页
// @Security  ApiKeyAuth
// @Produce   application/json
// @Success   200   {object}  middleware.Response{data=dto.PageCMDBHostOut,msg=string}  "获取主机分页"
// @Router    /api/cmdb/:groupID/pageHost [get]
func (c *cmdbController) PageHost(ctx *gin.Context) {
	userUUID, err := utils.GetUserUUID(ctx)
	if err != nil {
		v1.Log.ErrorWithCode(globalError.GetError, err)
		middleware.ResponseError(ctx, globalError.NewGlobalError(globalError.GetError, err))
		return
	}
	groupID, err := utils.ParseUint(ctx.Param("groupID"))
	if err != nil {
		v1.Log.ErrorWithCode(globalError.ParamBindError, err)
		middleware.ResponseError(ctx, globalError.NewGlobalError(globalError.ParamBindError, err))
		return
	}
	params := &dto.PageListCMDBHostInput{}
	if err := params.BindingValidParams(ctx); err != nil {
		v1.Log.ErrorWithCode(globalError.ParamBindError, err)
		middleware.ResponseError(ctx, globalError.NewGlobalError(globalError.ParamBindError, err))
		return
	}
	data, err := v1.CoreV1.CMDB().Host().PageHost(ctx, userUUID, groupID, params)
	if err != nil {
		v1.Log.ErrorWithCode(globalError.GetError, err)
		middleware.ResponseError(ctx, globalError.NewGlobalError(globalError.GetError, err))
		return
	}
	middleware.ResponseSuccess(ctx, data)
}

// @Description  删除主机
// @ID           /api/cmdb/:instanceID/deleteHost
// @Tags         Host
// @Summary   删除主机
// @Produce   application/json
// @Param     data  body       dto.Empty    true  "空"
// @Security  ApiKeyAuth
// @Success   200   {object}  middleware.Response{msg=string}  "删除主机"
// @Router    /api/cmdb/:instanceID/deleteHost [delete]
func (c *cmdbController) DeleteHost(ctx *gin.Context) {
	userUUID, err := utils.GetUserUUID(ctx)
	if err != nil {
		v1.Log.ErrorWithCode(globalError.GetError, err)
		middleware.ResponseError(ctx, globalError.NewGlobalError(globalError.GetError, err))
		return
	}
	instanceid := ctx.Param("instanceID")
	if err := v1.CoreV1.CMDB().Host().DeleteHost(ctx, userUUID, instanceid); err != nil {
		v1.Log.ErrorWithCode(globalError.DeleteError, err)
		middleware.ResponseError(ctx, globalError.NewGlobalError(globalError.DeleteError, err))
		return
	}
	middleware.ResponseSuccess(ctx, "")
}

// @Description  批量删除主机
// @ID           /api/cmdb/:instanceID/deleteHost
// @Tags         Host
// @Summary   批量删除主机
// @Produce   application/json
// @Param     data  body       dto.Empty    true  "空"
// @Security  ApiKeyAuth
// @Success   200   {object}  middleware.Response{msg=string}  "删除主机"
// @Router    /api/cmdb/:instanceID/deleteHost [delete]
func (c *cmdbController) DeleteHosts(ctx *gin.Context) {
	userUUID, err := utils.GetUserUUID(ctx)
	if err != nil {
		v1.Log.ErrorWithCode(globalError.GetError, err)
		middleware.ResponseError(ctx, globalError.NewGlobalError(globalError.GetError, err))
		return
	}
	params := &dto.InstancesReq{}
	if err := params.BindingValidParams(ctx); err != nil {
		v1.Log.ErrorWithCode(globalError.ParamBindError, err)
		middleware.ResponseError(ctx, globalError.NewGlobalError(globalError.ParamBindError, err))
		return
	}
	if err := v1.CoreV1.CMDB().Host().DeleteHosts(ctx, userUUID, params.Ids); err != nil {
		v1.Log.ErrorWithCode(globalError.DeleteError, err)
		middleware.ResponseError(ctx, globalError.NewGlobalError(globalError.DeleteError, err))
		return
	}
	middleware.ResponseSuccess(ctx, "")
}

// @Description  获取主机列表
// @ID       /api/cmdb/getHostsList
// @Tags      Host
// @Summary   获取主机列表
// @Security  ApiKeyAuth
// @Produce   application/json
// @Success   200   {object}  middleware.Response{data=model.CMDBHost,msg=string}  "获取主机列表"
// @Router    /api/cmdb/getHostsList [get]
func (c *cmdbController) GetHostList(ctx *gin.Context) {
	userUUID, err := utils.GetUserUUID(ctx)
	if err != nil {
		v1.Log.ErrorWithCode(globalError.GetError, err)
		middleware.ResponseError(ctx, globalError.NewGlobalError(globalError.GetError, err))
		return
	}
	data, err := v1.CoreV1.CMDB().Host().GetHostListWithGroupName(ctx, userUUID, nil)
	if err != nil {
		v1.Log.ErrorWithCode(globalError.GetError, err)
		middleware.ResponseError(ctx, globalError.NewGlobalError(globalError.GetError, err))
		return
	}
	middleware.ResponseSuccess(ctx, data)
}

func (c *cmdbController) WebShell(ctx *gin.Context) {
	instanceID := ctx.Query("instanceID")
	if utils.IsStrEmpty(instanceID) {
		v1.Log.Error("instanceID is empty")
		return
	}
	// 设置默认xterm窗口大小
	cols, _ := strconv.Atoi(ctx.DefaultQuery("cols", "188"))
	rows, _ := strconv.Atoi(ctx.DefaultQuery("rows", "42"))
	err := v1.CoreV1.CMDB().Host().WebShell(ctx, instanceID, cols, rows)
	if err != nil {
		v1.Log.ErrorWithCode(globalError.ServerError, err)
		return
	}
	v1.Log.Info("web shell connect success")
}
