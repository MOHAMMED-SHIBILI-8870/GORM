package routes

import (
	"gorm/internal/handler"

	"github.com/gin-gonic/gin"
)

func Routes(r *gin.Engine) {
	r.POST("/contacts",handler.CreateContact)
	r.GET("/contacts",handler.GetContacts)
	r.GET("/contacts/:id",handler.GetContactsByID)
	r.PATCH("contacts/:id",handler.UpdateContact)
	r.DELETE("/contacts/:id",handler.DeleteContact)
}