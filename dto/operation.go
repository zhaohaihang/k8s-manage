package dto

import (
	"github.com/gin-gonic/gin"

	"github.com/zhaohaihang/k8s-manage/dao/model"
	"github.com/zhaohaihang/k8s-manage/pkg"
)

type OperationListInput struct {
	PageInfo
	Method string `json:"method" form:"method" ` // 请求方法
	Path   string `json:"path" form:"path" `     // 请求路径
	Status int    `json:"status" form:"status" ` // 请求状态
}

type OperationListOutPut struct {
	Total         int64                       `json:"total"`
	OperationList []*model.SysOperationRecord `json:"list"`
	PageInfo
}

// BindingValidParams 绑定并校验参数
func (o *OperationListInput) BindingValidParams(ctx *gin.Context) error {
	return pkg.DefaultGetValidParams(ctx, o)
}
