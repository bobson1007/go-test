package main

import (
	"github.com/gin-gonic/gin"
	"test/src/api"
	"test/src/enum"
	"test/src/middleware"
)

func main() {
	router := gin.Default()

	// 處理管理者日誌
	router.Use(middleware.AdminLogger())
	// 原本的HandleError需要加入error日誌的處理
	router.Use(middleware.HandleError())

	router.POST(enum.CreatePolicyPath, api.CreatePolicy)

	_ = router.Run(":8081")
}
