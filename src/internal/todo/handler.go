package todo

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"labs.com/todofn/src/internal/common"
)

func GetAll(c *gin.Context) {
	todos, _ := GetAllTodos()
	c.JSON(http.StatusOK, todos)
}

func GetById(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	todo, _ := GetTodoById(int64(id))

	if todo.Id == 0 {
		c.JSON(http.StatusNotFound, common.Response{Code: http.StatusNotFound, Message: "Todo not found"})
		return
	}

	c.JSON(http.StatusOK, todo)
}
