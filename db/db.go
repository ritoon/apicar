package db

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"apicar/model"
)

var listUser map[string]*model.User = make(map[string]*model.User)

func GetUser(ctx *gin.Context) (*model.User, error) {
	id := ctx.Param("iduser")
	log.Println("get uuid", id)
	if len(id) == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "no given ID",
		})
		return nil, model.ErrAPIBadRequest
	}

	if u, ok := listUser[id]; ok {
		return u, nil
	}

	return nil, model.ErrAPINotfound
}

func CreateUser(u *model.User) (*model.User, error) {

}
