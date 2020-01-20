package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jasonlvhit/gocron"
)

func rootHandler(context *gin.Context) {
	context.JSON(200, gin.H{
		"hello": "world",
	})
}

func scheduledTask() {
	fmt.Println("Are you still there?")
}

func main() {
	go setupScheduler()
	setupHttpRouter()
}

func setupHttpRouter(){
	r:= gin.Default()
	r.GET("/", rootHandler)

	r.Run(":80")
}

func setupScheduler() {
	gocron.Every(3).Seconds().Do(scheduledTask)
	<-gocron.Start()
}
