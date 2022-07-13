package main

import (
	mqttclient "HomeIoT/mqtt"
	"HomeIoT/server"
	"fmt"

	MQTT "github.com/eclipse/paho.mqtt.golang"
)



func main() {
	
	ConnectToAWSIoT()
	server.Init()
	
}


func ConnectToAWSIoT () *mqttclient.AWSIoTConnection {

	const endPoint string = "a2d6gny2gotvvn-ats.iot.ap-northeast-2.amazonaws.com"
	const certDir string = "C:/Users/junwookim/Desktop/IOT/golangCrt"
	conn, err := mqttclient.NewConnection(mqttclient.Config{
		KeyPath:  certDir+"/private.pem.key",
		CertPath: certDir + "/certificate.pem.crt",
		CAPath:   certDir + "/AmazonRootCA1.pem",
		ClientId: "ping_client",
		Endpoint: endPoint,
	})
	mqttclient.HandleError(err)
	err=conn.SubscribeWithHandler("ping",0, func(client MQTT.Client, message MQTT.Message) {
		fmt.Println(string(message.Payload()))
	})
	mqttclient.HandleError(err)
	err=conn.Publish("ping","pong2",0)
	mqttclient.HandleError(err)

	return conn

}