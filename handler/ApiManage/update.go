package ApiManage

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"metadata/constant"
	"metadata/dal"
	"metadata/model"
	"metadata/util"
)

type ApiUpdateRequest struct {
	Id             int64  `form:"Id" binding:"required"`
	Path           string `form:"Path"`
	ServiceId      int64  `form:"ServiceId"`
	Name           string `form:"Name"`
	Protocol       string `form:"Protocol"`
	ConnectTimeout int    `form:"ConnectTimeout"`
	Retries        int    `form:"Retries"`
	Status         string `form:"Status"`
	Tags           string `form:"Tags"`
	Method         string `form:"Method"`
}

func Update(c *gin.Context) {

	var apiResponse ApiUpdateRequest
	if err := c.ShouldBindQuery(&apiResponse); err != nil {
		logrus.Errorf("parameter invalid %v", err.Error())
		util.ResponseError(c, 401, constant.PARAMETER_INVALID, "parameter invalid")
		return
	}
	var apiList []model.ApiStruct
	err, count := dal.ListApi(c, 0, 0, "", "", "", apiResponse.Id, apiResponse.ServiceId, &apiList)
	if err != nil {
		logrus.Errorf("search failed %v", err.Error())
		util.ResponseError(c, 500, constant.SEARCH_FAILED, "search failed")
		return
	}
	if count == 0 {
		util.ResponseError(c, 401, constant.SEARCH_NOT_FOUND, "search not found")
		return
	}
	api := apiList[0]
	if apiResponse.Tags != "" {
		api.Tags = apiResponse.Tags
	}
	if apiResponse.Method != "" {
		api.Method = apiResponse.Method
	}
	if apiResponse.Name != "" {
		api.Name = apiResponse.Name
	}
	if apiResponse.Path != "" {
		api.Path = apiResponse.Path
	}
	if apiResponse.Protocol != "" {
		api.Protocol = apiResponse.Protocol
	}
	if apiResponse.Status != "" {
		api.Status = apiResponse.Status
	}
	if apiResponse.ConnectTimeout != 0 {
		api.ConnectTimeout = apiResponse.ConnectTimeout
	}
	if apiResponse.Retries != 0 {
		api.Retries = apiResponse.Retries
	}

	err = dal.UpdateApi(c, api)
	if err != nil {
		logrus.Errorf("update api failed %v", err.Error())
		util.ResponseError(c, 500, constant.UPDATE_FAILED, "update api failed")
		return
	}

	util.ResponseSuccess(c, "success")
}
