package service

import(
	"github.com/douglasmakey/go-fcm"
	"log"
	"fmt"
	"bitcamp/common/queries"
)

var fcmClient *fcm.Client = nil

//in true hackathon fashion
const API_KEY = "AAAAwSuS614:APA91bESC0N22JD4agjb2tIAAEqpXXc0iaXBhCfhOw6CJYAbPHnhgVh2o4Zs9yQ58rwqx6xnD48f8KWoiY_e7QRY88VcumFg7SrLtalqcxx7p0uRVp5YaiCuJbkuhkXIbljuITnD_ema"

func BroadcastMessage(childId int, message, otherNum string, sent bool) error{

	if fcmClient == nil{
		fcmClient = fcm.NewClient(API_KEY)
	}

	var a string
	if sent{
		a = "sent to"
	}else{
		a = "received from"
	}

	parents, err := queries.GetParentsOfChildByName(childId)
	if err != nil{
		log.Println(err)
		return err
	}

	kid, err := queries.GetUserById(childId)
	if err != nil{
		log.Println(err)
		return err
	}

	data := map[string]interface{}{
		"message": fmt.Sprintf("%s %s %s", kid.Name, a, otherNum),
		"details": map[string]string{
			"content": message,
		},
	}

	parentFcm := []string{}

	for i := 0; i < len(parents); i++{
		parentFcm = append(parentFcm, parents[i].FCMId)
	}

	fcmClient.PushMultiple(parentFcm, data)

	fcmClient.AppendRegistrationIds(parentFcm)

	status, err := fcmClient.Send()
	if err != nil{
		log.Println(err)
		return err
	}

	log.Println(status.Err)
	log.Println(status.StatusCode)
	log.Println(status.RetryAfter)

	log.Println(status)

	return nil
}
