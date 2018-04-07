package rest

import (
	"github.com/gin-gonic/gin"
	"bitcamp/api/service"
	"bitcamp/common/models/dto"
	"encoding/json"
	"log"
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

