package controller

import (
	"fmt"
	"net/http"

	"ezone.xxxxx.com/xxxxx/xxxxx/alarm/common/consts"
	"ezone.xxxxx.com/xxxxx/xxxxx/alarm/modules/alarm/v2/dto"
	"ezone.xxxxx.com/xxxxx/xxxxx/alarm/modules/alarm/v2/model"
	"ezone.xxxxx.com/xxxxx/xxxxx/alarm/response"
	"ezone.xxxxx.com/xxxxx/xxxxx/communal/lib"
	"github.com/gin-gonic/gin"
)

// Save 告警列表
func (ctrl *AlarmController) Save(c *gin.Context) {
	var req model.NotificationData
	if err := c.ShouldBindJSON(&req); err != nil {
		lib.Log.TagError(lib.GetTraceContext(c), lib.DLTagHTTPFailed, map[string]interface{}{
			"message": fmt.Errorf("bind body params failed when save alert, %s", err).Error(),
		})
	}

	err := ctrl.svc.AlertSvc().CreateOrUpdate(c, &req)
	if err != nil {
		lib.Log.TagError(lib.GetTraceContext(c), lib.DLTagHTTPFailed, map[string]interface{}{
			"message": fmt.Errorf("get the alert list failed, %+v", err).Error(),
		})
	}
}

// AlertOverview 告警总览
func (ctrl *AlarmController) AlertOverview(c *gin.Context) {
	var req dto.OverviewRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		lib.Log.TagError(lib.GetTraceContext(c), lib.DLTagHTTPFailed, map[string]interface{}{
			"message": fmt.Errorf("bind body params failed, %s", err).Error(),
		})
		response.Failed(c, http.StatusBadRequest, consts.ParameterError, fmt.Errorf("参数绑定失败: %s", err).Error())
		return
	}

	data, err := ctrl.svc.AlertSvc().Overview(c, &req)
	if err != nil {
		lib.Log.TagError(lib.GetTraceContext(c), lib.DLTagHTTPFailed, map[string]interface{}{
			"message": fmt.Errorf("get the alert overview failed, %+v", err).Error(),
		})
		response.Failed(c, http.StatusInternalServerError, consts.RequestFail, consts.PolicyListFailedMsg)
		return
	}

	response.Success(c, data)

}

// AlertList 告警列表
func (ctrl *AlarmController) AlertList(c *gin.Context) {
	var req dto.ListAlertsRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		lib.Log.TagError(lib.GetTraceContext(c), lib.DLTagHTTPFailed, map[string]interface{}{
			"message": fmt.Errorf("bind body params failed, %s", err).Error(),
		})
		response.Failed(c, http.StatusBadRequest, consts.ParameterError, fmt.Errorf("参数绑定失败: %s", err).Error())
		return
	}

	data, err := ctrl.svc.AlertSvc().List(c, &req)
	if err != nil {
		lib.Log.TagError(lib.GetTraceContext(c), lib.DLTagHTTPFailed, map[string]interface{}{
			"message": fmt.Errorf("get the alert list failed, %+v", err).Error(),
		})
		response.Failed(c, http.StatusInternalServerError, consts.RequestFail, consts.PolicyListFailedMsg)
		return
	}

	response.Success(c, data)
}
