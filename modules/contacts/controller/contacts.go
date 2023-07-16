package controller

import (
	"fmt"
	"net/http"

	"ezone.xxxxx.com/xxxxx/xxxxx/alarm/common/consts"
	"ezone.xxxxx.com/xxxxx/xxxxx/alarm/modules/contacts/dto"
	"ezone.xxxxx.com/xxxxx/xxxxx/alarm/response"
	"ezone.xxxxx.com/xxxxx/xxxxx/communal/lib"
	"github.com/gin-gonic/gin"
)

// Create 创建联系人组
func (s *ContactsController) Create(c *gin.Context) {
	var req dto.CreateContactRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		lib.Log.TagError(lib.GetTraceContext(c), lib.DLTagHTTPFailed, map[string]interface{}{
			"message": fmt.Errorf("bind body params failed, %s", err).Error(),
		})
		response.Failed(c, http.StatusBadRequest, consts.ParameterError, fmt.Errorf("参数绑定失败: %s", err).Error())
		return
	}

	err := s.svc.Contacts().Create(c, &req)
	if err != nil {
		lib.Log.TagError(lib.GetTraceContext(c), lib.DLTagHTTPFailed, map[string]interface{}{
			"message": fmt.Errorf("create contacts failed, %v", err).Error(),
		})
		response.Failed(c, http.StatusInternalServerError, consts.ParameterError, "创建失败")
		return
	}

	response.Success(c, nil)
}

// Update 编辑联系人组
func (s *ContactsController) Update(c *gin.Context) {
	var req dto.UpdateContactRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		lib.Log.TagError(lib.GetTraceContext(c), lib.DLTagHTTPFailed, map[string]interface{}{
			"message": fmt.Errorf("bind body params failed, %s", err).Error(),
		})
		response.Failed(c, http.StatusBadRequest, consts.ParameterError, fmt.Errorf("参数绑定失败: %s", err).Error())
		return
	}

	err := s.svc.Contacts().Update(c, &req)
	if err != nil {
		lib.Log.TagError(lib.GetTraceContext(c), lib.DLTagHTTPFailed, map[string]interface{}{
			"message": fmt.Errorf("update contacts failed, %v", err).Error(),
		})
		response.Failed(c, http.StatusInternalServerError, consts.ParameterError, "更新失败")
		return
	}

	response.Success(c, nil)
}

// List 获取联系人组列表
func (s *ContactsController) List(c *gin.Context) {
	var req dto.ListContactRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		lib.Log.TagError(lib.GetTraceContext(c), lib.DLTagHTTPFailed, map[string]interface{}{
			"message": fmt.Errorf("bind body params failed, %s", err).Error(),
		})
		response.Failed(c, http.StatusBadRequest, consts.ParameterError, fmt.Errorf("参数绑定失败: %s", err).Error())
		return
	}

	respData, err := s.svc.Contacts().List(c, &req)
	if err != nil {
		lib.Log.TagError(lib.GetTraceContext(c), lib.DLTagHTTPFailed, map[string]interface{}{
			"message": fmt.Errorf("list contacts failed, %v", err).Error(),
		})
		response.Failed(c, http.StatusInternalServerError, consts.ParameterError, "获取联系人组列表失败")
		return
	}

	response.Success(c, respData)
}

// Delete 删除联系人组
func (s *ContactsController) Delete(c *gin.Context) {
	var req dto.DeleteContactRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		lib.Log.TagError(lib.GetTraceContext(c), lib.DLTagHTTPFailed, map[string]interface{}{
			"message": fmt.Errorf("bind body params failed, %s", err).Error(),
		})
		response.Failed(c, http.StatusBadRequest, consts.ParameterError, fmt.Errorf("参数绑定失败: %s", err).Error())
		return
	}

	err := s.svc.Contacts().Delete(c, &req)
	if err != nil {
		lib.Log.TagError(lib.GetTraceContext(c), lib.DLTagHTTPFailed, map[string]interface{}{
			"message": fmt.Errorf("delete contacts failed, %v", err).Error(),
		})
		response.Failed(c, http.StatusInternalServerError, consts.ParameterError, "删除联系人组失败")
		return
	}

	response.Success(c, nil)
}

// RemoveUserFromContact 从联系人组中移除用户
func (s *ContactsController) RemoveUserFromContact(c *gin.Context) {
	var req dto.RemoveUserFromContactRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		lib.Log.TagError(lib.GetTraceContext(c), lib.DLTagHTTPFailed, map[string]interface{}{
			"message": fmt.Errorf("bind body params failed, %s", err).Error(),
		})
		response.Failed(c, http.StatusBadRequest, consts.ParameterError, fmt.Errorf("参数绑定失败: %s", err).Error())
		return
	}

	err := s.svc.Contacts().RemoveUserFromContact(c, &req)
	if err != nil {
		lib.Log.TagError(lib.GetTraceContext(c), lib.DLTagHTTPFailed, map[string]interface{}{
			"message": fmt.Errorf("remove user from contacts failed, %v", err).Error(),
		})
		response.Failed(c, http.StatusInternalServerError, consts.ParameterError, "移除联系人失败")
		return
	}

	response.Success(c, nil)
}
