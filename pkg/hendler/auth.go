package hendler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zhashkevych/todo-app"
)

func (h *Hendler) singUp(c *gin.Context) {
	var input todo.User

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.service.Authorization.CreateUser(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})

}

func (h *Hendler) singIn(c *gin.Context) {

}
