package main

import (
	"github.com/gin-gonic/gin"
	"test/src/api"
	. "test/src/enum"
	"test/src/middleware"
)

func main() {
	router := gin.Default()

	// 處理管理者日誌
	router.Use(middleware.AdminLogger())
	// TODO　原本的HandleError需要加入error日誌的處理
	router.Use(middleware.HandleError())

	v1 := router.Group(GroupApiV1)

	v1.GET(GetAgentPolicyPath, api.DoNothing)
	v1.POST(UploadPolicyPath, api.DoNothing)

	v1.GET(GetAllPolicyPath, api.DoNothing)
	v1.POST(CreatePolicyPath, api.DoNothing)
	v1.PUT(UpdatePolicyContentPath, api.DoNothing)
	v1.GET(GetOnePolicyPath, api.DoNothing)
	v1.PATCH(UpdatePolicyNamePath, api.DoNothing)
	v1.DELETE(DeletePolicyPath, api.DoNothing)

	v1.GET(GetPublicKeyPath, api.DoNothing)
	v1.GET(GetLicenseEnablePath, api.DoNothing)
	v1.POST(UploadLicensePath, api.DoNothing)
	v1.GET(GetLicensePath, api.DoNothing)

	v1.GET(DownloadPolicyTemplatePath, api.DoNothing)

	_ = router.Run(":8081")
}
