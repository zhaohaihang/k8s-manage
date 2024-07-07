package v1

import (
	"github.com/zhaohaihang/k8s-manage/dao"
	"github.com/zhaohaihang/k8s-manage/pkg/core/kubemanage/v1/kube"
)

type CloudGetter interface {
	Cloud() CloudInterface
}

type CloudInterface interface {
	kube.PodsGetter
}

type cloud struct {
	app     *KubeManage
	factory dao.ShareDaoFactory
}

func (c *cloud) Pods(cloud string) kube.PodInterface {
	// TODO 临时添加，需要重构
	return kube.NewPods(nil, "", c.factory)
}

func NewCloud(c *KubeManage) CloudInterface {
	return &cloud{
		app:     c,
		factory: c.Factory,
	}
}
