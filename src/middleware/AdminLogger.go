package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
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

	// TODO 處理api內給的信息
	// TODO 處理HandleError的信息

	if ok {
		// 發送日誌給log-service
		fmt.Println("發送adminLog : ", adminLog.Operation, adminLog.Function, adminLog.AlertLevel)
	} else {
		fmt.Println("logEnum沒有此路徑，故不需發送日誌")
	}
}
