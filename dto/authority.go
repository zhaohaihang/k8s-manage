package dto

import (
	"github.com/gin-gonic/gin"

	"github.com/zhaohaihang/k8s-manage/dao/model"
	"github.com/zhaohaihang/k8s-manage/pkg"
)

type AuthorityList struct {
	PageInfo
	Total             int64                `json:"total"`
	AuthorityListItem []model.SysAuthority `json:"list"`
}

type AuthorityCreateUpdateInput struct {
	AuthorityId   uint   `json:"authorityId"`
	AuthorityName string `json:"authorityName"`
}

// BindingValidParams 绑定并校验参数
func (a *AuthorityCreateUpdateInput) BindingValidParams(ctx *gin.Context) error {
	return pkg.DefaultGetValidParams(ctx, a)
}
