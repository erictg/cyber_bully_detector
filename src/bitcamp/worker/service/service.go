package service

import (
	"log"
	"bitcamp/common/rabbit_helper"
	"bitcamp/common/models/dto"
	"encoding/json"
	"bitcamp/common/queries"
)

func CreateWorkers(workers int) error{
	_, channel, queue, err := rabbit_helper.Setup()
	if err != nil{
		log.Println(err)
		return err
	}

	for i := 0; i < workers; i++{
		msgs, err := channel.Consume(
			queue.Name, // queue
			"",     // consumer
			true,   // auto-ack
			false,  // exclusive
			false,  // no-local
			false,  // no-wait
			nil,    // args
		)

		if err != nil{
			log.Println(err)
			return err
		}

		go func() {
			for d := range msgs {
				log.Printf("Received a message: %s", d.Body)
				var tdto dto.TextDTO
				err := json.Unmarshal(d.Body, &tdto)
				if err != nil{
					log.Println(err)
					continue
				}

				handler(tdto)
			}
		}()
	}

	return nil
}

func handler(textDTO dto.TextDTO){
	log.Println(textDTO.Content)
	result, err := callPredictorAPI(textDTO)
	if err != nil{
		log.Println(err)
		return
	}

	//your child is either being bullied, or is a little fucker
	if result >= .75{
		err := queries.InsertFlaggedText(textDTO.Content, result, textDTO.Sent, textDTO.UserId, textDTO.OtherNumber)
		if err != nil{
			log.Println(err)
		}

	}
}

func callPredictorAPI(dto dto.TextDTO) (float32, error){
	//todo implement
	return .1, nil
}
