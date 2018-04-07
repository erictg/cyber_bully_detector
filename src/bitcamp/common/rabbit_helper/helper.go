package rabbit_helper

import (
	"log"
	"fmt"
	"bitcamp/common/rabbit_helper/connection"
	"github.com/streadway/amqp"
)

func FailOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
		panic(fmt.Sprintf("%s: %s", msg, err))
	}
}

func Setup() (p_conn *amqp.Connection, p_chan *amqp.Channel, p_queue *amqp.Queue, p_error error){
	conn, err := connection.GetDb()
	if err != nil{
		return nil, nil, nil, err
	}

	p_conn = conn

	ch, err := p_conn.Channel()
	if err != nil{
		log.Println(err.Error())
		return nil, nil, nil, err
	}

	p_chan = ch

	q, err := ch.QueueDeclare(
		"proc", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	if err != nil{
		log.Println(err)
		return nil, nil, nil, err
	}

	p_queue = &q

	return p_conn, p_chan, p_queue, nil
}