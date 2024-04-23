package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"labs.com/todofn/src/internal/todo"
)

var port string

func main() {
	if portValue, ok := os.LookupEnv("FUNCTIONS_CUSTOMHANDLER_PORT"); ok {
		port = fmt.Sprintf(":%v", portValue)
	}
	r := gin.Default()

	api := r.Group("/api")
	{
		todo.InitRouter(api)
	}

	r.Run(port)

}
