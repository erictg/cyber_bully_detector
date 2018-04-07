package service

import (
	"log"
)

func textRoutine(s *SendBroker){
	for{
		select {
		case dto := <- s.procChan:
			log.Println("in the chan")
			if dto.Quit{
				return
			}
			err := s.SubmitToQueue(dto)
			if err != nil{
				log.Println(err)
			}
			break
		}
	}
}
