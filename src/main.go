package main

import (
  "fmt"
  "github.com/gin-gonic/gin"
)
func main() {
  fmt.Println("Hello, Sunvo")

  r := gin.Default()
  r.GET("/", func(c *gin.Context){
    c.JSON(200, gin.H{
      "massage":"SunvoDz Wellcome",
    })
  })
  r.Run()
}

