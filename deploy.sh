#!/usr/bin/env bash
cd frontend

yarn build

cd ..

docker-compose -f docker-compose.yaml up -d --build
