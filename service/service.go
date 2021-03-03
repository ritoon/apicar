package service

import "github.com/gin-gonic/gin"

func Init(r *gin.Engine) {
	r.POST("/users", CreateUser)
	r.GET("/users/:iduser", GetUser)
	r.PUT("/users/:iduser", UpdateUser)
	r.DELETE("/users/:iduser", DeleteUser)
}
