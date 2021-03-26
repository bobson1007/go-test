package middleware

import "github.com/gin-gonic/gin"

func HandleError() gin.HandlerFunc {
	return func(context *gin.Context) {
		// TODO 在已有的HandleError上加入處理日誌的步驟
	}
}
