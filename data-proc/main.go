package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"strings"

	// "time"

	// viper_config "github.com/spf13/viper"
	homie "github.com/andig/homie"
	// autopaho "github.com/eclipse/paho.golang/autopaho"
	paho "github.com/eclipse/paho.golang/paho"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	// "github.com/influxdata/telegraf"
	// telegrafmqtt "github.com/influxdata/telegraf/plugins/outputs/mqtt"
)

type MQTT struct {
	TopicPrefix     string `toml:"topic_prefix" deprecated:"1.25.0;use 'topic' instead"`
	Topic           string `toml:"topic"`
	BatchMessage    bool   `toml:"batch" deprecated:"1.25.2;use 'layout = \"batch\"' instead"`
	Layout          string `toml:"layout"`
	HomieDeviceName string `toml:"homie_device_name"`
	HomieNodeID     string `toml:"homie_node_id"`
	// Log             telegraf.Logger `toml:"-"`
}

// pattern “homie/device/node/property”
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
type HomieTopic struct {
	Homie  string
	Device string
	Node   string
}
type MQTT_Config struct {
	Broker_Address string `json:"broker_address" mapstructure:"broker_address"`
	Broker_Port    uint   `json:"broker_port" mapstructure:"broker_port"`
	Username       string `json:"username" mapstructure:"username"`
	Password       string `json:"password" mapstructure:"password"`
}
type Config struct {
	ClientID string
	// BrokerURL *url.URL
	// main_config     *viper_config.Viper
	Hostname string
	MQTT     MQTT_Config
}
type Client struct {
	conn     *paho.Client
	cp       *paho.Connect
	pb       *paho.Publish
	co       *net.Conn
	topic    string
	qos      byte
	config   *Config
	gateways map[string]bool
}

var (
	Main *Config
)

func Config_Initialization() {
	Main = newConfig()
}

func newConfig() *Config {
	config := &Config{}

	hostname, err_hostname := os.Hostname()
	if err_hostname != nil {
		hostname = "localhost"
	}
	hostname = strings.ReplaceAll(hostname, " ", "")
	config.Hostname = hostname
	// config.main_config = viper_config.New()

	// config.main_config.SetConfigName(main_config_filename)
	// config.main_config.AddConfigPath("./")
	// config.main_config.SetConfigType("json")

	config.MQTT.Broker_Address = "localhost"
	config.MQTT.Broker_Port = 1883
	config.MQTT.Username = ""
	config.MQTT.Password = ""

	// config.main_config.SetDefault("mqtt", config.MQTT)

	return config
}

func myHandler(p *paho.Publish) {
	fmt.Printf("Received message on topic %s: %s\n", p.Topic, p.Payload)
}

func main() {
	fmt.Println("Hello, main")

	Config_Initialization()
	server := fmt.Sprintf("%s:%d", newConfig().MQTT.Broker_Address, newConfig().MQTT.Broker_Port)

	conn, err := net.Dial("tcp", server)
	if err != nil {
		log.Fatalf("Failed to connect to %s: %s", server, err)
	}

	c := new(Client)
	c.co = &conn
	c.Init()
	c.Connect(newConfig(), server)

	props := &paho.PublishProperties{}
	msg := "GREEN"
	p := &paho.Publish{
		QoS:        0,
		Topic:      "Homie/Light/Color",
		Properties: props,
		Payload:    []byte(msg),
	}
	c.Publish(p)
	conn.Close()

	// Create a client
	// You can generate an API Token from the "API Tokens Tab" in the UI
	// client := influxdb2.NewClient("$DOCKER_INFLUXDB_URL", "$DOCKER_INFLUXDB_INIT_ADMIN_TOKEN")
	// always close client at the end
	// defer client.Close()
}

func (client *Client) Init() *paho.Client {
	c := paho.NewClient(paho.ClientConfig{
		Router: paho.NewSingleHandlerRouter(myHandler),
		Conn:   *client.co,
	})
	client.conn = c

	return c
}

func (client *Client) Connect(config *Config, server string) {
	client.cp = &paho.Connect{
		KeepAlive:  30,
		ClientID:   config.ClientID,
		CleanStart: true,
		Username:   config.MQTT.Username,
		Password:   []byte(config.MQTT.Password),
	}

	ca, err := client.conn.Connect(context.Background(), client.cp)
	if err != nil {
		log.Fatalln(err)
	}
	if ca.ReasonCode != 0 {
		log.Fatalf("Failed to connect to %s : %d - %s", server, ca.ReasonCode, ca.Properties.ReasonString)
	}

	fmt.Printf("Connected to %s\n", server)
}

func (client *Client) Publish(p *paho.Publish) {
	client.pb = &paho.Publish{
		Topic:      *&p.Topic,
		QoS:        byte(*&p.QoS),
		Payload:    []byte(p.Payload),
		Properties: p.Properties,
	}

	if _, err := client.conn.Publish(context.Background(), client.pb); err != nil {
		log.Println(err)
	}

	fmt.Printf("message published...")
}

