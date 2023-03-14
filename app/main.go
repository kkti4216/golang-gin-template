package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"

	"github.com/gin-gonic/gin"
)


func main() {
	f, _ := os.OpenFile("/var/log/app/main.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	gin.DisableConsoleColor()
	gin.DefaultWriter = io.MultiWriter(os.Stdout, f)

    engine:= gin.Default()
	engine.GET("/",root)

    engine.Run(":3000")
}

func root(c *gin.Context) {
	output, err := exec.Command("/app/script.sh").Output()
	if err != nil {
		log.Println("fail to exec script.sh", err)
	}

	s := string(output)
	log.Println(s)

	c.String(http.StatusOK, s)
}
