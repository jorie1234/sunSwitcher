package main

import (
	"fmt"
	"github.com/yosssi/gmq/mqtt"
	"github.com/yosssi/gmq/mqtt/client"
	"log"
	"os"
	"os/signal"
	"syscall"
)

type mqttClient struct {
	client     *client.Client
	server     string
	clientname string
}

type mqttMsg struct {
	Topic   string
	Payload string
}

func NewMqtt(server, clientName string) *mqttClient {
	// Set up channel on which to send signal notifications.
	sigc := make(chan os.Signal, 1)
	signal.Notify(sigc, syscall.SIGTERM)

	// Create an MQTT Client.
	cli := client.New(&client.Options{
		// Define the processing of the error handler.
		ErrorHandler: func(err error) {
			fmt.Println(err)
		},
	})
	m := &mqttClient{
		client:     cli,
		server:     server,
		clientname: clientName,
	}

	m.Connect()

	return m
}

func (m *mqttClient) Connect() {

	// Connect to the MQTT Server.
	err := m.client.Connect(&client.ConnectOptions{
		Network:   "tcp",
		Address:   fmt.Sprintf("%s:1883", m.server),
		ClientID:  []byte(m.clientname),
		KeepAlive: 30,
	})
	if err != nil {
		panic(err)
	}

	log.Printf("Connected...")
}

func (m mqttClient) Subscribe(topic string) <-chan mqttMsg {
	c := make(chan mqttMsg)
	// Subscribe to topics.
	err := m.client.Subscribe(&client.SubscribeOptions{
		SubReqs: []*client.SubReq{
			&client.SubReq{
				TopicFilter: []byte(topic),
				QoS:         mqtt.QoS0,
				// Define the processing of the message handler.
				Handler: func(topicName, message []byte) {
					//fmt.Println(string(topicName), string(message))
					msg := mqttMsg{
						Topic:   string(topicName),
						Payload: string(message),
					}
					c <- msg
					//					s := strings.Split(string(topicName), "/")
				},
			},
		},
	})
	if err != nil {
		panic(err)
	}
	return c
}

func (m mqttClient) Publish(topic, msg string) error {
	// Publish a message.
	err := m.client.Publish(&client.PublishOptions{
		QoS:       mqtt.QoS0,
		TopicName: []byte(topic),
		Message:   []byte(msg),
	})
	if err != nil {
		return err
	}
	return nil
}

func (m mqttClient) Close() {
	// Disconnect the Network Connection.
	if err := m.client.Disconnect(); err != nil {
		panic(err)
	}

	// Terminate the Client.
	m.client.Terminate()
}
