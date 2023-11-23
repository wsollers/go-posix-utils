package web

import (
  "fmt"
  "log"
  "net/http"
  "strconv"
  "github.com/gin-gonic/gin"
  "github.com/spf13/cobra"
)

var host string
var port uint64

var WebServCmd = &cobra.Command{
  Use:   "webServ",
  Short: "web server",
  Long:  "web server",
  Run: func(cmd *cobra.Command, args []string) {
    fmt.Println("WebServCmd called")

    if len(args) < 1 {
      log.Fatal("usage: webServ ip port")
    }
    host = args[0]
    if len(host) < 1 {
      log.Fatal("No host provided")
    }
    err := error(nil)
    port, err = strconv.ParseUint(args[1], 10, 16)
    if err != nil {
      log.Fatal("Invalid port")
    }
    fmt.Println("host:", host)
    if host == "" {
      log.Fatal("No host provided")
    }
  
    serveStuff()
  },
}

type Param struct {
  Param1 string `form:"param1" json:"param1" binding:"required"`
  Param2 string `form:"param2" json:"param2" binding:"required"`
}

func serveStuff() {
  fmt.Println("Serving stuff")
  r := gin.Default()
  r.GET("/ping", func(c *gin.Context) {
    log.Println("Got a ping")
    c.JSON(http.StatusOK, gin.H{
      "message": "pong",
    })
  })

  r.GET("/ping/request-params", func(c *gin.Context) {

    var params Param

    if err := c.Bind(&params); err != nil {
      log.Fatal("ShouldBind: ",err)
      log.Println("Got a ping with bad params")
      log.Println(err)
      c.JSON(http.StatusBadRequest, gin.H{
        "error": err.Error(),
      })
    }
    log.Println("Got a ping with param1", params.Param1, "and param2", params.Param2)
    c.JSON(http.StatusOK, gin.H{
      "message": "pong",
      "param1": params.Param1,
      "param2": params.Param2,
    })
  })

  r.GET("/ping/:id", func(c *gin.Context) {
    log.Println("Got a ping with id:", c.Param("id"))
    c.JSON(http.StatusOK, gin.H{
     "message": "pong",
     "id": c.Param("id"),
    })
  })

  //default bad stuff handler
  r.NoRoute(func(c *gin.Context) {
    c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
  })
  r.NoMethod(func(c *gin.Context) {
    c.JSON(404, gin.H{"code": "METHOD_NOT_ALLOWED", "message": "METHOD_NOT_ALLOWED"})
  })
  
  r.Run() // listen and serve on
}
