package main

import (
	"github.com/gin-gonic/gin"
	"github.com/zhangyiming748/log"
)

type req struct {
	Apikey string `json:"apikey"`
	Ask    string `json:"ask"`
}

var ()

func setupRouter() *gin.Engine {
	r := gin.Default()
	// Get stop
	r.GET("/stop", func(c *gin.Context) {

	})
	r.POST("/chat", func(c *gin.Context) {
		var json req

		if err := c.BindJSON(&json); err != nil {
			return
		}
		log.Info.Printf("%v", &json)

		//in <- json.Ask
		//log.Debug.Println("pass")
		//c.JSON(http.StatusOK, gin.H{
		//	"apikey": json.Apikey,
		//	"ask":    json.Ask,
		//	"ans":    <-out,
		//})
	})

	return r
}

func main() {
	in := make(chan string)
	out := make(chan string, 1)
	key := "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"
	ChatGPT(key, in, out)
	r := setupRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}

// ToDo 单独接口设置key 使用循环阻塞拿到相应