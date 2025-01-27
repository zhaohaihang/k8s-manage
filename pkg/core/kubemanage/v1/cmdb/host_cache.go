package cmdb

import (
	"context"

	"github.com/zhaohaihang/k8s-manage/dao/model"
	"github.com/zhaohaihang/k8s-manage/runtime/queue"
)

// StartHostCheck 从数据库中不断查询放到queue中
func (h *hostService) StartHostCheck() error {
	hosts, err := h.getHostList(context.TODO(), model.CMDBHost{})

	if len(hosts) < 0 {
		return err
	}

	for _, host := range hosts {
		if h.queue.IsClosed() {
			return nil
		}
		h.queue.Push(&queue.Event{Type: "AddHOST", Data: host})
	}
	return nil
}
