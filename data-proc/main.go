package main

import (
	"context"
	"fmt"
	"time"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"

	"github.com/influxdata/telegraf"
	"github.com/influxdata/telegraf/plugins/outputs/mqtt"
)

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
type HomieGenerator struct {
	PluginName string
	// contains filtered or unexported fields
}
type TopicNameGenerator struct {
	Hostname    string
	TopicPrefix string
	PluginName  string
	// contains filtered or unexported fields
}
type InfluxDBType struct {
	influxdb2.Client
}
type FluxQuery struct {
	bucket       string
	rangeStart   string
	filterFn     func()
	groupColumns []string
}

func NewHomieGenerator(tmpl string) (*HomieGenerator, error)
func (t *HomieGenerator) Generate(m telegraf.Metric) (string, error)
func (t *HomieGenerator) Tag(key string) string

func (m *MQTT) Close() error {
	m.Close()
	// m.request_broker_exit <- true

	return nil
}

func (m *MQTT) Connect() error {
	m.Connect()

	return nil
}

func (m *MQTT) Init() error
func (*MQTT) SampleConfig() string
func (m *MQTT) Write(metrics []telegraf.Metric) error

func NewTopicNameGenerator(topicPrefix string, topic string) (*TopicNameGenerator, error)
func (t *TopicNameGenerator) Generate(hostname string, m telegraf.Metric) (string, error)
func (t *TopicNameGenerator) Tag(key string) string

func main() {
	fmt.Println("Hello, main")

	// Create a client
	// You can generate an API Token from the "API Tokens Tab" in the UI
	client := influxdb2.NewClient("$DOCKER_INFLUXDB_URL", "$DOCKER_INFLUXDB_INIT_ADMIN_TOKEN")
	// always close client at the end
	defer client.Close()

	// m := new(MQTT)
	// client. WriteDataToDB()

}

func (client *InfluxDBType) WriteDataToDB() error {
	// get non-blocking write client
	writeAPI := client.WriteAPI("$DOCKER_INFLUXDB_INIT_ORG", "$DOCKER_INFLUXDB_INIT_BUCKET")

	p := influxdb2.NewPoint("stat",
		map[string]string{"unit": "temperature"},
		map[string]interface{}{"avg": 24.5, "max": 45},
		time.Now())
	// write point asynchronously
	writeAPI.WritePoint(p)
	// create point using fluent style
	p = influxdb2.NewPointWithMeasurement("stat").
		AddTag("unit", "temperature").
		AddField("avg", 23.2).
		AddField("max", 45).
		SetTime(time.Now())
	// write point asynchronously
	writeAPI.WritePoint(p)
	// Flush writes
	writeAPI.Flush()

	return nil
}

func (client *InfluxDBType) FluxQuery(queryParams *FluxQuery) {
	// Get query client
	queryAPI := client.QueryAPI("$DOCKER_INFLUXDB_INIT_ORG")

	query := fmt.Sprintf(`from(bucket:%s)|> range(start: %s) |> filter(fn: (r) => r._measurement == "stat")`, queryParams.bucket, queryParams.rangeStart)

	// get QueryTableResult
	result, err := queryAPI.Query(context.Background(), query)
	if err != nil {
		panic(err)
	}

	// Iterate over query response
	for result.Next() {
		// Notice when group key has changed
		if result.TableChanged() {
			fmt.Printf("table: %s\n", result.TableMetadata().String())
		}
		// Access data
		fmt.Printf("value: %v\n", result.Record().Value())
	}
	// check for an error
	if result.Err() != nil {
		fmt.Printf("query parsing error: %\n", result.Err().Error())
	}
}
