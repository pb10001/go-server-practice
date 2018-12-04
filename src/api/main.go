package main

import (
	"log"

	"github.com/garyburd/redigo/redis"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	redi, err := redis.Dial("tcp", "redis:6379")
	if err != nil {
		log.Fatal(err)
	}
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, World!",
		})
	})
	r.GET("/counter", func(c *gin.Context) {
		res, err := redi.Do("incr", "counter")
		if err != nil {
			c.JSON(403, gin.H{
				"error": err.Error(),
			})
		} else {
			c.JSON(200, gin.H{
				"counter": res,
			})
		}
	})
	r.GET("/ping", func(c *gin.Context) {
		var m = Message{Text: "gwefwahrher"}
		c.JSON(200, m)
	})
	r.Run(":5000")
}
