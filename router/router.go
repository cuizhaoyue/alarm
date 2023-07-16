package router

import (
	"ezone.xxxxx.com/xxxxx/xxxxx/alarm/modules/alarm/v2/controller"
	contactCtrl "ezone.xxxxx.com/xxxxx/xxxxx/alarm/modules/contacts/controller"
	"github.com/gin-gonic/gin"
)

func InitRouter(engine *gin.Engine) {
	// 告警策略接口
	alarm := engine.Group("/alarm")

	v2 := alarm.Group("/v2")
	alarmController := controller.NewAlarmController()
	{
		policy := v2.Group("/policy")
		{
			// 创建告警策略
			policy.POST("/create", alarmController.Create)
			// 获取告警策略列表
			policy.POST("/list", alarmController.List)
			// 获取告警策略详情
			policy.GET("/describe", alarmController.Get)
			// 删除告警策略
			policy.POST("/delete", alarmController.DeleteCollection)
			// 更新告警策略
			policy.POST("/update", alarmController.Update)
			// 启动或停止告警策略
			policy.POST("/switch", alarmController.Switch)
		}

		remoteConfig := v2.Group("/remoteconfig")
		{
			// 获取短信配置
			remoteConfig.GET("/sms", alarmController.GetSMSConfig)
			// 更新短信配置
			remoteConfig.POST("/sms", alarmController.UpdateSMSConfig)

			// 获取邮箱配置
			remoteConfig.GET("/mailbox", alarmController.GetMailboxConfig)
			// 更新邮箱配置
			remoteConfig.POST("/mailbox", alarmController.UpdateMailboxConfig)

			// 获取企业微信配置
			remoteConfig.GET("/wechat", alarmController.GetWechatConfig)
			// 更新企业微信配置
			remoteConfig.POST("/wechat", alarmController.UpdateWechatConfig)

			// 获取飞书配置
			remoteConfig.GET("/feishu", alarmController.GetFeishuConfig)
			// 更新飞书配置
			remoteConfig.POST("/feishu", alarmController.UpdateFeishuConfig)

			// 获取钉钉配置列表
			remoteConfig.POST("/dingtalk/list", alarmController.ListDingtalkConfig)
			// 创建钉钉配置
			remoteConfig.POST("/dingtalk/create", alarmController.CreateDingtalkConfig)
			// 更新钉钉配置
			remoteConfig.POST("/dingtalk/update", alarmController.UpdateDingtalkConfig)
			// 删除钉钉配置
			remoteConfig.POST("/dingtalk/delete", alarmController.DeleteDingtalkConfig)
		}

		alert := v2.Group("/alert")
		{
			// 保存告警消息
			alert.POST("/save", alarmController.Save)

			// 获取告警总览
			alert.POST("/overview", alarmController.AlertOverview)

			// 获取告警列表
			alert.POST("/list", alarmController.AlertList)
		}
	}

	// 告警模板
	promqlTpl := v2.Group("/promql")
	{
		promqlTpl.GET("/list", alarmController.ListPromqlTpl) // 获取promql告警模板列表
	}

	// 联系人组接口
	contact := engine.Group("/contact")
	contactController := contactCtrl.NewContactsController()
	{
		// 创建联系人组
		contact.POST("/create", contactController.Create)
		// 删除联系人组
		contact.POST("/delete", contactController.Delete)
		// 获取联系人组列表
		contact.POST("/list", contactController.List)
		// 从联系人组中移除用户
		contact.POST("/remove", contactController.RemoveUserFromContact)
		// 编辑联系人组
		contact.POST("/update", contactController.Update)
	}
}
