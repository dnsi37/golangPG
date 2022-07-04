package main

import (
	"fmt"
	"practice/mqttPractice"
	"practice/netPractice"

	MQTT "github.com/eclipse/paho.mqtt.golang"
)

const endPoint string = "a2d6gny2gotvvn-ats.iot.ap-northeast-2.amazonaws.com"
const certDir string = "C:/Users/junwookim/Desktop/IOT/golangCrt"

func main() {
	conn, err := mqttPractice.NewConnection(mqttPractice.Config{
		KeyPath:  certDir+"/private.pem.key",
		CertPath: certDir + "/certificate.pem.crt",
		CAPath:   certDir + "/AmazonRootCA1.pem",
		ClientId: "ping_client",
		Endpoint: endPoint,
	})
	mqttPractice.HandleError(err)
	err=conn.SubscribeWithHandler("ping",0, func(client MQTT.Client, message MQTT.Message) {
		fmt.Println(string(message.Payload()))
	})
	mqttPractice.HandleError(err)
	err=conn.Publish("ping","pong2",0)
	mqttPractice.HandleError(err)

	
	netPractice.HttpPractice()
	
}
