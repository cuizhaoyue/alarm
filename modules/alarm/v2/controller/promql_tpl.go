package controller

import (
	"fmt"
	"net/http"

	"ezone.xxxxx.com/xxxxx/xxxxx/alarm/common/consts"
	_ "ezone.xxxxx.com/xxxxx/xxxxx/alarm/modules/alarm/v2/dto"
	"ezone.xxxxx.com/xxxxx/xxxxx/alarm/response"
	"ezone.xxxxx.com/xxxxx/xxxxx/communal/lib"
	"github.com/gin-gonic/gin"
)

// ListPromqlTpl 获取PromQL模板
func (ctrl *AlarmController) ListPromqlTpl(c *gin.Context) {
	rstype := c.Query("resource_sub_type")

	if rstype == "" {
		lib.Log.TagError(lib.GetTraceContext(c), lib.DLTagHTTPFailed, map[string]interface{}{
			"message": fmt.Errorf("not allow resource_sub_type is empty").Error(),
		})
		response.Failed(c, http.StatusBadRequest, consts.ParameterError, consts.ResourceSubTypeNotAllowEmpty)
		return
	}

	data, err := ctrl.svc.PromqlTplSvc().List(c, rstype)
	if err != nil {
		lib.Log.TagError(lib.GetTraceContext(c), lib.DLTagHTTPFailed, map[string]interface{}{
			"message": fmt.Errorf("get promql templates failed, %+v", err).Error(),
		})
		response.Failed(c, http.StatusInternalServerError, consts.RequestFail, consts.PolicyGetDetailFailedMsg)
		return
	}

	response.Success(c, data)
}
