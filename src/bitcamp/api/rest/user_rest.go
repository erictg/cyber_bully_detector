package rest

import (
	"github.com/gin-gonic/gin"
	"bitcamp/common/models/dto"
	"encoding/json"
	"bitcamp/common/queries"
	"log"
	"strconv"
)

func UserPOST(c *gin.Context){
	var userDto dto.UserCreateDTO
	_ = json.NewDecoder(c.Request.Body).Decode(&userDto)

	err := queries.CreateUser(userDto.Name, userDto.IsParent)
	if err != nil{
		log.Println(err)
		c.AbortWithError(400, err)
		return
	}

	c.Status(200)
}

func UserPairPOST(c *gin.Context){
	var pairDTO dto.PairDTO
	_ = json.NewDecoder(c.Request.Body).Decode(&pairDTO)

	err := queries.PairParentAndChild(pairDTO.PId, pairDTO.CId)
	if err != nil{
		log.Println(err)
		c.AbortWithError(400, err)
		return
	}

	c.Status(200)
}

//param 'name'
func GetUserByName(c *gin.Context){
	name := c.Param("name")

	user, err := queries.GetUserByName(name)
	if err != nil{
		log.Println(err)
		c.AbortWithError(400, err)
		return
	}

	c.JSON(200, user)
}

func GetUserById(c *gin.Context){
	id := c.Param("id")

	i, err := strconv.Atoi(id)
	if err != nil{
		log.Println(err)
		c.AbortWithError(400, err)
		return
	}

	user, err := queries.GetUserById(i)
	if err != nil{
		log.Println(err)
		c.AbortWithError(400, err)
		return
	}

	c.JSON(200, user)
}


