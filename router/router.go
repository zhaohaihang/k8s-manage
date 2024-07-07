package router

import (
	"github.com/zhaohaihang/k8s-manage/cmd/app/options"
	"github.com/zhaohaihang/k8s-manage/controller/api"
	"github.com/zhaohaihang/k8s-manage/controller/authority"
	"github.com/zhaohaihang/k8s-manage/controller/cmdb"
	"github.com/zhaohaihang/k8s-manage/controller/kubeController"
	"github.com/zhaohaihang/k8s-manage/controller/menu"
	"github.com/zhaohaihang/k8s-manage/controller/operation"
	"github.com/zhaohaihang/k8s-manage/controller/other"
	"github.com/zhaohaihang/k8s-manage/controller/system"
	"github.com/zhaohaihang/k8s-manage/controller/user"
	"github.com/zhaohaihang/k8s-manage/middleware"
)

func InstallRouters(opt *options.Options) {
	apiGroup := opt.GinEngine.Group("/api")
	middleware.InstallMiddlewares(apiGroup)
	{
		api.NewApiRouter(apiGroup)
		operation.NewOperationRouter(apiGroup)
		user.NewUserRouter(apiGroup)
		other.NewSwaggarRoute(apiGroup)
		kubeController.NewKubeRouter(apiGroup)
		menu.NewMenuRouter(apiGroup)
		authority.NewCasbinRouter(apiGroup)
		system.NewSystemController(apiGroup)
	}
	{
		// cmdb相关
		cmdb.NewCMDBRouter(apiGroup)
	}
}