// func (c *Client) Test() error {
// 	var homietopic *HomieTopic
// 	homietopic = &HomieTopic{Homie: "homie", Device: "firefly", Node: "Panel1"}
// 	topic := homietopic.Homie + "/" + homietopic.Device + "/" + homietopic.Node
// 	var property *homie.Property
// 	property = &homie.Property{Name: "Color", Value: "GREEN"}

// 	// config := SampleConfig(&homietopic, &property)
// 	c.Publish(context.Background(), topic, property.Value, byte(1))

// 	return nil
// }

func SampleConfig(topic *HomieTopic, property *homie.Property) *homie.Device {
	d := homie.NewDevice()

	if n, _ := d.NewNode(topic.Node); true {
		n.Name = topic.Node

		if p, _ := n.NewProperty(property.Name); true {
			p.DataType = homie.DataTypeFloat
			p.Name = property.Name
			p.Value = property.Value
			p.Settable = true
		}
	}

	return d

	// // template mqtt client options
	// opt := mqtt.NewClientOptions()
	// opt.AddBroker(broker)
	// opt.SetAutoReconnect(true)

	// root topic for device
	// rootTopic := homie.DefaultRootTopic + "/test"

	// // mqtt client connection with cloned options and last will
	// handler := paho.NewHandler(topic, opt, qos)
	// handler.Timeout = 1 * time.Second
	// handler.ErrorHandler = paho.Log
	// if t := handler.Client.Connect(); !t.WaitTimeout(time.Second) {
	// 	log.Fatalf("could not connect: %v", t.Error())
	// }

	// // publish the device using handler's Publish method
	// d.Publish(handler.Publish, topic)
	// time.Sleep(time.Second)

	// // omitting the Disconnect() will set the device state to "lost"
	// handler.Client.Disconnect(1000)

	return nil
}

type MqttConfig struct {
	URL      string `help:"MQTT broker to send messages to" required:""`
	Username string `help:"Username to authenticate with"`
	Password string `help:"Password to authenticate with"`
}

// func New(ctx context.Context, cfg MqttConfig) (*Client, error) {
// 	u, err := url.Parse(cfg.URL)
// 	if err != nil {
// 		return nil, fmt.Errorf("parsing broker URL '%s': %w", cfg.URL, err)
// 	}

// 	if u.Scheme != "mqtt" && u.Scheme != "tls" {
// 		return nil, fmt.Errorf("URL scheme '%s' not supported, should be 'mqtt' or 'tls'", u.Scheme)
// 	}

// 	mqttConfig := autopaho.ClientConfig{
// 		BrokerUrls: []*url.URL{u},
// 		TlsCfg:     nil,
// 		KeepAlive:  5 * 60,
// 	}

// 	mqttConfig.SetUsernamePassword(cfg.Username, []byte(cfg.Password))

// 	conn, err := autopaho.NewConnection(ctx, mqttConfig)

// 	return &Client{
// 		conn: conn,
// 	}, nil
// }

// func (client *InfluxDBType) WriteDataToDB() error {
// 	// get non-blocking write client
// 	writeAPI := client.WriteAPI("$DOCKER_INFLUXDB_INIT_ORG", "$DOCKER_INFLUXDB_INIT_BUCKET")

// 	p := influxdb2.NewPoint("stat",
// 		map[string]string{"unit": "temperature"},
// 		map[string]interface{}{"avg": 24.5, "max": 45},
// 		time.Now())
// 	// write point asynchronously
// 	writeAPI.WritePoint(p)
// 	// create point using fluent style
// 	p = influxdb2.NewPointWithMeasurement("stat").
// 		AddTag("unit", "temperature").
// 		AddField("avg", 23.2).
// 		AddField("max", 45).
// 		SetTime(time.Now())
// 	// write point asynchronously
// 	writeAPI.WritePoint(p)
// 	// Flush writes
// 	writeAPI.Flush()

// 	return nil
// }

// func (client *InfluxDBType) FluxQuery(queryParams *FluxQuery) {
// 	// Get query client
// 	queryAPI := client.QueryAPI("$DOCKER_INFLUXDB_INIT_ORG")

// 	query := fmt.Sprintf(`from(bucket:%s)|> range(start: %s) |> filter(fn: (r) => r._measurement == "stat")`, queryParams.bucket, queryParams.rangeStart)

// 	// get QueryTableResult
// 	result, err := queryAPI.Query(context.Background(), query)
// 	if err != nil {
// 		panic(err)
// 	}

