package cmdb

import (
	"context"

	"gorm.io/gorm"

	"github.com/zhaohaihang/k8s-manage/dao"
	"github.com/zhaohaihang/k8s-manage/dao/model"
)

const (
	checkSuccess = 1
	checkFailed  = 2
)

var checkFailedHandler = func(factory dao.ShareDaoFactory, host model.CMDBHost) error {
	host.Status = checkFailed
	if err := factory.CMDB().Host().Updates(context.TODO(), func(db *gorm.DB) *gorm.DB {
		return db.Where("instanceID = ? ", host.InstanceID)
	}, &host); err != nil {
		return err
	}
	return nil
}

var checkSuccessHandler = func(factory dao.ShareDaoFactory, host model.CMDBHost) error {
	host.Status = checkSuccess
	if err := factory.CMDB().Host().Updates(context.TODO(), func(db *gorm.DB) *gorm.DB {
		return db.Where("instanceID = ? ", host.InstanceID)
	}, &host); err != nil {
		return err
	}
	return nil
}
