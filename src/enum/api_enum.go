package enum

type api struct {
	Operation  string `json:"operation"`
	Function   string `json:"function"`
	AlertLevel int    `json:"alertLevel"`
}

var apiMap = make(map[string]api)

/**
以map實現enum，存放adminLog初始化需要的固定資訊
*/
func init() {
	addApiEnum(CreatePolicyPath, "policy", "createPolicy", 2)
	addApiEnum(UpdatePolicyPath, "policy", "updatePolicy", 2)
	addApiEnum(RemovePolicyPath, "policy", "removePolicy", 2)
}

func addApiEnum(path string, operation string, function string, alertLevel int) {
	api := new(api)
	api.Operation = operation
	api.Function = function
	api.AlertLevel = alertLevel
	apiMap[path] = *api
}

func GetApiMap() map[string]api {
	return apiMap
}
