#!/bin/bash

echo "restarting container ..."
docker-compose restart web
docker-compose logs -f web
