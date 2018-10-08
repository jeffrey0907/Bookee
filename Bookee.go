package main

import (
	"Bookee/controller"
	"github.com/gin-gonic/gin"
	"io"
	"os"
)

func main() {
	gin.DisableConsoleColor()
	// Logging to a file.
	logFile, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(logFile, gin.DefaultWriter)

	root := gin.Default()

	userRouter := root.Group(`user`)
	controller.RegistUserController(userRouter)

	bookRouter := root.Group(`book`)
	controller.RegisterBookController(bookRouter)

	root.Run(`:8080`)
}
