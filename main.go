package main

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"io"
	"simple-api-mock/application/config"
	"strconv"
)

func main() {
	router := gin.Default()
	conf := config.Load()
	logger, err := config.NewLogger(conf.Config.LogPath)
	if err != nil {
		panic(err)
	}

	router.Any("/*path", func(c *gin.Context) {
		logger.Info("path:" + c.Request.RequestURI + "\tmethod:" + c.Request.Method + "\tbody:" + streamToString(c.Request.Body))
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "*")
		res, err := conf.Contain(c.Request.RequestURI, c.Request.Method)
		if err != nil {
			c.String(404, "api not found")
		}
		c.String(res.Code, res.Body)
	})

	router.Run(":" + strconv.Itoa(conf.Config.Port))
}

func streamToString(stream io.Reader) string {
	buf := new(bytes.Buffer)
	buf.ReadFrom(stream)
	return buf.String()
}
