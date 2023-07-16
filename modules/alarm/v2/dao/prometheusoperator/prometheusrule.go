package prometheusoperator

import (
	"context"
	"fmt"

	"ezone.xxxxx.com/xxxxx/xxxxx/alarm/modules/alarm/v2/model"
	"ezone.xxxxx.com/xxxxx/xxxxx/communal/lib"
	clientversioned "github.com/prometheus-operator/prometheus-operator/pkg/client/versioned"
	typedmonitoringv1 "github.com/prometheus-operator/prometheus-operator/pkg/client/versioned/typed/monitoring/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	Namespace = "monitoring"
)

// PromRuleOperator 操作PrometheusRule的接口
type PromRuleOperator interface {
	Create(ctx context.Context, pr *model.PrometheusRule) (*model.PrometheusRule, error)
	Update(ctx context.Context, pr *model.PrometheusRule) (*model.PrometheusRule, error)
	Get(ctx context.Context, name string) (*model.PrometheusRule, error)
	Delete(ctx context.Context, name string) error
	DeleteCollection(ctx context.Context, listOpts metav1.ListOptions) error
}

var _ PromRuleOperator = &prometheusRuleOperator{}

type prometheusRuleOperator struct {
	client typedmonitoringv1.PrometheusRuleInterface
}

func NewPromRuleOperator(cs *clientversioned.Clientset) *prometheusRuleOperator {
	return &prometheusRuleOperator{
		client: cs.MonitoringV1().PrometheusRules(Namespace),
	}
}

// Create 创建k8s集群中的PrometheusRule对象
func (p *prometheusRuleOperator) Create(ctx context.Context, pr *model.PrometheusRule) (*model.PrometheusRule, error) {
	pr, err := p.client.Create(ctx, pr, metav1.CreateOptions{})
	if err != nil {
		lib.Log.TagError(lib.GetTraceContext(ctx), lib.DLTagMySqlFailed, map[string]interface{}{
			"message": fmt.Errorf("create PrometheusRule [%s] failed, %s", pr.Name, err).Error(),
		})

		return nil, err
	}

	return pr, nil
}

// Update 更新PrometheusRule对象
func (p *prometheusRuleOperator) Update(ctx context.Context, pr *model.PrometheusRule) (*model.PrometheusRule, error) {
	pr, err := p.client.Update(ctx, pr, metav1.UpdateOptions{})
	if err != nil {
		lib.Log.TagError(lib.GetTraceContext(ctx), lib.DLTagMySqlFailed, map[string]interface{}{
			"message": fmt.Errorf("update PrometheusRule failed, %s", err).Error(),
		})

		return nil, err
	}

	return pr, nil
}

// Get 获取PrometheusRule对象
func (p *prometheusRuleOperator) Get(ctx context.Context, name string) (*model.PrometheusRule, error) {
	pr, err := p.client.Get(ctx, name, metav1.GetOptions{})
	if err != nil {
		lib.Log.TagError(lib.GetTraceContext(ctx), lib.DLTagMySqlFailed, map[string]interface{}{
			"message": fmt.Errorf("get PrometheusRule failed, %s", err).Error(),
		})

		return nil, err
	}

	return pr, nil
}

// Delete 删除指定的PrometheusRule
func (p *prometheusRuleOperator) Delete(ctx context.Context, name string) error {
	err := p.client.Delete(ctx, name, metav1.DeleteOptions{})
	if err != nil {
		lib.Log.TagError(lib.GetTraceContext(ctx), lib.DLTagMySqlFailed, map[string]interface{}{
			"message": fmt.Errorf("delete PrometheusRule failed, %s", err).Error(),
		})

		return err
	}

	return nil
}

// DeleteCollection 批量删除PrometheusRule
func (p *prometheusRuleOperator) DeleteCollection(ctx context.Context, listOpts metav1.ListOptions) error {
	// ListOptions指定标签批量删除
	err := p.client.DeleteCollection(ctx, metav1.DeleteOptions{}, listOpts)
	if err != nil {
		lib.Log.TagError(lib.GetTraceContext(ctx), lib.DLTagMySqlFailed, map[string]interface{}{
			"message": fmt.Errorf("delete PrometheusRule collection failed, %s", err).Error(),
		})
		return err
	}

	return nil
}
