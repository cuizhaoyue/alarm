package controller

import (
	"fmt"
	"net/http"

	"ezone.xxxxx.com/xxxxx/xxxxx/alarm/common/consts"
	"ezone.xxxxx.com/xxxxx/xxxxx/alarm/modules/alarm/v2/dto"
	"ezone.xxxxx.com/xxxxx/xxxxx/alarm/response"
	"ezone.xxxxx.com/xxxxx/xxxxx/communal/lib"
	"github.com/gin-gonic/gin"
)

// GetSMSConfig 获取短信配置
func (ctrl *AlarmController) GetSMSConfig(c *gin.Context) {
	data, err := ctrl.svc.RemoteConfigSvc().GetSMSConfig(c)
	if err != nil {
		lib.Log.TagError(lib.GetTraceContext(c), lib.DLTagHTTPFailed, map[string]interface{}{
			"message": fmt.Errorf("get sms config failed, %+v", err).Error(),
		})
		response.Failed(c, http.StatusInternalServerError, consts.RequestFail, consts.GetSMSConfigFailedMsg)
		return
	}

	response.Success(c, data)
}

// UpdateSMSConfig 更新短信配置
func (ctrl *AlarmController) UpdateSMSConfig(c *gin.Context) {
	var req dto.UpdateSMSConfigRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		lib.Log.TagError(lib.GetTraceContext(c), lib.DLTagHTTPFailed, map[string]interface{}{
			"message": fmt.Errorf("bind body params failed, %s", err).Error(),
		})
		response.Failed(c, http.StatusBadRequest, consts.ParameterError, fmt.Errorf("参数绑定失败: %s", err).Error())
		return
	}

	err := ctrl.svc.RemoteConfigSvc().CreateOrUpdateSMSConfig(c, &req)
	if err != nil {
		lib.Log.TagError(lib.GetTraceContext(c), lib.DLTagHTTPFailed, map[string]interface{}{
			"message": fmt.Errorf("update sms config failed, %+v", err).Error(),
		})
		response.Failed(c, http.StatusInternalServerError, consts.RequestFail, consts.UpdateSMSConfigFailedMsg)
		return
	}

	response.Success(c, nil)
}

// GetMailboxConfig 获取邮箱配置
func (ctrl *AlarmController) GetMailboxConfig(c *gin.Context) {
	data, err := ctrl.svc.RemoteConfigSvc().GetMailboxConfig(c)
	if err != nil {
		lib.Log.TagError(lib.GetTraceContext(c), lib.DLTagHTTPFailed, map[string]interface{}{
			"message": fmt.Errorf("get mailbox config failed, %+v", err).Error(),
		})
		response.Failed(c, http.StatusInternalServerError, consts.RequestFail, consts.GetMailboxConfigFailedMsg)
		return
	}

	response.Success(c, data)
}

// UpdateMailboxConfig 更新邮箱配置
func (ctrl *AlarmController) UpdateMailboxConfig(c *gin.Context) {
	var req dto.UpdateMailboxConfigRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		lib.Log.TagError(lib.GetTraceContext(c), lib.DLTagHTTPFailed, map[string]interface{}{
			"message": fmt.Errorf("bind body params failed, %s", err).Error(),
		})
		response.Failed(c, http.StatusBadRequest, consts.ParameterError, fmt.Errorf("参数绑定失败: %s", err).Error())
		return
	}

	err := ctrl.svc.RemoteConfigSvc().CreateOrUpdateMailboxConfig(c, &req)
	if err != nil {
		lib.Log.TagError(lib.GetTraceContext(c), lib.DLTagHTTPFailed, map[string]interface{}{
			"message": fmt.Errorf("update mailbox config failed, %+v", err).Error(),
		})
		response.Failed(c, http.StatusInternalServerError, consts.RequestFail, consts.UpdateMailboxConfigFailedMsg)
		return
	}

	response.Success(c, nil)
}

// GetWechatConfig 获取企业微信配置
func (ctrl *AlarmController) GetWechatConfig(c *gin.Context) {
	data, err := ctrl.svc.RemoteConfigSvc().GetWechatConfig(c)
	if err != nil {
		lib.Log.TagError(lib.GetTraceContext(c), lib.DLTagHTTPFailed, map[string]interface{}{
			"message": fmt.Errorf("get Wechat config failed, %+v", err).Error(),
		})
		response.Failed(c, http.StatusInternalServerError, consts.RequestFail, consts.GetWechatConfigFailedMsg)
		return
	}

	response.Success(c, data)
}

// UpdateWechatConfig 更新企业微信配置
func (ctrl *AlarmController) UpdateWechatConfig(c *gin.Context) {
	var req dto.UpdateWechatConfigRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		lib.Log.TagError(lib.GetTraceContext(c), lib.DLTagHTTPFailed, map[string]interface{}{
			"message": fmt.Errorf("bind body params failed, %s", err).Error(),
		})
		response.Failed(c, http.StatusBadRequest, consts.ParameterError, fmt.Errorf("参数绑定失败: %s", err).Error())
		return
	}

	err := ctrl.svc.RemoteConfigSvc().CreateOrUpdateWechatConfig(c, &req)
	if err != nil {
		lib.Log.TagError(lib.GetTraceContext(c), lib.DLTagHTTPFailed, map[string]interface{}{
			"message": fmt.Errorf("update Wechat config failed, %+v", err).Error(),
		})
		response.Failed(c, http.StatusInternalServerError, consts.RequestFail, consts.UpdateWechatConfigFailedMsg)
		return
	}

	response.Success(c, nil)
}

