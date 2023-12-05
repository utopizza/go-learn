package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Conf struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func main() {
	str := `{"key":"k1","value":"https://www.baidu.com?aid=1128&did=1234\u0026uid=5678"}`
	conf := &Conf{}
	if err := json.Unmarshal([]byte(str), conf); err != nil {
		fmt.Printf("err=%v", err)
		return
	}

	// http
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.GET("/decode", func(c *gin.Context) {
		c.JSON(http.StatusOK, conf)
	})
	r.Run()
}
