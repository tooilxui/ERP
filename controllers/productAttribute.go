package controllers

import (
	service "golangERP/services"
	"golangERP/utils"
)

// ProductAttributeContriller 城市模块
type ProductAttributeContriller struct {
	BaseController
}

// Get get attributes
func (ctl *ProductAttributeContriller) Get() {
	response := make(map[string]interface{})
	IDStr := ctl.Ctx.Input.Param(":id")
	var err error
	// 获得城市列表信息
	if IDStr == "" {
		query := make(map[string]interface{})
		exclude := make(map[string]interface{})
		cond := make(map[string]map[string]interface{})
		fields := make([]string, 0, 0)
		sortby := make([]string, 0, 0)
		order := make([]string, 0, 0)
		offsetStr := ctl.Input().Get("offset")
		var offset int64
		var limit int64 = 20
		if offsetStr != "" {
			offset, _ = utils.GetInt64(offsetStr)
		}
		limitStr := ctl.Input().Get("limit")
		if limitStr != "" {
			if limit, err = utils.GetInt64(limitStr); err != nil {
				limit = 20
			}
		}
		var attributes []map[string]interface{}
		var paginator utils.Paginator
		var access utils.AccessResult
		if access, paginator, attributes, err = service.ServiceGetProductAttribute(&ctl.User, query, exclude, cond, fields, sortby, order, offset, limit); err == nil {
			response["code"] = utils.SuccessCode
			response["msg"] = utils.SuccessMsg
			data := make(map[string]interface{})
			data["attributes"] = &attributes
			data["paginator"] = &paginator
			data["access"] = access
			response["data"] = data
		} else {
			response["code"] = utils.FailedCode
			response["msg"] = utils.FailedMsg
			response["err"] = err
		}
	} else {
		// 获得某个城市的信息
		if attributeID, err := utils.GetInt64(IDStr); err == nil {
			if access, attribute, err := service.ServiceGetProductAttributeByID(&ctl.User, attributeID); err == nil {
				response["code"] = utils.SuccessCode
				response["msg"] = utils.SuccessMsg
				data := make(map[string]interface{})
				data["attribute"] = &attribute
				data["access"] = access
				response["data"] = data
			} else {
				response["code"] = utils.FailedCode
				response["msg"] = utils.FailedMsg
				response["err"] = err
			}
		}
	}
	ctl.Data["json"] = response

	ctl.ServeJSON()
}