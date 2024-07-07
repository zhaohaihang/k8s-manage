package sys

import (
	"context"
	"github.com/zhaohaihang/k8s-manage/dao"
	"github.com/zhaohaihang/k8s-manage/dao/model"
)

type APIServiceGetter interface {
	Api() APIService
}

type APIService interface {
	GetApiList(ctx context.Context) ([]model.SysApi, error)
}

var _ APIService = &apiService{}

func NewApiService(factory dao.ShareDaoFactory) APIService {
	return &apiService{factory: factory}
}

type apiService struct {
	factory dao.ShareDaoFactory
}

func (a *apiService) GetApiList(ctx context.Context) ([]model.SysApi, error) {
	// 不做任何限制查询全量数据
	var search model.SysApi
	return a.factory.Api().FindList(ctx, search)
}
