package controller 

import (
	"database/sql"
	"net/http"
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"

	"server/model"
)

//UserCtr is a structure for using databases
type UserCtr struct {
	DB *sql.DB
}


//HelloWorld is a function to return Hello world
func (u* UserCtr) HelloWorld(c *gin.Context)  {
	c.JSON(http.StatusOK,gin.H{
		"message": "Hello World!!",
	})
}
// GetUserAll is a function to list up all users
func (u *UserCtr) GetUserAll(c *gin.Context) {
	users, err := model.UserAll(u.DB)
	if err != nil {
		resp := errors.New(err.Error())
		c.JSON(http.StatusInternalServerError, resp)
		return
	}

	if len(users) == 0 {
		c.JSON(http.StatusOK, make([]*model.User, 0))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"result": users,
	})
	return
}

//GetUserByid is a funciton to list up user with id
func (u *UserCtr) GetUserByid(c *gin.Context) {
	id,err := strconv.Atoi(c.Param("id"))
	if err != nil {
		resp := errors.New(err.Error())
		c.JSON(http.StatusInternalServerError, resp)
		return
	}

	user, err := model.UserByID(u.DB,id)
	if err != nil {
		resp := errors.New(err.Error())
		c.JSON(http.StatusInternalServerError, resp)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"result": user,
	})
	return

}

//InsertUser is a function to insert new user
func (u *UserCtr) InsertUser(c *gin.Context) {
	var newuser model.User
	c.BindJSON(&newuser)
	user, err := model.InsertUser(u.DB,newuser.Name,newuser.Email)
	if err != nil {
		resp := errors.New(err.Error())
		c.JSON(http.StatusInternalServerError, resp)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"result": user,
	})
	return
}

//UpdateUser is a function to update user with id 
func (u *UserCtr) UpdateUser(c *gin.Context) {
	id,err := strconv.Atoi(c.Param("id"))
	if err != nil {
		resp := errors.New(err.Error())
		c.JSON(http.StatusInternalServerError, resp)
		return
	}

	var newuser model.User
	c.BindJSON(&newuser)
	user, err := model.UpdateUser(u.DB,id,newuser.Name,newuser.Email)
	if err != nil {
		resp := errors.New(err.Error())
		c.JSON(http.StatusInternalServerError, resp)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"result": user,
	})
	return
}

//DeleteUser is a function to delete user with id
func (u *UserCtr) DeleteUser(c *gin.Context) {
	id,err := strconv.Atoi(c.Param("id"))
	if err != nil {
		resp := errors.New(err.Error())
		c.JSON(http.StatusInternalServerError, resp)
		return
	}

	err = model.DeleteUser(u.DB,id)
	if err != nil {
		resp := errors.New(err.Error())
		c.JSON(http.StatusInternalServerError, resp)
		return
	}

	c.JSON(http.StatusOK, gin.H{
	})
	return
}