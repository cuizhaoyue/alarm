package service

import (
	"context"

	"ezone.xxxxx.com/xxxxx/xxxxx/alarm/modules/alarm/v2/dao"
	"ezone.xxxxx.com/xxxxx/xxxxx/alarm/modules/alarm/v2/dto"
)

type PromqlTplSvc interface {
	List(ctx context.Context, rstype string) (*dto.ListPromqlTplResponseData, error)
}

var _ PromqlTplSvc = &promqlTplSvc{}

type promqlTplSvc struct {
	store dao.Factory
}

func newPromqlTplSvc(svc *service) *promqlTplSvc {
	return &promqlTplSvc{
		store: svc.store,
	}
}

// List 获取PromQL模板数据
func (p *promqlTplSvc) List(ctx context.Context, rstype string) (*dto.ListPromqlTplResponseData, error) {

	promqlTpls, err := p.store.PromQLTpl().List(ctx, rstype)
	if err != nil {
		return nil, err
	}

	data := &dto.ListPromqlTplResponseData{}
	for _, item := range promqlTpls.Items {
		tplInfo := &dto.TplInfo{
			Name:            item.Name,
			DisplayName:     item.DisplayName,
			Promql:          item.Promql,
			Unit:            item.Unit,
			Algo:            item.Algo,
			AlgoDisplayName: item.AlgoDisplayName,
		}
		data.Items = append(data.Items, tplInfo)
	}

	return data, nil
}
