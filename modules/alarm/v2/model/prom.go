package model

import (
	monitoringv1 "github.com/prometheus-operator/prometheus-operator/pkg/apis/monitoring/v1"
	monitoringv1alpha1 "github.com/prometheus-operator/prometheus-operator/pkg/apis/monitoring/v1alpha1"
	clientversioned "github.com/prometheus-operator/prometheus-operator/pkg/client/versioned"
	typedmonitoringv1alpha1 "github.com/prometheus-operator/prometheus-operator/pkg/client/versioned/typed/monitoring/v1alpha1"
)

const (
	PrometheusRuleKind    = monitoringv1.PrometheusRuleKind
	PrometheusRuleName    = monitoringv1.PrometheusRuleName
	PrometheusRuleKindKey = monitoringv1.PrometheusRuleKindKey
)

type (
	PrometheusRule     = monitoringv1.PrometheusRule
	PrometheusRuleSpec = monitoringv1.PrometheusRuleSpec
	RuleGroup          = monitoringv1.RuleGroup
	Rule               = monitoringv1.Rule
	Duration           = monitoringv1.Duration
)

type (
	AlertmanagerConfig         = monitoringv1alpha1.AlertmanagerConfig
	AlertmanagerConfigList     = monitoringv1alpha1.AlertmanagerConfigList
	AlertmanagerConfigInterfac = typedmonitoringv1alpha1.AlertmanagerConfigInterface
	AlertmanagerConfigSpec     = monitoringv1alpha1.AlertmanagerConfigSpec
	Receiver                   = monitoringv1alpha1.Receiver
	Matcher                    = monitoringv1alpha1.Matcher
	Route                      = monitoringv1alpha1.Route
	WebhookConfig              = monitoringv1alpha1.WebhookConfig
	Clientset                  = clientversioned.Clientset
)
