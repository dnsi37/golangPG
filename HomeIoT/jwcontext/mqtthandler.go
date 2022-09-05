package jwcontext

import (
	"HomeIoT/db"
	"fmt"
	"log"

	MQTT "github.com/eclipse/paho.mqtt.golang"
)

var ledStatusHandler MQTT.MessageHandler = func(conn MQTT.Client, msg MQTT.Message) {
	fmt.Printf("TOPIC: %s\n", msg.Topic())
	fmt.Printf("MSG: %s\n", msg.Payload())

	if string(msg.Payload()) == "on" {
		rowsAffected ,err := db.UpdateLedStatus(Context.DB,"on")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Led Status Updated \n row affected : ",rowsAffected)
	}

	if string(msg.Payload()) == "off" {
		rowsAffected ,err := db.UpdateLedStatus(Context.DB,"off")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Led Status Updated \n row affected : ",rowsAffected)
	}
}