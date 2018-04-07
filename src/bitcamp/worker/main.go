package main

import (
	"log"
	"bitcamp/worker/service"
)

func main(){
	err := service.CreateWorkers(1)
	if err != nil{
		log.Println(err)
		return
	}
	neverEnd := make(chan bool)
	log.Println("worker")
	<-neverEnd
}
