package main

import (
  "fmt"
  "github.com/gin-gonic/gin"
  "github.com/garyburd/redigo/redis"
)

func main() {
  // redis port
  redisPort, err := redis.Dial("tcp", ":4000")

	if err != nil {
		panic(err)
	}

	defer redisPort.Close()

	// notification api
	router := gin.Default()

	router.GET("/notification", func(c *gin.Context) {
    message := c.Query("message")
    if message != "" {
      value := fmt.Sprintf("{\"class\":\"notifier\",\"args\":[\"%s\"]}", message)
      redisPort.Do("RPUSH", "resque:queue:slack",  value)
		  c.JSON(200, gin.H{
			  "message": 200,
		  })
    } else {
		  c.JSON(403, gin.H{
			  "message": "Please send message",
		  })
    }
	})

	router.Run() // listen and serve on 0.0.0.0:8080
}
