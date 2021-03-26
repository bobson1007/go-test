package middleware

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
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

	apiMap := enum.GetApiMap()
	api, ok := apiMap[c.FullPath()]
	if ok {
		var log model.AdminLog
		log.SetOperation(api.Operation).SetFunction(api.Function).SetAlertLevel(api.AlertLevel)
		c.Set("logModel", log)
	}
}

func adminLogSender(c *gin.Context) {
	fmt.Println("進入LogSender")

	log, ok := c.Get("logModel")
	adminLog := log.(model.AdminLog)
	printLog(c)

	// TODO 處理request

	if ok {
		// 發送日誌給log-service
		fmt.Println("發送adminLog : ", adminLog.Operation, adminLog.Function, adminLog.AlertLevel)

	} else {
		fmt.Println("logEnum沒有此路徑，故不需發送日誌")
	}
}

func printLog(c *gin.Context) {
	data, err := c.GetRawData()
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Printf("data: %v\n", string(data))

	c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(data)) // 关键点

	accessLogMap := make(map[string]interface{})
	//accessLogMap["request_time"]      = startTime
	accessLogMap["request_method"] = c.Request.Method
	accessLogMap["request_uri"] = c.Request.RequestURI
	accessLogMap["request_proto"] = c.Request.Proto
	accessLogMap["request_ua"] = c.Request.UserAgent()
	accessLogMap["request_referer"] = c.Request.Referer()
	accessLogMap["request_post_data"] = c.Request.PostForm.Encode()
	accessLogMap["request_client_ip"] = c.ClientIP()
	accessLogMap["request_body"] = c.Request.Body

	fmt.Println("========>", accessLogMap)
}
