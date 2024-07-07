package checker

import "github.com/zhaohaihang/k8s-manage/runtime/queue"

type Checker interface {
	Run()
	Check(event *queue.Event) error
	HandlerErr(err error)
}
