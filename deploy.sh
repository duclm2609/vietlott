#!/bin/bash

docker network create docker-network

docker-compose -f docker-compose.mongodb.yml up -d
