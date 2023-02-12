package main

import (
	"github.com/gin-gonic/gin"
	"github.com/zhangyiming748/log"
	"net/http"
)

type req struct {
	Apikey string `json:"apikey"`
	Ask    string `json:"ask"`
}

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.POST("/chat", func(c *gin.Context) {
		var json req
		if err := c.BindJSON(&json); err != nil {
			return
		}
		log.Info.Printf("%v", &json)
		ans := ChatGPT(json.Apikey, json.Ask)
		c.JSON(http.StatusOK, gin.H{
			"apikey": json.Apikey,
			"ask":    json.Ask,
			"ans":    ans,
		})
	})
	return r
}

func main() {
	r := setupRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
