package dao

import (
	"context"
	"fmt"

	"ezone.xxxxx.com/xxxxx/xxxxx/alarm/modules/alarm/v2/model"
	"ezone.xxxxx.com/xxxxx/xxxxx/communal/lib"
	"gorm.io/gorm"
)

type PromQLTplStore interface {
	List(ctx context.Context, rstype string) (*model.PromqlTemplateList, error)
}

var _ PromQLTplStore = &promqlTplStore{}

type promqlTplStore struct {
	db *gorm.DB
}

// 使用公共库中的lib.GROMDefaultPool实例，去掉参数datastore参数
func newPromTplStore(ds *datastore) *promqlTplStore {
	return &promqlTplStore{ds.db}
}

// List 获取PromQL模板列表
func (p *promqlTplStore) List(ctx context.Context, rstype string) (*model.PromqlTemplateList, error) {
	ret := &model.PromqlTemplateList{}

	err := p.db.Where("resource_sub_type = ?", rstype).Find(&ret.Items).Error

	if err != nil {
		lib.Log.TagError(lib.GetTraceContext(ctx), lib.DLTagMySqlFailed, map[string]interface{}{
			"message": fmt.Errorf("got promql templates failed, %s", err).Error(),
		})

		return nil, err
	}

	return ret, nil
}
