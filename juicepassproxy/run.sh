#!/usr/bin/env bashio

MQTT_HOST=$(bashio::services mqtt "host")
MQTT_USER=$(bashio::services mqtt "username")
MQTT_PASS=$(bashio::services mqtt "password")

JUICEBOX_HOST=$(bashio::config 'juicebox_host')
UPDATE_UDPC=true

JPP_HOST=$(bashio::config 'ha_host')

DEVICE_NAME=$(bashio::config 'juicebox_device_name')

/juicepassproxy/docker_entrypoint.sh
