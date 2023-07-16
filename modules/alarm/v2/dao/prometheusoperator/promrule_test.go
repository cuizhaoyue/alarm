package prometheusoperator_test

import (
	"context"
	"log"
	"strings"
	"testing"

	"ezone.xxxxx.com/xxxxx/xxxxx/alarm/modules/alarm/v2/dao/clientset"
	"ezone.xxxxx.com/xxxxx/xxxxx/alarm/modules/alarm/v2/model"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
)

// func Test_Name(t *testing.T) {
// 	tests := []struct {
// 		name string
// 	}{
// 		// : Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			Name()
// 		})
// 	}
// }

const (
	PromruleName = "testpr"
	LabelKey     = "name"
)

func init() {
	_, err := clientset.PromOperatorClientSet()
	if err != nil {
		log.Fatalf("get client set failed, %s", err)
	}
}

func Test_PromRuleCreate(t *testing.T) {
	pr := clientset.ClientSet().MonitoringV1().PrometheusRules("monitoring")
	ctx := context.Background()

	// 构建PrometheusRule对象
	promrule := &model.PrometheusRule{
		TypeMeta: metav1.TypeMeta{
			Kind:       model.PrometheusRuleKind,
			APIVersion: "monitoring.coreos.com/v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:        PromruleName,
			Namespace:   "monitoring",
			Labels:      map[string]string{"name": PromruleName},
			Annotations: map[string]string{},
		},
		Spec: model.PrometheusRuleSpec{
			Groups: []model.RuleGroup{
				{
					Name:     "testprgroup",
					Interval: "1m",
					Rules: []model.Rule{
						{
							Alert: "testup",
							Expr: intstr.IntOrString{
								Type:   intstr.String,
								StrVal: "up == 1",
							},
							For:         "5m",
							Labels:      map[string]string{},
							Annotations: map[string]string{},
						},
					},
					// Limit:                   new(int),
				},
			},
		},
	}

	prule, err := pr.Create(ctx, promrule, metav1.CreateOptions{})
	if err != nil {
		// create PrometheusRule failed, prometheusrules.monitoring.coreos.com "testpr" already exists
		t.Logf("error type: %T", err)

		t.Logf("%t", strings.Contains(err.Error(), "already exists"))
		t.Logf("create PrometheusRule failed, %s", err)
		return
	}
	t.Logf("PrometheusRule [%s] created successfully", prule.Name)
	t.Logf("returned PrometheusRule: %#v", prule)
}

// 测试更新PrometheusRule对象
func Test_PromRuleUpdate(t *testing.T) {
	pr := clientset.ClientSet().MonitoringV1().PrometheusRules("monitoring")
	ctx := context.Background()

	// 构建PrometheusRule对象
	promeRule, _ := pr.Get(ctx, PromruleName, metav1.GetOptions{})
	promeRule.Spec = model.PrometheusRuleSpec{
		Groups: []model.RuleGroup{
			{
				Name:     "testprgroup",
				Interval: "1m",
				Rules: []model.Rule{
					{
						Alert: "testup",
						Expr: intstr.IntOrString{
							Type:   intstr.String,
							StrVal: "up == 0",
						},
						For:         "5m",
						Labels:      map[string]string{},
						Annotations: map[string]string{},
					},
				},
				// Limit:                   new(int),
			},
			{
				Name:     "testprgroup11", // rule的名称不能相同
				Interval: "1m",
				Rules: []model.Rule{
					{
						Alert: "testup2",
						Expr: intstr.IntOrString{
							Type:   intstr.String,
							StrVal: "up == 0",
						},
						For:         "5m",
						Labels:      map[string]string{},
						Annotations: map[string]string{},
					},
				},
			},
		},
	}

	_, err := pr.Update(ctx, promeRule, metav1.UpdateOptions{})
	if err != nil {
		t.Errorf("update failed, %s", err)
	}
}

// 测试删除PrometheusRule对象
func Test_PromRuleDelete(t *testing.T) {
	pr := clientset.ClientSet().MonitoringV1().PrometheusRules("monitoring")
	ctx := context.Background()

	// err := pr.Delete(ctx, PromruleName, metav1.DeleteOptions{})

	// opts := metav1.ListOptions{
	// 	LabelSelector: labels.Set{"name": PromruleName}.AsSelector().String(),
	// }
	// t.Log(labels.Set{"name": PromruleName}.AsSelector().String())

	// 使用 'key in (v1, v2)' 的格式
	labelValues := []string{PromruleName}
	labelSelector := LabelKey + " in " + "(" + strings.Join(labelValues, ",") + ")"
	opts := metav1.ListOptions{
		LabelSelector: labelSelector,
	}

	t.Log(labelSelector)

	err := pr.DeleteCollection(ctx, metav1.DeleteOptions{}, opts)

	if err != nil {
		t.Logf("delete PrometheusRule failed, %s", err)
	}
}

// 测试获取PrometheusRule对象
func Test_PromRuleGet(t *testing.T) {
	clientset := clientset.ClientSet()

	pro, err := clientset.MonitoringV1().PrometheusRules("monitoring").Get(context.TODO(), PromruleName, metav1.GetOptions{})
	if err != nil {
		t.Logf("error type: %T", err)
		t.Logf("get prometheus rule failed, %s", err)
		t.Logf("%v", pro)
	}

	t.Logf("got PrometheusRule: %s", pro.Name)
}
