package handler

import (
	"gorm/internal/config"
	"gorm/internal/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateContact(c *gin.Context){
	var Contact models.Contact

	if err:=c.ShouldBindJSON(&Contact);err != nil{
		c.JSON(http.StatusBadRequest,gin.H{
			"error":err.Error(),
		})
		return
	}
	if err :=config.DB.Create(&Contact).Error;err!=nil{
		c.JSON(http.StatusInternalServerError,gin.H{
			"error":err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated,gin.H{
		"msg":"contact created successfully",
		"contacs":Contact,
	})
}

func GetContacts(c *gin.Context){
	var contacts []models.Contact

	config.DB.Find(&contacts)
	c.JSON(200,contacts)
}

func GetContactsByID(c *gin.Context){
	idparams:=c.Param("id")
	id,err:=strconv.Atoi(idparams)
	if err != nil{
		c.JSON(http.StatusBadRequest,gin.H{
			"error":"Invalid ID",
		})
		return
	}
	
	var contacts models.Contact
	if err:=config.DB.First(&contacts,id).Error;err != nil{
		c.JSON(404,gin.H{
			"error":"user not found",
		})
		return
	}
	c.JSON(200,contacts)
}

func GetContactsByName(c *gin.Context){
	name:=c.Param("name")

	var contacts []models.Contact

	result:=config.DB.Where("name = ?",name).Find(&contacts)

	if result.RowsAffected == 0{
		c.JSON(http.StatusNotFound,gin.H{
			"error":"no users found",
		})
		return
	}

	c.JSON(200,contacts)

}

func UpdateContact(c *gin.Context){
	id := c.Param("id")

	var contact models.Contact

	if err:=config.DB.First(&contact,id).Error;err != nil{
		c.JSON(404,gin.H{
			"error":"record not found",
		})
		return
	}

	var input models.Contact

	if err:=c.ShouldBindJSON(&input);err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{
			"error":err.Error(),
		})
		return
	}

	contact.Name=input.Name
	contact.Email=input.Email
	contact.Phone=input.Phone

	if err:=config.DB.Save(&contact).Error;err!=nil{
		c.JSON(http.StatusInternalServerError,gin.H{
			"error":err.Error(),
		})
		return
	}

	c.JSON(200,gin.H{
		"data":contact,
	})
}

func DeleteContact(c *gin.Context){
	id := c.Param("id")

	var contact models.Contact

	if err:=config.DB.First(&contact,id).Error;err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{
			"msg":"Record not found",
		})
		return
	}
	config.DB.Delete(&contact)
	c.JSON(200,gin.H{
		"msg":"Deleted successfuly",
	})
}