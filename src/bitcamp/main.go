package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
	"time"
)

func main(){
	r := gin.Default()

	var domain string
	var origins []string
	if gin.IsDebugging(){
		domain = "192.168.0.111:8080"
		origins = []string{"http://localhost:4200", "https://localhost", "http://192.168.0.111:4200"}
	}else{
		domain = ":8080"
		origins = []string{"https://eip.umbc.edu"}
	}
	r.Use(cors.New(cors.Config{
		AllowOrigins:     origins,
		AllowMethods:     []string{"PUT", "PATCH", "DELETE", "GET", "POST"},
		AllowHeaders:     []string{"Origin", "Authorization", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	r.Run(domain)
}