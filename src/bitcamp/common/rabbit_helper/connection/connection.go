package connection

import(
	"github.com/streadway/amqp"
	"log"
	"os"
	"fmt"
)

var rabbitCon *amqp.Connection

func GetDb() (*amqp.Connection, error){
	if rabbitCon != nil{
		return rabbitCon, nil
	}else{

		var location string

		if os.Getenv("DOCKER") == "true"{
			location = "rabbitmq"
		}else{
			location = "localhost"
		}
		locDb, err := amqp.Dial(fmt.Sprintf("amqp://guest:guest@%s:5672/", location))
		if err != nil{
			log.Println(err)
			return nil, err
		}

		rabbitCon = locDb
		return rabbitCon, nil
	}

}
