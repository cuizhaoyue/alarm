package service

import (
	"context"
	"strconv"
	"strings"
	"time"

	"ezone.xxxxx.com/xxxxx/xxxxx/alarm/common/consts"
	"ezone.xxxxx.com/xxxxx/xxxxx/alarm/modules/alarm/v2/dao"
	"ezone.xxxxx.com/xxxxx/xxxxx/alarm/modules/alarm/v2/dto"
	"ezone.xxxxx.com/xxxxx/xxxxx/alarm/modules/alarm/v2/model"
	"github.com/AlekSi/pointer"
	"github.com/flosch/pongo2/v6"
	"github.com/pkg/errors"
	"github.com/rs/xid"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"k8s.io/apimachinery/pkg/util/intstr"
)

const (
	PromRuleNameKey  = "name"
	ProductionKey    = "production"
	EvaluateInterval = "1m"
	GroupInterval    = "3m"

	Namespace       = "monitoring"
	ResourceGroup   = "monitoring.coreos.com"
	ResourceVersion = "v1alpha1"
	ResourceKind    = "AlertmanagerConfig"
)

// Create 创建告警策略
func (s *alarmService) Create(ctx context.Context, req *dto.CreatePolicyRequest) error {
	// 根据请求构建AlarmPolicy实例
	policy := buildPolicyInstanceFromCreateReq(req)

	lables, annotations := buildCommonLablesAndAnnotations(policy)

	// 根据告警策略构造告警规则列表
	rules := buildRulesFromPolicy(policy, lables, annotations)

	// 为每个告警规则构建AlertmanagerConfig对象
	var amcfgs []*model.AlertmanagerConfig
	if checkBuildAm(lables, annotations) {
		amcfgs = buildAlertmanagerConfig(rules)
	}

	// 创建AlertmanagerConfig
	for _, amcfg := range amcfgs {
		if err := s.store.AlertmanagerConfig().Create(ctx, amcfg); err != nil {
			s.deleteAlertmanagerConfigs(ctx, []string{policy.InstanceId})
			return errors.Wrap(err, "create AlertmanagerConfig failed")
		}
	}

	// 构建PrometheusRule对象
	promRule := buildPrometheusRule(policy.InstanceId, policy.InstanceId, rules, policy.Limit)

	// 调用operator接口，创建PrometheusRule，获取返回的PrometheusRule实例
	_, err := s.store.PromRuleOperator().Create(ctx, promRule)
	if err != nil {
		return errors.Wrap(err, "create PrometheusRule failed")
	}

	resources := buildResourceOnPolicy(policy)

	// 在事务中操作策略表和资源表
	if err := s.store.Alarm().Tx(ctx, func(ctx context.Context, store dao.Factory) error {
		// 保存告警策略到数据库中
		if err := store.Alarm().Create(ctx, policy); err != nil {
			return err
		}

		// 如果是表单格式告警重要则需要保存选择的资源列表到数据库中
		if req.Type == model.FormAlarmPolicyType {
			if err := store.ResourceOnPolicy().Create(ctx, resources); err != nil {
				return err
			}
		}

		return nil
	}); err != nil {
		// 如果入库失败要删除k8s中的PrometheusRule和AlertmanagerConfig
		if deleteErr := s.store.PromRuleOperator().Delete(ctx, promRule.Name); deleteErr != nil {
			return errors.Wrapf(deleteErr, "delete PrometheusRule [%s] failed", promRule.Name)

		}
		if len(amcfgs) > 0 {
			if deleteErr := s.deleteAlertmanagerConfigs(ctx, []string{policy.InstanceId}); deleteErr != nil {
				return errors.Wrapf(deleteErr, "delete AlertmanagerConfig for policy [%s] failed", policy.InstanceId)
			}
		}

		return errors.Wrapf(err, "create alarm policy [%s] failed", policy.Name)
	}

	return err
}

// 通过告警策略构建要保存的资源实例列表
func buildResourceOnPolicy(policy *model.AlarmPolicy) []*model.ResourceOnPolicy {
	var resources []*model.ResourceOnPolicy
	for _, ins := range policy.FormPolicy.Resources.Instances {
		resources = append(resources, &model.ResourceOnPolicy{
			PolicyId: policy.InstanceId,
			Name:     ins.Value,
			Region:   ins.Region,
			Az:       ins.Az,
		})
	}

	return resources
}

