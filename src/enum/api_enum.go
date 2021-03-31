package enum

import (
	"net/http"
)

type Api struct {
	Operation     string
	Function      string
	Group         string
	RequestMethod string
	ModelName     string
	AlertLevel    int
}

var apiMap = make(map[string]Api)

/**
以map實現enum，存放adminLog初始化需要的固定資訊
*/
func init() {
	addApiEnum(GroupApiV1, http.MethodPost, CreatePolicyPath, OperationPolicy, "createPolicy", 2)
	addApiEnum(GroupApiV1, http.MethodPut, UpdatePolicyContentPath, OperationPolicy, "updatePolicy", 2)
	addApiEnum(GroupApiV1, http.MethodDelete, DeletePolicyPath, OperationPolicy, "removePolicy", 2)
}

func addApiEnum(group string, requestMethod string, path string, operation string, function string, alertLevel int) {
	api := new(Api)
	api.Group = group
	api.RequestMethod = requestMethod
	api.Operation = operation
	api.Function = function
	api.AlertLevel = alertLevel
	key := requestMethod + "_" + group + path
	apiMap[key] = *api
}

func GetApiMap() map[string]Api {
	return apiMap
}
