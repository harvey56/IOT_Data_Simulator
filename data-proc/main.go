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

	return nil
}
