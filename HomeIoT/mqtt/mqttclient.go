package mqtt

import (
	iottopic "HomeIoT/mqtt/topic"
	"crypto/tls"
	"crypto/x509"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"time"

	MQTT "github.com/eclipse/paho.mqtt.golang"
)

const endPoint string = "a2d6gny2gotvvn-ats.iot.ap-northeast-2.amazonaws.com"
const certDir string = "C:/Users/junwookim/Desktop/IOT/golangCrt"

var mDefaultHandler MQTT.MessageHandler = func(conn MQTT.Client, msg MQTT.Message) {
	fmt.Printf("TOPIC: %s\n", msg.Topic())
	fmt.Printf("MSG: %s\n", msg.Payload())
}
// ConnectToAWSIoT 
// init function to access and receive conn pointer
func ConnectToAWSIoT () *AWSIoTConnection {

	const endPoint string = "a2d6gny2gotvvn-ats.iot.ap-northeast-2.amazonaws.com"
	const certDir string = "C:/Users/junwookim/Desktop/IOT/golangCrt"
	conn, err := NewConnection(Config{
		KeyPath:  certDir+"/private.pem.key",
		CertPath: certDir + "/certificate.pem.crt",
		CAPath:   certDir + "/AmazonRootCA1.pem",
		ClientId: "ping_client",
		Endpoint: endPoint,
	})
	HandleError(err)
	err=conn.SubscribeWithHandler("ping",0, func(client MQTT.Client, message MQTT.Message) {
		fmt.Println(string(message.Payload()))
	})
	HandleError(err)
	err=conn.Publish("ping","pong2",0)
	HandleError(err)

	return conn

}
func NewConnection (config Config) (*AWSIoTConnection,error) {

	connection := AWSIoTConnection{}
	if err := connection.init(config); err!= nil {
		return nil,err
	}
	connection.client = MQTT.NewClient(connection.options)
	token := connection.client.Connect()
	token.Wait()
	if err := token.Error(); err != nil {
		return nil,err
	}else {
		fmt.Println(token)
		return &connection,nil
	}
}


type Config struct {
	KeyPath  string `json:"keyPath" binding:"required"`
	CertPath string `json:"certPath" binding:"required"`
	CAPath   string `json:"caPath" binding:"required"`
	ClientId string `json:"clientId" binding:"required"`
	Endpoint string `json:"endpoint" binding:"required"`
}

type AWSIoTConnection struct {
	client  MQTT.Client
	options *MQTT.ClientOptions
}



func (c *AWSIoTConnection) init (config Config) error {

	tlsConfig,err := NewTLSConfig(config)
	HandleError(err)
	c.options = MQTT.NewClientOptions()
	c.options.AddBroker("tls://"+config.Endpoint+":8883")
	c.options.SetMaxReconnectInterval(60* 2 *time.Second)
	c.options.SetClientID(config.ClientId)
	c.options.SetTLSConfig(tlsConfig)
	c.options.SetKeepAlive(60* 2 *time.Second)
	c.options.SetDefaultPublishHandler(mDefaultHandler)
	c.options.SetConnectionLostHandler(func(c MQTT.Client, err error) {
		fmt.Println("disconnected")
		log.Fatal(err)
	})
	c.options.SetOnConnectHandler(func(c MQTT.Client) {
		fmt.Println("connected")
		c.Publish(iottopic.Home,0,false,"Server Connected")
	})
	c.options.SetAutoReconnect(true)
	return nil
}


func (c *AWSIoTConnection) Disconnect () bool {
	if c.client.IsConnected() {
		c.client.Disconnect(0)
		return true
	} else {
		return false
	}
}
func (c *AWSIoTConnection) Subscribe(topic string, qos byte) error {
	return c.SubscribeWithHandler(topic, qos, nil)
}

func (c *AWSIoTConnection) SubscribeWithHandler(topic string, qos byte, handler MQTT.MessageHandler) error {
	if !c.client.IsConnected() {
		log.Fatal("client not connected")
		return errors.New("client not connected")
	} else {
		token := c.client.Subscribe(topic, qos, handler)
		token.Wait()
		fmt.Println("token created")
		if err := token.Error(); err != nil {
			fmt.Println("error occured")
			return err
		} else {
			return nil
		}
	}
}
// Unsubscribe function removes subscription for specified topic
func (c *AWSIoTConnection) Unsubscribe(topic string) error {
	if !c.client.IsConnected() {
		log.Fatal("client not connected")
		return errors.New("client not connected")
	} else {
		token := c.client.Unsubscribe(topic)
		token.Wait()
		if err := token.Error(); err != nil {
			return err
		} else {
			return nil
		}
	}
}

// Publish function publishes data in interface on topic with level of qos (Quality of service)
// currently supported 0 & 1 (2 coming in future)
func (c *AWSIoTConnection) Publish(topic string, data interface{}, qos byte) error {
	token := c.client.Publish(topic, qos, false, data)
	token.Wait()
	if err := token.Error(); err != nil {
		return err
	} else {
		return nil
	}
}


func NewTLSConfig(awsConfig Config) (tlsConfig *tls.Config, err error) {

	certpool := x509.NewCertPool()
	pemCerts, err := ioutil.ReadFile(awsConfig.CAPath)
	HandleError(err)
	certpool.AppendCertsFromPEM(pemCerts)

	cert,err := tls.LoadX509KeyPair(awsConfig.CertPath,awsConfig.KeyPath)
	HandleError(err)

	tlsConfig = &tls.Config{
		RootCAs: certpool,
		Certificates: []tls.Certificate{cert},
	}
	return 

}


func HandleError ( err error ){
	if err != nil {
		log.Fatal(err)
	}
}

func MqttPractice() {

	conn, err := NewConnection(Config{
		KeyPath:  certDir+"/private.pem.key",
		CertPath: certDir + "/certificate.pem.crt",
		CAPath:   certDir + "/AmazonRootCA1.pem",
		ClientId: "ping_client",
		Endpoint: endPoint,
	})
	HandleError(err)
	go func() {
		err=conn.SubscribeWithHandler("ping",0, func(client MQTT.Client, message MQTT.Message) {
			fmt.Println(string(message.Payload()))
		})
		HandleError(err)
	}()
	err=conn.Publish("ping","pong2",0)
	HandleError(err)
	
}