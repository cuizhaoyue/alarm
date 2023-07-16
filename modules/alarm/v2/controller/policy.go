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

// Create 创建告警策略
func (ctrl *AlarmController) Create(c *gin.Context) {
	var req dto.CreatePolicyRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		lib.Log.TagError(lib.GetTraceContext(c), lib.DLTagHTTPFailed, map[string]interface{}{
			"message": fmt.Errorf("bind body params failed, %s", err).Error(),
		})
		response.Failed(c, http.StatusBadRequest, consts.ParameterError, fmt.Errorf("参数绑定失败: %s", err).Error())
		return
	}

	err := ctrl.svc.AlarmSvc().Create(c, &req)
	if err != nil {
		lib.Log.TagError(lib.GetTraceContext(c), lib.DLTagHTTPFailed, map[string]interface{}{
			"message": fmt.Errorf("create contacts failed, %+v", err).Error(),
		})
		response.Failed(c, http.StatusInternalServerError, consts.RequestFail, consts.PolicyCreateFailedMsg)
		return
	}

	response.Success(c, nil)
}

// List 创建告警策略
func (ctrl *AlarmController) List(c *gin.Context) {
	var req dto.ListPolicyRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		lib.Log.TagError(lib.GetTraceContext(c), lib.DLTagHTTPFailed, map[string]interface{}{
			"message": fmt.Errorf("bind body params failed, %s", err).Error(),
		})
		response.Failed(c, http.StatusBadRequest, consts.ParameterError, fmt.Errorf("参数绑定失败: %s", err).Error())
		return
	}

	data, err := ctrl.svc.AlarmSvc().List(c, &req)
	if err != nil {
		lib.Log.TagError(lib.GetTraceContext(c), lib.DLTagHTTPFailed, map[string]interface{}{
			"message": fmt.Errorf("get the policy list failed, %+v", err).Error(),
		})
		response.Failed(c, http.StatusInternalServerError, consts.RequestFail, consts.PolicyListFailedMsg)
		return
	}

	response.Success(c, data)
}

// Get 获取告警策略详情
func (ctrl *AlarmController) Get(c *gin.Context) {
	// 获取告警策略id
	insId := c.Query("InstanceId")
	if insId == "" {
		lib.Log.TagError(lib.GetTraceContext(c), lib.DLTagHTTPFailed, map[string]interface{}{
			"message": fmt.Errorf("not allow the policy instance id is empty").Error(),
		})
		response.Failed(c, http.StatusBadRequest, consts.ParameterError, consts.PolicyIdNotAllowEmpty)
		return
	}

	data, err := ctrl.svc.AlarmSvc().Get(c, insId)
	if err != nil {
		lib.Log.TagError(lib.GetTraceContext(c), lib.DLTagHTTPFailed, map[string]interface{}{
			"message": fmt.Errorf("get policy [%s] detail failed, %+v", insId, err).Error(),
		})
		response.Failed(c, http.StatusInternalServerError, consts.RequestFail, consts.PolicyGetDetailFailedMsg)
		return
	}

	response.Success(c, data)
}

// DeleteCollection 删除告警策略
func (ctrl *AlarmController) DeleteCollection(c *gin.Context) {
	var req dto.DeletePolicyRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		lib.Log.TagError(lib.GetTraceContext(c), lib.DLTagHTTPFailed, map[string]interface{}{
			"message": fmt.Errorf("bind body params failed, %s", err).Error(),
		})
		response.Failed(c, http.StatusBadRequest, consts.ParameterError, fmt.Errorf("参数绑定失败: %s", err).Error())
		return
	}

	err := ctrl.svc.AlarmSvc().DeleteCollection(c, &req)
	if err != nil {
		lib.Log.TagError(lib.GetTraceContext(c), lib.DLTagHTTPFailed, map[string]interface{}{
			"message": fmt.Errorf("delete policy failed, %+v", err).Error(),
		})
		response.Failed(c, http.StatusBadRequest, consts.ParameterError, consts.PolicyDeleteFailedMsg)
		return
	}

	response.Success(c, nil)
}

// Update 编辑更新告警策略
func (ctrl *AlarmController) Update(c *gin.Context) {
	var req dto.UpdatePolicyRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		lib.Log.TagError(lib.GetTraceContext(c), lib.DLTagHTTPFailed, map[string]interface{}{
			"message": fmt.Errorf("bind body params failed, %s", err).Error(),
		})
		response.Failed(c, http.StatusBadRequest, consts.ParameterError, fmt.Errorf("参数绑定失败: %s", err).Error())
		return
	}

	err := ctrl.svc.AlarmSvc().Update(c, &req)
	if err != nil {
		lib.Log.TagError(lib.GetTraceContext(c), lib.DLTagHTTPFailed, map[string]interface{}{
			"message": fmt.Errorf("update policy failed, %+v", err).Error(),
		})
		response.Failed(c, http.StatusBadRequest, consts.ParameterError, consts.PolicyUpdateFailedMsg)
		return
	}

	response.Success(c, nil)
}

// Switch 启动或停止告警策略
func (ctrl *AlarmController) Switch(c *gin.Context) {
	var req dto.SwitchPolicyRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		lib.Log.TagError(lib.GetTraceContext(c), lib.DLTagHTTPFailed, map[string]interface{}{
			"message": fmt.Errorf("bind body params failed, %s", err).Error(),
		})
		response.Failed(c, http.StatusBadRequest, consts.ParameterError, fmt.Errorf("参数绑定失败: %s", err).Error())
		return
	}

	err := ctrl.svc.AlarmSvc().Switch(c, &req)
	if err != nil {
		lib.Log.TagError(lib.GetTraceContext(c), lib.DLTagHTTPFailed, map[string]interface{}{
			"message": fmt.Errorf("switch policy status failed, %+v", err).Error(),
		})
		response.Failed(c, http.StatusBadRequest, consts.ParameterError, consts.PolicySwitchFailedMsg)
		return
	}

	response.Success(c, nil)
}
