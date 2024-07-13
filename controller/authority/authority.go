package authority

import (
	"github.com/gin-gonic/gin"

	"github.com/zhaohaihang/k8s-manage/dto"
	"github.com/zhaohaihang/k8s-manage/middleware"
	v1 "github.com/zhaohaihang/k8s-manage/pkg/core/kubemanage/v1"
	"github.com/zhaohaihang/k8s-manage/pkg/globalError"
	"github.com/zhaohaihang/k8s-manage/pkg/utils"
)

// @Description  创建角色
// @ID        /api/authority/createAuthority
// @Tags      Authority
// @Summary   创建角色
// @Accept    application/json
// @Produce   application/json
// @Param     data  body      dto.AuthorityCreateUpdateInput 	true 	"角色信息"
// @Security  ApiKeyAuth
// @Success   200   {object}  middleware.Response{msg=string}  "创建角色"
// @Router    /api/authority/createAuthority [post]
func (a *authorityController) CreateAuthority(ctx *gin.Context) {
	params := &dto.AuthorityCreateUpdateInput{}
	if err := params.BindingValidParams(ctx); err != nil {
		v1.Log.ErrorWithCode(globalError.ParamBindError, err)
		middleware.ResponseError(ctx, globalError.NewGlobalError(globalError.ParamBindError, err))
		return
	}
	if err := v1.CoreV1.System().Authority().CreateAuthority(ctx, params.AuthorityId, params.AuthorityName); err != nil {
		v1.Log.ErrorWithCode(globalError.CreateError, err)
		middleware.ResponseError(ctx, globalError.NewGlobalError(globalError.CreateError, err))
		return
	}
	middleware.ResponseSuccess(ctx, "创建成功")
}

// @Description  获取角色列表
// @ID        /api/authority/getAuthorityList
// @Tags      Authority
// @Summary   获取角色列表
// @Accept    application/json
// @Produce   application/json
// @Param     data  body      dto.PageInfo  true  "分页信息"
// @Security  ApiKeyAuth
// @Success   200   {object}  middleware.Response{data=dto.AuthorityList,msg=string}  "获取角色列表"
// @Router    /api/authority/getAuthorityList [get]
func (a *authorityController) GetAuthorityList(ctx *gin.Context) {
	params := &dto.PageInfo{}
	if err := params.BindingValidParams(ctx); err != nil {
		v1.Log.ErrorWithCode(globalError.ParamBindError, err)
		middleware.ResponseError(ctx, globalError.NewGlobalError(globalError.ParamBindError, err))
		return
	}
	data, err := v1.CoreV1.System().Authority().PageAuthority(ctx, *params)
	if err != nil {
		v1.Log.ErrorWithCode(globalError.GetError, err)
		middleware.ResponseError(ctx, globalError.NewGlobalError(globalError.GetError, err))
		return
	}
	middleware.ResponseSuccess(ctx, data)
}

// @Description  更新角色
// @ID	      /api/authority/updateAuthority
// @Tags      Authority
// @Summary   更新角色
// @Accept    application/json
// @Produce   application/json
// @Param     data  body      dto.AuthorityCreateUpdateInput    true  "角色信息"
// @Security  ApiKeyAuth
// @Success   200   {object}  middleware.Response{msg=string}  "更新角色"
// @Router    /api/authority/updateAuthority [put]
func (a *authorityController) UpdateAuthority(ctx *gin.Context) {
	params := &dto.AuthorityCreateUpdateInput{}
	if err := params.BindingValidParams(ctx); err != nil {
		v1.Log.ErrorWithCode(globalError.ParamBindError, err)
		middleware.ResponseError(ctx, globalError.NewGlobalError(globalError.ParamBindError, err))
		return
	}
	if err := v1.CoreV1.System().Authority().UpdateAuthority(ctx, params.AuthorityId, params.AuthorityName); err != nil {
		v1.Log.ErrorWithCode(globalError.UpdateError, err)
		middleware.ResponseError(ctx, globalError.NewGlobalError(globalError.UpdateError, err))
		return
	}
	middleware.ResponseSuccess(ctx, "修改成功")
}

// @Description  删除角色
// @ID		/api/authority/:authID/delAuthority  
// @Tags      Authority
// @Summary   删除角色
// @Produce   application/json
// @Param     data  body      dto.Empty    true  "空"
// @Security  ApiKeyAuth
// @Success   200   {object}  middleware.Response{msg=string}  "删除角色"
// @Router    /api/authority/:authID/delAuthority [delete]
func (a *authorityController) DeleteAuthority(ctx *gin.Context) {
	authID, err := utils.ParseUint(ctx.Param("authID"))
	if err != nil {
		v1.Log.ErrorWithCode(globalError.ParamBindError, err)
		middleware.ResponseError(ctx, globalError.NewGlobalError(globalError.ParamBindError, err))
		return
	}
	if err := v1.CoreV1.System().Authority().DeleteAuthority(ctx, authID); err != nil {
		v1.Log.ErrorWithCode(globalError.DeleteError, err)
		middleware.ResponseError(ctx, globalError.NewGlobalError(globalError.DeleteError, err))
		return
	}
	middleware.ResponseSuccess(ctx, "删除成功")
}
