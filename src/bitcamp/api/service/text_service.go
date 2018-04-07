package service

import (
	"bitcamp/common/models/dto"
	"errors"
	"log"
	"github.com/streadway/amqp"
	"encoding/json"
	"bitcamp/common/rabbit_helper"
)

type SendBroker struct{
	procChan chan dto.TextDTO
	connection *amqp.Connection
	queue *amqp.Queue
	channel *amqp.Channel
}

func NewSendBroker() (*SendBroker){
	toReturn := new(SendBroker)

	toReturn.procChan = make(chan dto.TextDTO, 100)

	return toReturn
}

func (s *SendBroker) PushText(text dto.TextDTO) error{

	select {
	case s.procChan <- text:
		return nil
	default:
		return errors.New("failed to add event to the channel buffer")
	}
}

func (s *SendBroker) Start(){
	err := s.setupChannel()
	if err != nil{
		panic(err)
	}
	go textRoutine(s)
}

func (s *SendBroker) Stop(){
	s.PushText(dto.TextDTO{Quit:true})
}

func (s *SendBroker) setupChannel() error{

	var err error

	s.connection, s.channel, s.queue, err = rabbit_helper.Setup()
	if err != nil{
		return err
	}

	return nil
}

func (s *SendBroker) SubmitToQueue(textDTO dto.TextDTO) error{

	jDto, err := json.Marshal(textDTO)
	if err != nil{
		log.Println(err)
		return err
	}

	err = s.channel.Publish(
		"",     // exchange
		s.queue.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing {
		ContentType: "application/json",
		Body:        jDto,
	})
	if err != nil{
		log.Println(err)
		return err
	}

	return nil
}