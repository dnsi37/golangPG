package jwcontext

import (
	"HomeIoT/mqtt"
	"fmt"

	MQTT "github.com/eclipse/paho.mqtt.golang"
)

type JwContext struct {
	MC *mqtt.AWSIoTConnection
}

func Init ()  *JwContext {
	Context.ConnectToAWSIoT()
	return Context
}
var Context *JwContext = &JwContext{}

func (c *JwContext) ConnectToAWSIoT (){

	const endPoint string = "a2d6gny2gotvvn-ats.iot.ap-northeast-2.amazonaws.com"
	const certDir string = "C:/Users/junwookim/Desktop/IOT/golangCrt"
	conn, err := mqtt.NewConnection(mqtt.Config{
		KeyPath:  certDir+"/private.pem.key",
		CertPath: certDir + "/certificate.pem.crt",
		CAPath:   certDir + "/AmazonRootCA1.pem",
		ClientId: "ping_client",
		Endpoint: endPoint,
	})
	mqtt.HandleError(err)
	err=conn.SubscribeWithHandler("ping",0, func(client MQTT.Client, message MQTT.Message) {
		fmt.Println(string(message.Payload()))
	})
	mqtt.HandleError(err)
	err=conn.Publish("ping","pong2",0)
	mqtt.HandleError(err)

	c.MC = conn
}