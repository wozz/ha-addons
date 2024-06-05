#!/usr/bin/env bashio

# handle deprecated config
if [ bashio::config.exists 'dst' ]; then
  if [ ! bashio::config.exists 'enelx_ip' ]; then
    bashio::option 'enelx_ip' $(bashio::config 'dst')
  fi
  bashio::option 'dst'
fi

export MQTT_HOST="$(bashio::services mqtt 'host')"
export MQTT_USER="$(bashio::services mqtt 'username')"
export MQTT_PASS="$(bashio::services mqtt 'password')"

export JUICEBOX_HOST="$(bashio::config 'juicebox_host')"
export DEVICE_NAME="$(bashio::config 'juicebox_device_name')"
export ENELX_IP="$(bashio::config 'enelx_ip')"
export DEBUG="$(bashio::config 'debug')"

/juicepassproxy/docker_entrypoint.sh
