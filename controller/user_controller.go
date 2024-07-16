package controller

import (
	"github.com/gin-gonic/gin"
	"go-web/service"
	"go-web/util"
	"net/http"
)

func GetUsers(c *gin.Context) {
	users, err := service.GetAllUsers()
	if err != nil {
		util.RespondJSON(c, http.StatusInternalServerError, err.Error())
		return
	}

	util.RespondJSON(c, http.StatusOK, users)
}

func GetUserByID(c *gin.Context) {
	id := c.Param("id")
	user, err := service.GetUserByID(id)
	if err != nil {
		util.RespondJSON(c, http.StatusInternalServerError, err.Error())
		return
	}
	util.RespondJSON(c, http.StatusOK, user)
}

func CreateUser(c *gin.Context) {
	var user service.User
	if err := c.ShouldBind(&user); err != nil {
		util.RespondJSON(c, http.StatusBadRequest, err.Error())
		return
	}
	err := service.CreateUser(&user)
	if err != nil {
		util.RespondJSON(c, http.StatusInternalServerError, err.Error())
		return
	}
	util.RespondJSON(c, http.StatusCreated, user)
}
