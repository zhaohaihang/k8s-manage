package kubeController

import (
	"github.com/gin-gonic/gin"
	"github.com/zhaohaihang/k8s-manage/dto/kubeDto"
	"github.com/zhaohaihang/k8s-manage/pkg/core/kubemanage/v1"

	"github.com/zhaohaihang/k8s-manage/middleware"
	"github.com/zhaohaihang/k8s-manage/pkg/core/kubemanage/v1/kube"
	"github.com/zhaohaihang/k8s-manage/pkg/globalError"
)

var StatefulSet statefulSet

type statefulSet struct{}

// DeleteStatefulSet 删除statefulSet
// ListPage godoc
// @Summary      删除statefulSet
// @Description  删除statefulSet
// @Tags         statefulSet
// @ID           /api/k8s/statefulset/del
// @Accept       json
// @Produce      json
// @Param        name       query  string  true  "statefulSet名称"
// @Param        namespace    query  string  true  "命名空间"
// @Success       200  {object}  middleware.Response "{"code": 200, msg="","data": "删除成功}"
// @Router       /api/k8s/statefulset/del [delete]
func (s *statefulSet) DeleteStatefulSet(ctx *gin.Context) {
	params := &kubeDto.StatefulSetNameNS{}
	if err := params.BindingValidParams(ctx); err != nil {
		v1.Log.ErrorWithCode(globalError.ParamBindError, err)
		middleware.ResponseError(ctx, globalError.NewGlobalError(globalError.ParamBindError, err))
		return
	}
	if err := kube.StatefulSet.DeleteStatefulSet(params.Name, params.NameSpace); err != nil {
		v1.Log.ErrorWithCode(globalError.DeleteError, err)
		middleware.ResponseError(ctx, globalError.NewGlobalError(globalError.DeleteError, err))
		return
	}
	middleware.ResponseSuccess(ctx, "删除成功")
}

// UpdateStatefulSet 更新statefulSet
// ListPage godoc
// @Summary      更新statefulSet
// @Description  更新statefulSet
// @Tags         statefulSet
// @ID           /api/k8s/statefulset/update
// @Accept       json
// @Produce      json
// @Param        name       query  string  true  "无状态控制器名称"
// @Param        namespace  query  string  true  "命名空间"
// @Param        content    query  string  true  "更新内容"
// @Success       200  {object}  middleware.Response"{"code": 200, msg="","data": "更新成功}"
// @Router       /api/k8s/statefulset/update [put]
func (s *statefulSet) UpdateStatefulSet(ctx *gin.Context) {
	params := &kubeDto.StatefulSetUpdateInput{}
	if err := params.BindingValidParams(ctx); err != nil {
		v1.Log.ErrorWithCode(globalError.ParamBindError, err)
		middleware.ResponseError(ctx, globalError.NewGlobalError(globalError.ParamBindError, err))
		return
	}
	if err := kube.StatefulSet.UpdateStatefulSet(params.Content, params.NameSpace); err != nil {
		v1.Log.ErrorWithCode(globalError.UpdateError, err)
		middleware.ResponseError(ctx, globalError.NewGlobalError(globalError.UpdateError, err))
		return
	}
	middleware.ResponseSuccess(ctx, "更新成功")
}

// GetStatefulSetList 查看statefulSet列表
// ListPage godoc
// @Summary      查看statefulSet列表
// @Description  查看statefulSet列表
// @Tags         statefulSet
// @ID           /api/k8s/statefulset/list
// @Accept       json
// @Produce      json
// @Param        filter_name  query  string  false  "过滤"
// @Param        namespace  query  string  false  "命名空间"
// @Param        page         query  int     false  "页码"
// @Param        limit        query  int     false  "分页限制"
// @Success       200  {object}  middleware.Response"{"code": 200, msg="","data": }"
// @Router       /api/k8s/statefulset/list [get]
func (s *statefulSet) GetStatefulSetList(ctx *gin.Context) {
	params := &kubeDto.StatefulSetListInput{}
	if err := params.BindingValidParams(ctx); err != nil {
		v1.Log.ErrorWithCode(globalError.ParamBindError, err)
		middleware.ResponseError(ctx, globalError.NewGlobalError(globalError.ParamBindError, err))
		return
	}
	data, err := kube.StatefulSet.GetStatefulSets(params.FilterName, params.NameSpace, params.Limit, params.Page)
	if err != nil {
		v1.Log.ErrorWithCode(globalError.GetError, err)
		middleware.ResponseError(ctx, globalError.NewGlobalError(globalError.GetError, err))
		return
	}
	middleware.ResponseSuccess(ctx, data)
}

// GetStatefulSetDetail 获取statefulSet详情
// ListPage godoc
// @Summary      获取statefulSet详情
// @Description  获取statefulSet详情
// @Tags         statefulSet
// @ID           /api/k8s/statefulSet/detail
// @Accept       json
// @Produce      json
// @Param        name       query  string  true  "statefulSet名称"
// @Param        namespace  query  string  true  "命名空间"
// @Success      200        {object}  middleware.Response"{"code": 200, msg="","data":v1.Deployment }"
// @Router       /api/k8s/statefulset/detail [get]
func (s *statefulSet) GetStatefulSetDetail(ctx *gin.Context) {
	params := &kubeDto.StatefulSetNameNS{}
	if err := params.BindingValidParams(ctx); err != nil {
		v1.Log.ErrorWithCode(globalError.ParamBindError, err)
		middleware.ResponseError(ctx, globalError.NewGlobalError(globalError.ParamBindError, err))
		return
	}
	data, err := kube.StatefulSet.GetStatefulSetDetail(params.Name, params.NameSpace)
	if err != nil {
		v1.Log.ErrorWithCode(globalError.GetError, err)
		middleware.ResponseError(ctx, globalError.NewGlobalError(globalError.GetError, err))
		return
	}
	middleware.ResponseSuccess(ctx, data)
}
