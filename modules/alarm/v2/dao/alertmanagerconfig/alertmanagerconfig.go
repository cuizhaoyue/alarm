package alertmanagerconfig

import (
	"context"
	"fmt"

	"ezone.xxxxx.com/xxxxx/xxxxx/alarm/modules/alarm/v2/model"
	"ezone.xxxxx.com/xxxxx/xxxxx/communal/lib"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	Namespace       = "monitoring"
	ResourceGroup   = "monitoring.coreos.com"
	ResourceVersion = "v1alpha1"
	ResourceKind    = "AlertmanagerConfig"
)

type AlertmanagerConfiger interface {
	Create(ctx context.Context, am *model.AlertmanagerConfig) error
	Delete(ctx context.Context, name string) error
	DeleteCollection(ctx context.Context, lstOpts metav1.ListOptions) error
}

var _ AlertmanagerConfiger = &amConfig{}

type amConfig struct {
	client model.AlertmanagerConfigInterfac
}

func NewAlertmanagerConfiger(cs *model.Clientset) *amConfig {
	return &amConfig{
		client: cs.MonitoringV1alpha1().AlertmanagerConfigs(Namespace),
	}
}

// Create 创建alertmanagerconfig实例
func (c *amConfig) Create(ctx context.Context, am *model.AlertmanagerConfig) error {
	_, err := c.client.Create(ctx, am, metav1.CreateOptions{})

	if err != nil {
		lib.Log.TagError(lib.GetTraceContext(ctx), lib.DLTagMySqlFailed, map[string]interface{}{
			"message": fmt.Errorf("create AlertmanagerConfig [%s] failed, %s", am.Name, err).Error(),
		})

		return err
	}

	return nil
}

// Delete 删除alertmanagerconfig实例
func (c *amConfig) Delete(ctx context.Context, name string) error {
	if err := c.client.Delete(ctx, name, metav1.DeleteOptions{}); err != nil {
		lib.Log.TagError(lib.GetTraceContext(ctx), lib.DLTagMySqlFailed, map[string]interface{}{
			"message": fmt.Errorf("delete AlertmanagerConfig [%s] failed, %s", name, err).Error(),
		})

		return err
	}

	return nil
}

// DeleteCollection 批量删除AlertmanagerConfig实例
func (c *amConfig) DeleteCollection(ctx context.Context, lstOpts metav1.ListOptions) error {
	if err := c.client.DeleteCollection(ctx, metav1.DeleteOptions{}, lstOpts); err != nil {
		lib.Log.TagError(lib.GetTraceContext(ctx), lib.DLTagMySqlFailed, map[string]interface{}{
			"message": fmt.Errorf("batch delete AlertmanagerConfig failed, %s", err).Error(),
		})

		return err
	}

	return nil
}
