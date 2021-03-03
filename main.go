package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"apicar/db"
	"apicar/model"
)

func main() {
	r := gin.Default()
	r.POST("/users", CreateUser)
	r.GET("/users/:iduser", GetUser)
	r.PUT("/users/:iduser", UpdateUser)
	r.DELETE("/users/:iduser", DeleteUser)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func GetUser(ctx *gin.Context) {
	u, err := db.GetUser(ctx)
	if err != nil {
		ctx.JSON(err.(model.ErrAPI).Code, gin.H{
			"error": err.(model.ErrAPI).Message,
		})
	}
	ctx.JSON(http.StatusOK, u)
}

func CreateUser(ctx *gin.Context) {
	var u model.User
	err := ctx.BindJSON(&u)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	u.ID = uuid.New().String()
	listUser[u.ID] = &u
	ctx.JSON(200, u)
}

func UpdateUser(ctx *gin.Context) {

	var u2 User
	err := ctx.BindJSON(&u2)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	u, err := dbGetUser(ctx)
	if err != nil {
		ctx.JSON(err.(ErrAPI).Code, gin.H{
			"error": err.(ErrAPI).Message,
		})
	}

	u.FirstName = u2.FirstName
	u.LastName = u2.LastName
	u.Email = u2.Email
	u.Password = u2.Password

	ctx.JSON(200, u)
}

func DeleteUser(ctx *gin.Context) {
	_, err := dbGetUser(ctx)
	if err != nil {
		ctx.JSON(err.(ErrAPI).Code, gin.H{
			"error": err.(ErrAPI).Message,
		})
	}

	id := ctx.Param("iduser")
	delete(listUser, id)

	ctx.JSON(200, gin.H{
		"delete": id,
	})
}
