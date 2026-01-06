package handler

import (
	"gorm/internal/config"
	"gorm/internal/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateContact(c *gin.Context) {
	var Contact models.Contact

	if err:=c.ShouldBindJSON(&Contact);err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{
			"error":err.Error(),
		})
		return
	}

	if err:=config.DB.Create(&Contact).Error;err!=nil{
		c.JSON(http.StatusInternalServerError,gin.H{
			"error":err.Error(),
		})
		return
	}

	c.JSON(200,gin.H{
		"data":Contact,
		"msg":"Data's are created successflly",
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
	if err!=nil{
		c.JSON(400,gin.H{
			"error":"Invalid id",
		})
		return
	}

	var Contact models.Contact

	if err:=config.DB.First(&Contact,id).Error;err!=nil{
		c.JSON(404,gin.H{
			"error":"User not found",
		})
		return
	}

	c.JSON(200,Contact)
}

func GetContactsByName(c *gin.Context){
	name := c.Param("name")
	
	var contact models.Contact

	if err:=config.DB.Where("name = ?",name).First(&contact).Error;
	err != nil{
		c.JSON(404,gin.H{
			"error":"User not found",
		})
		return
	}

	c.JSON(200,contact)
}

func PutContact(c *gin.Context){
	id:=c.Param("id")

	var contact models.Contact

	if err:=config.DB.First(&contact,id).Error;err!=nil{
		c.JSON(http.StatusNotFound,gin.H{
			"error":"Record not found",
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

	if err:= config.DB.Model(&contact).Updates(models.Contact{
		Name: input.Name,
		Email: input.Email,
		Phone: input.Phone,
	}).Error;err!=nil{
		c.JSON(http.StatusInternalServerError,gin.H{
			"error":err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK,gin.H{
		"msg":"Updated Successfully",
		"data":contact,
	})
}

func PatchContact(c *gin.Context){
	id:= c.Param("id")

	var contact models.Contact

	if err:=config.DB.First(&contact,id).Error;err !=nil{
		c.JSON(http.StatusNotFound,gin.H{
			"error":"Record not Found",
		})
		return
	}

	var input models.PatchContactInput
	if err:=c.ShouldBindJSON(&input);err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{
			"error":err.Error(),
		})
		return
	}
	if err:= config.DB.Model(&contact).Updates(input).Error;err!=nil{
		c.JSON(http.StatusInternalServerError,gin.H{
			"error":err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK,gin.H{
		"msg":"Updated successfully",
		"data":contact,
	})
}

func DeleteContact(c *gin.Context){
	id:=c.Param("id")

	var  contact models.Contact

	result:=config.DB.Delete(&contact,id)

	if result.Error != nil{
		c.JSON(http.StatusInternalServerError,gin.H{
			"error":result.Error.Error(),
		})
		return
	}

	if result.RowsAffected == 0{
		c.JSON(http.StatusNotFound,gin.H{
			"message":"contact not found",
		})
		return
	}

	c.JSON(200,gin.H{
		"msg":"Deleted successfully",
	})
}