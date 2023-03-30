package main

import (
	"fmt"

	"github.com/influxdata/telegraf"
	"github.com/influxdata/telegraf/plugins/outputs/mqtt"
)

func main() {
	fmt.Println("Hello, main")

}

type MQTT struct {
	TopicPrefix     string          `toml:"topic_prefix" deprecated:"1.25.0;use 'topic' instead"`
	Topic           string          `toml:"topic"`
	BatchMessage    bool            `toml:"batch" deprecated:"1.25.2;use 'layout = \"batch\"' instead"`
	Layout          string          `toml:"layout"`
	HomieDeviceName string          `toml:"homie_device_name"`
	HomieNodeID     string          `toml:"homie_node_id"`
	Log             telegraf.Logger `toml:"-"`
	mqtt.MQTT
}

func (m *MQTT) Init(broker string, client_id string) error {
	m = new(MQTT)

	m.ClientID = client_id
	m.

	// m.connection = make(chan bool)
	// m.outgoing = make(chan MQTT_Outgoing, 50)
	// m.request_broker_exit = make(chan bool)
	// m.requested_broker_state = make(chan bool, 1)
	// m.broker_connected = make(chan bool, 1)
	// m.params.KeepAlive = 100
	// m.params.PingTimeout = 100
	// m.params.ConnectTimeout = 1
	// m.params.OutgoingMessageTimeout_ms = config.MQTT_Wait_Timeout_ms() // ms

	// m.wg.Add(1)

	// m.params.Broker_Address = broker
	// m.params.ClientId = client_id

	// m.process_list = make(map[string]*MQTT_Topic)

	// go m.runMqtt()

	return m, nil
}

func (m *MQTT) Connect() error {

}