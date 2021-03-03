package service

import (
	"apicar/db"
	"apicar/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUser(ctx *gin.Context) {
	u, err := db.GetUser(ctx.Param("iduser"))
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
	db.CreateUser(&u)
	ctx.JSON(200, u)
}

func UpdateUser(ctx *gin.Context) {
	var u2 model.User
	err := ctx.BindJSON(&u2)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	id := ctx.Param("iduser")
	u, _ := db.UpdateUser(id, &u2)

	ctx.JSON(200, u)
}

func DeleteUser(ctx *gin.Context) {
	_, err := db.GetUser(ctx.Param("iduser"))
	if err != nil {
		ctx.JSON(err.(model.ErrAPI).Code, gin.H{
			"error": err.(model.ErrAPI).Message,
		})
	}

	id := ctx.Param("iduser")
	db.DeleteUser(id)

	ctx.JSON(200, gin.H{
		"delete": id,
	})
}