// 	// Iterate over query response
// 	for result.Next() {
// 		// Notice when group key has changed
// 		if result.TableChanged() {
// 			fmt.Printf("table: %s\n", result.TableMetadata().String())
// 		}
// 		// Access data
// 		fmt.Printf("value: %v\n", result.Record().Value())
// 	}
// 	// check for an error
// 	if result.Err() != nil {
// 		fmt.Printf("query parsing error: %\n", result.Err().Error())
// 	}
// }

// func (m *MQTT) Write(metrics []telegraf.Metric) error

// func NewTopicNameGenerator(topicPrefix string, topic string) (*TopicNameGenerator, error)

// func (t *TopicNameGenerator) Generate(hostname string, m telegraf.Metric) (string, error)
// func (t *TopicNameGenerator) Tag(key string) string

// func NewHomieGenerator(tmpl string) (*HomieGenerator, error)

// func (t *HomieGenerator) Generate(m telegraf.Metric) (string, error)
// func (t *HomieGenerator) Tag(key string) string

// const (
// 	announcementTopicPrefix = "hb/announce/gw/#"
// 	commandTopicFmt         = "hb/gw/%s/cmd/#"  // expects the GWid
// 	sampleTopicFmt          = "hb/gw/%s/data/#" // expects the GWid
// 	keepAliveSeconds        = 15
// 	connectRetryDelay       = 5 * time.Second
// 	connectTimeout          = 5 * time.Second
// 	commandTopicQoS         = 1
// )

// const (
// 	connectionWaitTime = 1000 // milliseconds
// 	timeFormat         = "2006-01-02-15-04-05"
// )

// type mqttAdapter struct {
// 	connectionManager *autopaho.ConnectionManager
// 	clientID          string
// 	topic             string
// 	qos               byte
// }

// func NewMqttAdapter(ctx context.Context) (repository.PayloadUploadRepository, error) {
// 	cfg, err := getConfig()
// 	if err != nil {
// 		return nil, err
// 	}

// 	tlsCfg, err := newTLSConfig()
// 	if err != nil {
// 		return nil, err
// 	}

// 	mqttCfg := getMqttConfig(cfg, tlsCfg)

// 	cm, err := autopaho.NewConnection(ctx, mqttCfg)
// 	if err != nil {
// 		return nil, err
// 	}

// 	time.Sleep(connectionWaitTime * time.Millisecond)

// 	return &mqttAdapter{
// 		connectionManager: cm,
// 		clientID:          cfg.clientID,
// 		topic:             cfg.topic,
// 		qos:               cfg.qos,
// 	}, nil
// }

// func (a *mqttAdapter) Upload(ctx context.Context, payload *model.Payload) ([]model.BaseFilePath, error) {
// 	if len(payload.FilePaths) == 0 {
// 		return []model.BaseFilePath{}, nil
// 	}

// 	topic, err := a.createTopic(payload)
// 	if err != nil {
// 		return nil, err
// 	}

// 	if _, err := a.connectionManager.Publish(ctx, &paho.Publish{
// 		QoS:     a.qos,
// 		Topic:   topic,
// 		Payload: payload.Message,
// 	}); err != nil {
// 		return nil, errors.Wrapf(err, "failed to publish")
// 	}

// 	return payload.FilePaths, nil
// }

// func getMqttConfig(cfg config, tlsCfg *tls.Config) autopaho.ClientConfig {
// 	return autopaho.ClientConfig{
// 		BrokerUrls:        []*url.URL{cfg.serverURL},
// 		KeepAlive:         cfg.keepAlive,
// 		ConnectRetryDelay: cfg.connectRetryDelay,
// 		OnConnectionUp:    func(*autopaho.ConnectionManager, *paho.Connack) { log.Print("mqtt connection up") },
// 		OnConnectError:    func(err error) { log.Printf("error whilst attempting connection: %s\n", err) },
// 		Debug:             paho.NOOPLogger{},
// 		TlsCfg:            tlsCfg,
// 		ClientConfig: paho.ClientConfig{
// 			ClientID:      cfg.clientID,
// 			OnClientError: func(err error) { log.Printf("server requested disconnect: %s\n", err) },
// 			OnServerDisconnect: func(d *paho.Disconnect) {
// 				if d.Properties != nil {
// 					log.Printf("server requested disconnect: %s\n", d.Properties.ReasonString)
// 				} else {
// 					log.Printf("server requested disconnect; reason code: %d\n", d.ReasonCode)
// 				}
// 			},
// 		},
// 	}
// }

// func (a *mqttAdapter) createTopic(payload *model.Payload) (string, error) {
// 	t, err := time.Parse(timeFormat, strings.TrimSuffix(string(payload.FilePaths[0]), ".dat"))
// 	if err != nil {
// 		return "", errors.Wrapf(err, "failed to parse date")
// 	}

// 	topic := a.topic + "/thing=" + a.clientID + t.Format("/year=2006/month=01/day=02/") + string(payload.FilePaths[0])

// 	return topic, nil
// }
