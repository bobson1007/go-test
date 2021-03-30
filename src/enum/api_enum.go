package enum

import (
	"reflect"
)

type Api struct {
	Operation  string `json:"operation"`
	Function   string `json:"function"`
	AlertLevel int    `json:"alertLevel"`
	ModelName  string
}

var apiMap = make(map[string]Api)

/**
以map實現enum，存放adminLog初始化需要的固定資訊
*/
func init() {
	addApiEnum(CreatePolicyPath, "policy", "createPolicy", 2)
	addApiEnum(UpdatePolicyPath, "policy", "updatePolicy", 2)
	addApiEnum(RemovePolicyPath, "policy", "removePolicy", 2)
	addApiEnum(UpdatePolicyName, "policy", "removePolicy", 2)
}

func addApiEnum(path string, operation string, function string, alertLevel int) {
	api := new(Api)
	api.Operation = operation
	api.Function = function
	api.AlertLevel = alertLevel
	//Api.ModelName = reflect.TypeOf(model).Name()
	apiMap[path] = *api

	//registerType(reflect.TypeOf(model))

}

func GetApiMap() map[string]Api {
	return apiMap
}

////////////////// 反射註冊 ////////////////////////

//this is the registry of types by name
var registry = map[string]reflect.Type{}

// add a type to the registry
func registerType(t reflect.Type) {
	name := t.Name()
	registry[name] = t
}

// create a new object by name, returning it as interface{}
func NewByName(name string) interface{} {

	t, found := registry[name]
	if !found {
		panic("name not found!")
	}

	return reflect.New(t).Elem().Interface()
}
