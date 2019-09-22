# TestGolang-Gin



1.Create Folder TestGolang-Gin

2.Create Folder src in TestGolang-Gin

3.Create file main.go

4.go mod init example.com/m

5.write

==========================

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


==========================

6.go run src/main.go

7.http://localhost:8080/

##SunvoDz##
