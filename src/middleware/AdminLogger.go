package middleware

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"strings"
	"test/src/enum"
	"test/src/model"
)

func AdminLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		adminLogInit(c)
		c.Next()
		adminLogSender(c)
	}
}

/**
處理enum內的初始值，與adminID和ip，requset的處理
*/
func adminLogInit(c *gin.Context) {
	fmt.Println("進入LogFilter")

	api, ok := getApiByPathAndRequestMethod(c)
	if ok {
		var log model.AdminLog
		log.SetOperation(api.Operation).SetFunction(api.Function).SetAlertLevel(api.AlertLevel).SetIP(c.ClientIP())
		// TODO 調用已有函式紀錄adminID
		handleRequestParams(c, &log)
		c.Set("logModel", log)
	} else {
		fmt.Println("enum無該api，不做初始化")
	}
}

func adminLogSender(c *gin.Context) {
	fmt.Println("進入LogSender")
	l, ok := c.Get("logModel")
	if ok {
		// TODO 處理HandleError的結果
		// TODO 調用已有函式發送log給log-service
		log := l.(model.AdminLog)
		fmt.Println("發送adminLog : ", log)
	} else {
		fmt.Println("logEnum沒有此路徑，故不需發送日誌")
	}
}

func handleRequestParams(c *gin.Context, log *model.AdminLog) {
	// TODO 處理Request中的OriginalValue
	updateValueMap := make(map[string]interface{}, 0)
	getParamsInPath(c, updateValueMap)
	getParamsInQuerystring(c, updateValueMap)
	getParamsInJson(c, updateValueMap)
	getParamsInForm(c, updateValueMap)
	log.SetUpdateValue(updateValueMap)
}

func getParamsInPath(c *gin.Context, updateValueMap map[string]interface{}) {
	if len(c.Params) > 0 {
		for _, e := range c.Params {
			updateValueMap[e.Key] = e.Value
		}
	}
}

func getParamsInQuerystring(c *gin.Context, updateValueMap map[string]interface{}) {
	if len(c.Request.URL.Query()) > 0 {
		for k, v := range c.Request.URL.Query() {
			if len(v) == 1 {
				updateValueMap[k] = v[0]
			}
			updateValueMap[k] = v
		}
	}
}
func getParamsInJson(c *gin.Context, updateValueMap map[string]interface{}) {
	data, err := c.GetRawData()
	if err != nil {
		fmt.Println(err.Error())
	}
	c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(data))

	if len(data) > 0 {
		var f interface{}
		_ = json.Unmarshal(data, &f)
		m, err := f.(map[string]interface{})
		if !err {
			for k, v := range m {
				updateValueMap[k] = v
			}
		}
	}
}

func getParamsInForm(c *gin.Context, updateValueMap map[string]interface{}) {
	if c.Request.Form == nil {
		_ = c.Request.ParseMultipartForm(32 << 20)
	}
	if len(c.Request.Form) > 0 {
		for k, v := range c.Request.Form {
			updateValueMap[k] = v
		}
	}
}

func getApiByPathAndRequestMethod(c *gin.Context) (enum.Api, bool) {
	fullPath := c.FullPath()
	for _, p := range c.Params {
		fullPath = strings.Replace(fullPath, p.Value, ":"+p.Key, 1)
	}
	requestMethod := c.Request.Method
	key := requestMethod + "_" + fullPath
	apiMap := enum.GetApiMap()
	api, ok := apiMap[key]
	return api, ok
}
