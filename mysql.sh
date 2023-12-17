#!/bin/bash

service_name="mysql"

if [ "$1" == "run" ]; then
    echo "Starting MySQL!"
    brew services start $service_name

elif [ "$1" == "kill" ]; then
    echo "Killing MySQL"
    brew services stop $service_name
else
    echo "Undefined argument!"
fi