package service

import (
	"log"
	"bitcamp/common/rabbit_helper"
	"bitcamp/common/models/dto"
	"encoding/json"
	"bitcamp/common/queries"
	"net/http"
	"bytes"
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

		err = BroadcastMessage(textDTO.UserId, textDTO.Content, textDTO.OtherNumber, textDTO.Sent)
		if err != nil{
			log.Println(err)
		}

	}
}

func callPredictorAPI(to dto.TextDTO) (float32, error){
	baseUrl := "http://ann:4200/rest/classify"

	d := dto.NNReq{Content:to.Content}

	jsonVal, _ := json.Marshal(d)

	resp, err := http.Post(baseUrl, "application/json", bytes.NewBuffer(jsonVal))
	if err != nil{
		log.Println(err)
		return -1, err
	}

	var response dto.NNResp
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil{
		log.Println(err)
		return -1, err
	}

	if response.Insult > response.NotInsult{
		return response.Insult, nil
	}else{
		return 0, nil
	}
}
