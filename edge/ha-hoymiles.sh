#!/usr/bin/with-contenv bashio
set +u

bashio::log.green "Starting hoymiles add-on container..."

CONFIG_PATH=/data/options.json
SYSTEM_USER=/data/system_user.json

export MQTT_HOST=$(bashio::services mqtt "host")
export MQTT_USER=$(bashio::services mqtt "username")
export MQTT_PASSWORD=$(bashio::services mqtt "password")

bashio::log.blue "Starting Hoymiles client service..."

./ha-hoymiles