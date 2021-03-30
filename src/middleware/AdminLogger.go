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

func adminLogInit(c *gin.Context) {
	fmt.Println("進入LogFilter")

	api, ok := getApiByPath(c)
	if ok {
		var log model.AdminLog
		log.SetOperation(api.Operation).SetFunction(api.Function).SetAlertLevel(api.AlertLevel)
		c.Set("logModel", log)
	} else {
		fmt.Println("enum無該api，不做初始化")
	}
}

func adminLogSender(c *gin.Context) {
	fmt.Println("進入LogSender")
	l, ok := c.Get("logModel")
	if ok {
		log := l.(model.AdminLog)
		// TODO 處理request
		printLog(c)
		fmt.Println("發送adminLog : ", log.Operation, log.Function, log.AlertLevel)
		// TODO 發送日誌給log-service
	} else {
		fmt.Println("logEnum沒有此路徑，故不需發送日誌")
	}
}

func printLog(c *gin.Context) {

	//accessLogMap := make(map[string]interface{})
	////accessLogMap["request_time"]      = startTime
	//accessLogMap["request_method"] = c.Request.Method
	//accessLogMap["request_uri"] = c.Request.RequestURI
	//accessLogMap["request_proto"] = c.Request.Proto
	//accessLogMap["request_ua"] = c.Request.UserAgent()
	//accessLogMap["request_referer"] = c.Request.Referer()
	//accessLogMap["request_post_data"] = c.Request.PostForm.Encode()
	//accessLogMap["request_client_ip"] = c.ClientIP()
	//accessLogMap["request_body"] = c.Request.Body

	updateValueMap := make(map[string]interface{}, 0)

	// TODO 需補上進入的判斷

	fmt.Println("========param in path========>")
	fmt.Println(c.Params) // Param陣列，Param包含k和v
	for _, e := range c.Params {
		//fmt.Println(k, v)
		updateValueMap[e.Key] = e.Value
	}

	fmt.Println("========GET param not in path========>")
	for k, v := range c.Request.URL.Query() {
		if len(v) == 1 {
			updateValueMap[k] = v[0]
		}
		updateValueMap[k] = v
	}
	fmt.Println(updateValueMap)

	fmt.Println("========param in JSON========>")
	// request body只能讀取一次的解決辦法
	data, err := c.GetRawData()
	if err != nil {
		fmt.Println(err.Error())
	}
	c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(data)) // 关键点

	if len(data) > 0 {
		var f interface{}
		_ = json.Unmarshal(data, &f)
		m := f.(map[string]interface{})
		for k, v := range m {
			//fmt.Println(k, v)
			updateValueMap[k] = v
		}
	}

	fmt.Println("======== param in x-www-form-urlencode / form-data ========>")
	if c.Request.Form == nil {
		_ = c.Request.ParseMultipartForm(32 << 20)
	}
	for k, v := range c.Request.Form {
		//fmt.Println(k, v)
		updateValueMap[k] = v
	}

}

func getApiByPath(c *gin.Context) (enum.Api, bool) {
	apiMap := enum.GetApiMap()
	fullPath := c.FullPath()
	for _, p := range c.Params {
		fullPath = strings.Replace(fullPath, p.Value, ":"+p.Key, 1)
	}
	api, ok := apiMap[fullPath]
	return api, ok
}
