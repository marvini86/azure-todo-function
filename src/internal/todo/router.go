package todo

import "github.com/gin-gonic/gin"

func InitRouter(r *gin.RouterGroup) {
	r.GET("/todos", GetAll)
	r.GET("/todos/:id", GetById)
}
