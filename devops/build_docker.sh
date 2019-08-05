#!/usr/bin/env bash

cp ./config/service.template.config.json ./config/service.config.json

sed -i "s/Your Username/$1/g" "./config/service.config.json"
sed -i "s/Your Password/$2/g" "./config/service.config.json"
