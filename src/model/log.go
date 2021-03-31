package model

import (
	//"bytes"
	//"license_service/app/handler"
	//"license_service/config"
	//"license_service/utils"

	"github.com/gin-gonic/gin"
)

// lincense-service內的AdminLog模型
type AdminLog struct {
	Operation     string      `json:"operation"`
	Function      string      `json:"function"`
	AlertLevel    int         `json:"alertLevel"`
	IP            string      `json:"IP"`
	AdminID       int         `json:"adminId"`
	AdminAccount  string      `json:"adminAccount"`
	OriginalValue interface{} `json:"originalValue,omitempty"`
	UpdateValue   interface{} `json:"updateValue,omitempty"`
	TargetID      int         `json:"targetId,omitempty"`
	TargetName    string      `json:"targetName,omitempty"`
	Detail        string      `json:"detail,omitempty"`
	Notes         string      `json:"notes,omitempty"`
	Result        int         `json:"result"`
}

// Init - Function,Operation,AlertLevel
func (a *AdminLog) Init(module, oper string, level int) {
	a.Function = module
	a.Operation = oper
	a.AlertLevel = level
}

// Base - AdminID,AdminAccount,IP
func (a *AdminLog) Base(ctx *gin.Context) {
	//a.AdminID, a.AdminAccount = handler.GetAdmin(ctx)
	a.IP = ctx.ClientIP()
}

func (a *AdminLog) Target(ID int, name string) {
	a.TargetID = ID
	a.TargetName = name
}

func (a *AdminLog) Send(c *gin.Context) {
	//url := config.ServiceConf.Log + "/admin"
	a.Result = 1
	//bodyStr, _ := utils.StructToJsonStr(&a)
	//ReqWithToken("POST", url, c.GetHeader("Authorization"), bytes.NewBuffer([]byte(bodyStr)))
}

func (a *AdminLog) SetOperation(operation string) *AdminLog {
	a.Operation = operation
	return a
}

func (a *AdminLog) SetFunction(function string) *AdminLog {
	a.Function = function
	return a
}

func (a *AdminLog) SetAlertLevel(alertLevel int) *AdminLog {
	a.AlertLevel = alertLevel
	return a
}

func (a *AdminLog) SetIP(ip string) *AdminLog {
	a.IP = ip
	return a
}

func (a *AdminLog) SetUpdateValue(updateValue interface{}) *AdminLog {
	a.UpdateValue = updateValue
	return a
}
