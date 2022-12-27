package jwcontext

import (
	"HomeIoT/mqtt"
	iottopic "HomeIoT/mqtt/topic"
	"database/sql"
	"fmt"
	"log"
	"time"

	MQTT "github.com/eclipse/paho.mqtt.golang"
)

type JwContext struct {
	MC *mqtt.AWSIoTConnection
	DB *sql.DB
}

// Init : initialize context as global variable 
func Init ()  *JwContext {
	Context.ConnectToAWSIoT()
	go Context.InitDB()
	return Context
}
// Global Variable :: careful when it be used 
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
	err=conn.SubscribeWithHandler(iottopic.Home,0, func(client MQTT.Client, message MQTT.Message) {
		fmt.Println(string(message.Payload()))
	})
	if err != nil {
		log.Fatal(err)
	}
	err=conn.SubscribeWithHandler(iottopic.LedStatus,0,ledStatusHandler)
	if err != nil {
		log.Fatal(err)
	}
	mqtt.HandleError(err)
	err=conn.Publish(iottopic.Home,"Back end Server Connected",0)
	if err != nil {
		log.Fatal(err)
	}
	mqtt.HandleError(err)

	c.MC = conn
}

func (c *JwContext) InitDB ()  {

	db, err := sql.Open("mysql", "junwoo:junwoo123@tcp(junwoodb.clcwfeh6dtye.ap-northeast-2.rds.amazonaws.com:3306)/iotdb")
	if err != nil {
		log.Fatal(err)
	}	
	c.DB = db
	defer db.Close()
	
	// Check db connection periodically and retry if conn is lost 
	for {
		err := c.DB.Ping()
		if err!= nil {
			fmt.Println("DB Ping Error", err)
			c.DB.Close()
			c.DB , err = sql.Open("mysql", "junwoo:junwoo123@tcp(junwoodb.clcwfeh6dtye.ap-northeast-2.rds.amazonaws.com:3306)/iotdb")
			if err!= nil {
				fmt.Println("Failed to reconnect DB", err)
			}
		}
		time.Sleep(time.Second * 60)
	}
	
}