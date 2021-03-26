package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
)

func CreatePolicy(c *gin.Context) {
	fmt.Println("進入Api")
	c.Data(200, "text/plain", []byte("success"))
	body, _ := ioutil.ReadAll(c.Request.Body)
	fmt.Println(body)
	form, _ := c.GetPostForm("age")
	fmt.Println(form)
	// TODO api的處理結果放進Context給中間件LogSender處理

}