// GetFeishuConfig 获取企业微信配置
func (ctrl *AlarmController) GetFeishuConfig(c *gin.Context) {
	data, err := ctrl.svc.RemoteConfigSvc().GetFeishuConfig(c)
	if err != nil {
		lib.Log.TagError(lib.GetTraceContext(c), lib.DLTagHTTPFailed, map[string]interface{}{
			"message": fmt.Errorf("get Feishu config failed, %+v", err).Error(),
		})
		response.Failed(c, http.StatusInternalServerError, consts.RequestFail, consts.GetFeishuConfigFailedMsg)
		return
	}

	response.Success(c, data)
}

// UpdateFeishuConfig 更新企业微信配置
func (ctrl *AlarmController) UpdateFeishuConfig(c *gin.Context) {
	var req dto.UpdateFeishuConfigRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		lib.Log.TagError(lib.GetTraceContext(c), lib.DLTagHTTPFailed, map[string]interface{}{
			"message": fmt.Errorf("bind body params failed, %s", err).Error(),
		})
		response.Failed(c, http.StatusBadRequest, consts.ParameterError, fmt.Errorf("参数绑定失败: %s", err).Error())
		return
	}

	err := ctrl.svc.RemoteConfigSvc().CreateOrUpdateFeishuConfig(c, &req)
	if err != nil {
		lib.Log.TagError(lib.GetTraceContext(c), lib.DLTagHTTPFailed, map[string]interface{}{
			"message": fmt.Errorf("update Feishu config failed, %+v", err).Error(),
		})
		response.Failed(c, http.StatusInternalServerError, consts.RequestFail, consts.UpdateFeishuConfigFailedMsg)
		return
	}

	response.Success(c, nil)
}

// ListDingtalkConfig 获取钉钉配置列表
func (ctrl *AlarmController) ListDingtalkConfig(c *gin.Context) {
	var req dto.ListDingtalkConfigRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		lib.Log.TagError(lib.GetTraceContext(c), lib.DLTagHTTPFailed, map[string]interface{}{
			"message": fmt.Errorf("bind body params failed, %s", err).Error(),
		})
		response.Failed(c, http.StatusBadRequest, consts.ParameterError, fmt.Errorf("参数绑定失败: %s", err).Error())
		return
	}

	data, err := ctrl.svc.RemoteConfigSvc().ListDingtalkConfig(c, &req)
	if err != nil {
		lib.Log.TagError(lib.GetTraceContext(c), lib.DLTagHTTPFailed, map[string]interface{}{
			"message": fmt.Errorf("get dingtalk config list failed, %s", err).Error(),
		})
		response.Failed(c, http.StatusInternalServerError, consts.RequestFail, consts.ListDingtalkConfigFailedMsg)
		return
	}

	response.Success(c, data)
}

// CreateDingtalkConfig 创建钉钉配置
func (ctrl *AlarmController) CreateDingtalkConfig(c *gin.Context) {
	var req dto.CreateDingtalkConfigRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		lib.Log.TagError(lib.GetTraceContext(c), lib.DLTagHTTPFailed, map[string]interface{}{
			"message": fmt.Errorf("bind body params failed, %s", err).Error(),
		})
		response.Failed(c, http.StatusBadRequest, consts.ParameterError, fmt.Errorf("参数绑定失败: %s", err).Error())
		return
	}

	err := ctrl.svc.RemoteConfigSvc().CreateDingtalkConfig(c, &req)
	if err != nil {
		lib.Log.TagError(lib.GetTraceContext(c), lib.DLTagHTTPFailed, map[string]interface{}{
			"message": fmt.Errorf("create dingtalk config failed, %s", err).Error(),
		})
		response.Failed(c, http.StatusInternalServerError, consts.RequestFail, consts.CreateDingtalkConfigFailedMsg)
		return
	}

	response.Success(c, nil)
}

// UpdateDingtalkConfig 更新钉钉配置
func (ctrl *AlarmController) UpdateDingtalkConfig(c *gin.Context) {
	var req dto.UpdateDingtalkConfigRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		lib.Log.TagError(lib.GetTraceContext(c), lib.DLTagHTTPFailed, map[string]interface{}{
			"message": fmt.Errorf("bind body params failed, %s", err).Error(),
		})
		response.Failed(c, http.StatusBadRequest, consts.ParameterError, fmt.Errorf("参数绑定失败: %s", err).Error())
		return
	}

	err := ctrl.svc.RemoteConfigSvc().UpdateDingtalkConfig(c, &req)
	if err != nil {
		lib.Log.TagError(lib.GetTraceContext(c), lib.DLTagHTTPFailed, map[string]interface{}{
			"message": fmt.Errorf("update dingtalk config failed, %s", err).Error(),
		})
		response.Failed(c, http.StatusInternalServerError, consts.RequestFail, consts.UpdateDingtalkConfigFailedMsg)
		return
	}

	response.Success(c, nil)
}

// DeleteDingtalkConfig 删除钉钉配置
func (ctrl *AlarmController) DeleteDingtalkConfig(c *gin.Context) {
	var req dto.DeleteDingtalkConfigRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		lib.Log.TagError(lib.GetTraceContext(c), lib.DLTagHTTPFailed, map[string]interface{}{
			"message": fmt.Errorf("bind body params failed, %s", err).Error(),
		})
		response.Failed(c, http.StatusBadRequest, consts.ParameterError, fmt.Errorf("参数绑定失败: %s", err).Error())
		return
	}

	err := ctrl.svc.RemoteConfigSvc().DeleteDingtalkConfig(c, &req)
	if err != nil {
		lib.Log.TagError(lib.GetTraceContext(c), lib.DLTagHTTPFailed, map[string]interface{}{
			"message": fmt.Errorf("delete dingtalk config failed, %s", err).Error(),
		})
		response.Failed(c, http.StatusInternalServerError, consts.RequestFail, consts.DeleteDingtalkConfigFailedMsg)
		return
	}

	response.Success(c, nil)
}