// 根据创建请求构建出告警策略实例
func buildPolicyInstanceFromCreateReq(req *dto.CreatePolicyRequest) *model.AlarmPolicy {
	policy := &model.AlarmPolicy{
		InstanceId:  xid.New().String(),
		Name:        req.Name,
		Creator:     req.Creator,
		Updater:     req.Updater,
		Enabled:     true,
		Production:  req.Production,
		Comment:     req.Comment,
		Labels:      req.Labels,
		Limit:       req.Limit,
		Type:        req.Type,
		NotifySetup: req.NotifySetup,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	// 判断策略类型
	switch req.Type {
	case model.FormAlarmPolicyType:
		// 表单类型告警策略
		policy.FormPolicy = req.FormPolicy
		policy.ResourceType = req.FormPolicy.ResourceType
		policy.ResourceSubType = req.FormPolicy.ResourceSubType

		// 添加告警规则id
		for _, r := range policy.FormPolicy.Rules {
			if r.RuleId == "" {
				r.RuleId = xid.New().String()
			}
		}
	case model.PromAlarmPolicyType:
		// PromQL类型告警策略
		policy.PromqlPolicy = req.PromqlPolicy
		// 添加告警规则id
		if policy.PromqlPolicy.RuleId == "" {
			policy.PromqlPolicy.RuleId = xid.New().String()
		}
	}

	// 判断是否传入了联系人
	if req.Receivers != nil {
		policy.Receivers = req.Receivers
	}

	// 转换数据格式
	policy.FE2DB()

	return policy
}

// 通过告警策略构建告警规则实例列表和对应的AlertmanagerConfig实例
// func buildRulesFromPolicy(policy *model.AlarmPolicy) ([]model.Rule, []*model.AlertmanagerConfig) {
func buildRulesFromPolicy(policy *model.AlarmPolicy, labels, annotations map[string]string) []model.Rule {
	var rules []model.Rule

	// labels := make(map[string]string)
	// annotations := make(map[string]string)

	// 构建通用labels和annotations
	// buildCommonLablesAndAnnotations(policy, labels, annotations)

	// 根据策略类型构建不同的Rule
	switch policy.Type {
	case model.FormAlarmPolicyType:
		// 表单格式的告警策略

		for _, r := range policy.FormPolicy.Rules {
			// 构建告警规则的labels和annotations
			rLabels := make(map[string]string)
			rAnnotations := make(map[string]string)
			copyMap(labels, rLabels)
			copyMap(annotations, rAnnotations)

			buildLablesAndAnnotationsForFormPolicy(r, rLabels, rAnnotations)

			// 渲染PromQL语句
			// promql := buildPromQLForFormPolicy(policy.FormPolicy, r)
			promql := `up == 1` // 仅测试用

			rule := model.Rule{
				Alert: r.Name,
				Expr: intstr.IntOrString{
					Type:   intstr.String,
					StrVal: promql,
				},
				For:         model.Duration(r.For),
				Labels:      rLabels,
				Annotations: rAnnotations,
			}

			rules = append(rules, rule)
		}
	case model.PromAlarmPolicyType:
		// promql格式的告警策略
		rLabels := make(map[string]string)
		rAnnotations := make(map[string]string)
		copyMap(labels, rLabels)
		copyMap(annotations, rAnnotations)

		buildLablesAndAnnotationsForPromqlPolicy(policy.PromqlPolicy, rLabels, rAnnotations)

		rule := model.Rule{
			Alert: policy.PromqlPolicy.Name,
			Expr: intstr.IntOrString{
				Type:   intstr.String,
				StrVal: policy.PromqlPolicy.Promql,
			},
			For:         model.Duration(policy.PromqlPolicy.For),
			Labels:      rLabels,
			Annotations: rAnnotations,
		}

		rules = append(rules, rule)
	}

	return rules
}

func copyMap(src, dst map[string]string) {
	for k, v := range src {
		dst[k] = v
	}
}

// 检测是否需要为rules创建AlertmanagerConfig实例
// TODO: 添加内部告警使用的接收标志
func checkBuildAm(labels, annotations map[string]string) bool {
	// 如果接收人设置和通知设置有一个不为空则需要添加WebhookConfig来发送告警消息
	if annotations[consts.ToEmails] != "" ||
		annotations[consts.ToSms] != "" ||
		annotations[consts.EnableDingtalk] == "true" ||
		annotations[consts.EnableFeishu] == "true" ||
		annotations[consts.EnableWechat] == "true" {
		return true
	}

	return false
}

// 构建AlertmanagerConfig对象实例
func buildAlertmanagerConfig(rules []model.Rule) []*model.AlertmanagerConfig {
	var amcfgs []*model.AlertmanagerConfig

	for _, rule := range rules {
		labels := rule.Labels
		annotations := rule.Annotations
		amLabels := make(map[string]string)

		// AlertmanagerConfig对象的label
		amLabels[consts.RuleId] = labels[consts.RuleId]
		amLabels[consts.PolicyId] = labels[consts.PolicyId]

		// 构建WebhookConfig实例
		// sendResolved, _ := strconv.ParseBool(annotations[consts.SendResolved])
		webhoookConfigs := []model.WebhookConfig{
			{
				// TODO: 确定好url后要添加上
				URL: pointer.ToString("http://alarmv2.xxxxx/alarm/v2/webhook"),
				// 告警恢复必须为true，否则无法收到resolved状态的告警
				// 发送告警消息前要增加判断annotations[consts.SendResolved]的逻辑
				SendResolved: pointer.ToBool(true),
			},
		}

		amcfg := &model.AlertmanagerConfig{
			TypeMeta: metav1.TypeMeta{
				Kind:       ResourceKind,
				APIVersion: ResourceGroup + "/" + ResourceVersion,
			},
			ObjectMeta: metav1.ObjectMeta{
				Name:      labels[consts.RuleId],
				Namespace: Namespace,
				Labels:    amLabels,
			},
			Spec: model.AlertmanagerConfigSpec{
				Receivers: []model.Receiver{
					{
						Name:           labels[consts.RuleId],
						WebhookConfigs: webhoookConfigs,
					},
				},
				Route: &model.Route{
					GroupBy: []string{"alertname"},
					Matchers: []model.Matcher{
						{
							Name:  consts.RuleId,
							Value: labels[consts.RuleId],
						},
					},
					Receiver:       labels[consts.RuleId],
					RepeatInterval: annotations[consts.RepeatInterval],
					GroupWait:      annotations[consts.GroupWait],
					GroupInterval:  GroupInterval,
					Continue:       false,
				},
			},
		}

		amcfgs = append(amcfgs, amcfg)
	}

	return amcfgs
}

// 构造PrometheusRule实例
// name: PrometheusRule实例的名称
// groupName: 告警规则组名称
// rules: 规则组中的所有规则
func buildPrometheusRule(name, groupName string, rules []model.Rule, limit int) *model.PrometheusRule {
	// 设置PrometheusRule的labels和annotations
	// TODO: 预留出来，暂时不添加值
	labels := map[string]string{
		PromRuleNameKey:         name,
		consts.DefaultLabelname: consts.DefaultLabelvalue,
	}
	annotations := make(map[string]string)

	return &model.PrometheusRule{
		TypeMeta: metav1.TypeMeta{
			Kind:       model.PrometheusRuleKind,
			APIVersion: "monitoring.coreos.com/v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:        name,
			Namespace:   "monitoring",
			Labels:      labels,
			Annotations: annotations,
		},
		Spec: model.PrometheusRuleSpec{
			Groups: []model.RuleGroup{
				{
					Name:     groupName,
					Interval: EvaluateInterval,
					Rules:    rules,
					Limit:    pointer.ToInt(limit), // 设置告警次数，请求参数应该放在告警规则之外，不能对单个告警规则设置告警次数
				},
			},
		},
	}
}

// TODO: 构建告警规则中的PromQL语句
func buildPromQLForFormPolicy(formPolicy *model.FormPolicy, rule *model.AlertRule) string {
	// 构造模板
	tpl, err := pongo2.FromString(rule.PromqlTpl.Promql)
	if err != nil {
		panic(err)
	}

	var region string
	var az string
	tplCtx := pongo2.Context{}

	// 选择的资源
	var resourceNames []string
	if formPolicy.Resources.Index == 3 {
		// 按对象选择
		for _, r := range formPolicy.Resources.Instances {
			resourceNames = append(resourceNames, r.Value)
			if region == "" {
				region = r.Region
			}
			if az == "" {
				az = r.Az
			}
		}
	}

	resources := strings.Join(resourceNames, "|")

	// TODO: 根据不同的模板需要添加其它的变量值
	// 填充模板中的变量值
	tplCtx["Servers"] = resources // 被选择的所有资源
	tplCtx["Region"] = region
	tplCtx["Az"] = az
	tplCtx["Operator"] = rule.Operator   // 运算符
	tplCtx["Threshold"] = rule.Threshold // 阈值

	promql, err := tpl.Execute(tplCtx)
	if err != nil {
		panic(err)
	}

	return promql
}

// 构建通用的labels和annotations
func buildCommonLablesAndAnnotations(policy *model.AlarmPolicy) (labels, annotations map[string]string) {
	labels = make(map[string]string)
	annotations = make(map[string]string)
	// 策略名称
	labels[consts.PolicyName] = policy.Name

	// 策略实例id
	labels[consts.PolicyId] = policy.InstanceId

	// 添加策略类型 form promql
	labels[consts.PolicyType] = string(policy.Type)

	// 添加业务线标签
	if policy.Production != "" {
		labels[ProductionKey] = policy.Production
	}

	// 添加策略标签
	for k, v := range policy.Labels {
		labels[k] = v
	}

	// 添加资源类型
	labels[consts.ResourceType] = policy.ResourceType
	labels[consts.ResourceSubType] = policy.ResourceSubType

	// 根据接收人设置添加annotations
	var toEmails []string
	var toSms []string
	mEmails := make(map[string]bool)
	mSms := make(map[string]bool)
	if policy.Receivers != nil {
		// 获取接收人中的邮箱和手机号码并添加到annotations
		annotations[consts.SendResolved] = strconv.FormatBool(policy.Receivers.SendResolved)
		for _, u := range policy.Receivers.NoticeUsers {
			if u.EnableEmail {
				mEmails[u.Email] = true
			}
			if u.EnableSms {
				mSms[u.Telephone] = true
			}
		}
		for _, cg := range policy.Receivers.ContactsGroup {
			if cg.EnableEmail {
				for _, u := range cg.Users {
					mEmails[u.Email] = true
				}
			}
			if cg.EnableSms {
				for _, u := range cg.Users {
					mSms[u.Telephone] = true
				}
			}
		}

		if len(mEmails) > 0 {
			for email := range mEmails {
				toEmails = append(toEmails, email)
			}
		}
		if len(mSms) > 0 {
			for tel := range mSms {
				toSms = append(toSms, tel)
			}
		}

		annotations[consts.ToEmails] = strings.Join(toEmails, ",")
		annotations[consts.ToSms] = strings.Join(toSms, ",")

	} else {
		annotations[consts.SendResolved] = "false"
	}

	// 根据通知设置添加annotations
	annotations[consts.EnableWechat] = strconv.FormatBool(policy.NotifySetup.EnableWecom)
	annotations[consts.EnableFeishu] = strconv.FormatBool(policy.NotifySetup.EnableFeishu)
	annotations[consts.EnableDingtalk] = strconv.FormatBool(policy.NotifySetup.EnableDingtalk)
	if policy.NotifySetup.EnableDingtalk {
		var dingtalkIds []string
		for _, d := range policy.NotifySetup.DingtalkRobots {
			dingtalkIds = append(dingtalkIds, strconv.Itoa(d.Id))
		}
		annotations[consts.ToDingtalk] = strings.Join(dingtalkIds, ",")
	}

	return
}

// 构建表单告警规则需要的labels和annotations
func buildLablesAndAnnotationsForFormPolicy(rule *model.AlertRule, labels, annotations map[string]string) {
	// 添加自定义标签
	if len(rule.Labels) > 0 {
		for k, v := range rule.Labels {
			labels[k] = v
		}
	}

	// 添加标签键值对
	labels[consts.AlertName] = rule.Name                 // 告警显示名称
	labels[consts.MonitorName] = rule.MonitorName        // 监控项名称
	labels[consts.MonitorDisplayName] = rule.DisplayName // 监控项显示名称
	labels[consts.RuleId] = rule.RuleId                  // 告警规则id

	// flaot64格式数据转换成字符串
	threshold := strconv.FormatFloat(rule.Threshold, 'f', 2, 64)

	// 添加annotations键值对
	annotations[consts.Level] = rule.Level                      // 告警等级
	annotations[consts.For] = rule.For                          // 持续时间
	annotations[consts.BinaryOperator] = rule.Operator          // 运算操作符
	annotations[consts.Threshold] = threshold                   // 阈值
	annotations[consts.Unit] = rule.Unit                        // 单位
	annotations[consts.RepeatInterval] = rule.Interval          // 告警间隔
	annotations[consts.GroupWait] = rule.GroupWaitTime          // 等待时间
	annotations[consts.Expression] = rule.ExpressionWithChinese // 告警规则描述

	annotations[consts.CurrentValue] = "{{ .Value }}" // 当前值
}

// 构建promql告警规则需要的labels和annotations
func buildLablesAndAnnotationsForPromqlPolicy(rule *model.PromqlPolicy, labels, annotations map[string]string) {
	// 添加annotations键值对
	labels[consts.RuleId] = rule.RuleId // 告警规则id

	annotations[consts.AlertName] = rule.Name                   // 告警显示名称
	annotations[consts.Level] = rule.Level                      // 告警等级
	annotations[consts.For] = rule.For                          // 持续时间
	annotations[consts.RepeatInterval] = rule.Interval          // 告警间隔
	annotations[consts.GroupWait] = rule.GroupWaitTime          // 等待时间
	annotations[consts.Expression] = rule.ExpressionWithChinese // 告警规则描述

	annotations[consts.CurrentValue] = "{{ .Value }}" // 当前值
}
