#!/usr/bin/env bash
docker-compose down ; docker-compose up --build -d ; docker-compose logs -f
