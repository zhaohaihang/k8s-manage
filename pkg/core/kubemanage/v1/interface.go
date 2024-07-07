package v1

import (
	"github.com/zhaohaihang/k8s-manage/cmd/app/config"
	"github.com/zhaohaihang/k8s-manage/dao"
	"github.com/zhaohaihang/k8s-manage/pkg/logger"
)

type CoreService interface {
	WorkFlowServiceGetter
	CloudGetter
	SystemGetter
	CMDBGetter
}

func New(cfg *config.Config, factory dao.ShareDaoFactory) CoreService {
	return &KubeManage{
		Cfg:     cfg,
		Factory: factory,
	}
}

type Logger interface {
	logger.Logger
}

type KubeManage struct {
	Cfg     *config.Config
	Factory dao.ShareDaoFactory
}

func (c *KubeManage) WorkFlow() WorkFlowService {
	return NewWorkFlow(c)
}

func (c *KubeManage) Cloud() CloudInterface {
	return NewCloud(c)
}

func (c *KubeManage) System() SystemInterface {
	return NewSystem(c)
}

func (c *KubeManage) CMDB() CMDBService {
	return NewCMDBService(c.Factory)
}
