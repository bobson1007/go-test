package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

/**
api內不允許有Logger相關的code
*/
func DoNothing(c *gin.Context) {
	fmt.Println("進入Api")
	c.Data(200, "text/plain", []byte("success"))
}
