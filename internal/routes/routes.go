package routes

import (
	"gorm/internal/handler"

	"github.com/gin-gonic/gin"
)

func Routes(r *gin.Engine) {

	r.POST("/contacts/post",handler.CreateContact)
	r.GET("/contacts/get",handler.GetContacts)
	r.GET("/contacts/get/:id",handler.GetContactsByID)
	r.GET("/contacts/get/name/:name",handler.GetContactsByName)
	r.PUT("/contacts/put/:id",handler.PutContact)
	r.PATCH("/contacts/patch/:id",handler.PatchContact)
	r.DELETE("/contacts/delete/:id",handler.DeleteContact)
}