package rest

import (
	"github.com/gin-gonic/gin"
	"bitcamp/api/service"
	"bitcamp/common/models/dto"
	"encoding/json"
	"log"
	"strconv"
	"bitcamp/common/queries"
	"bitcamp/common/models"
)

func TextPostWrap(broker *service.SendBroker) func(ctx *gin.Context){
	return func(c *gin.Context) {
		var textDto dto.TextDTO
		err := json.NewDecoder(c.Request.Body).Decode(&textDto)
		if err != nil{
			log.Println(err)
			c.AbortWithError(400, err)
			return
		}
		broker.PushText(textDto)
		c.Status(200)
	}
}

func GetTextById(c *gin.Context){
	id := c.Param("id")

	i, err := strconv.Atoi(id)
	if err != nil{
		log.Println(err)
		c.AbortWithError(400, err)
		return
	}

	user, err := queries.GetTextById(i)
	if err != nil{
		log.Println(err)
		c.AbortWithError(400, err)
		return
	}

	c.JSON(200, user)
}

func GetAllTextsForUser(c *gin.Context){
	id := c.Param("id")

	i, err := strconv.Atoi(id)
	if err != nil{
		log.Println(err)
		c.AbortWithError(400, err)
		return
	}

	user, err := queries.GetTextsFromUser(i)
	if err != nil{
		log.Println(err)
		c.AbortWithError(400, err)
		return
	}



	c.JSON(200, TextContainer{Texts:user})
}

type TextContainer struct{
	Texts []models.FlaggedText `json:"texts"`
}