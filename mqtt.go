package main

import (
	MQTT "github.com/eclipse/paho.mqtt.golang"
	"time"
)

type mqttClient struct {
	client     MQTT.Client
	server     string
	clientname string
}

type mqttMsg struct {
	Topic   string
	Payload string
}

func NewMqtt(server, clientName string) *mqttClient {
	// Set up channel on which to send signal notifications.
	opts := MQTT.NewClientOptions().AddBroker(server)
	opts.SetClientID(clientName)
	m := mqttClient{}

	//create and start a client using the above ClientOptions
	m.client = MQTT.NewClient(opts)
	if token := m.client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	return &m
}

func (m mqttClient) Publish(topic, msg string) error {
	// Publish a message.
	token := m.client.Publish(topic, 0, false, msg)
	token.WaitTimeout(30 * time.Second)
	return token.Error()
}

func (m mqttClient) Close() {
	m.client.Disconnect(250)
}
