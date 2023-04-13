# IOT_Data_Simulator

IOT Data Simulator

## Homie

The application follows the Homie convention. For more details, see here : https://homieiot.github.io/

## Telegraf

To learn about Telegraf : https://www.influxdata.com/time-series-platform/telegraf/

MQTT consumer plugin : https://github.com/influxdata/telegraf/tree/master/plugins/inputs/mqtt_consumer

MQTT producer plugin : https://github.com/influxdata/telegraf/tree/master/plugins/outputs/mqtt

# Run container

docker-compose up -d

To check the services are running :

docker container ps

You should see the 3 services running (influxdb, grafana, telegraf).
If you don't see telegraf listed, that most likely means there is an error with the telegraf.conf file.

docker logs telegraf will indicate where the configuration error is.

Crate a network so all containers can exchange data :
docker network create iot

To access influxDB CLI :
docker exec -it influxdb /bin/bash

Check that your InfluxDB server is correctly listening on port 8086 on your host :
netstat -tulpn | grep 8086
