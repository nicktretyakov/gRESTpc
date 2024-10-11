#!/bin/bash

start_test_db() {
  sudo docker run --name postgresbeaver -p 5432:5432 -e POSTGRES_USER=postgress -e POSTGRES_PASSWORD=postgress -e POSTGRES_DB -d postgres
}

stop_test_db() {
  docker stop pspdb
  docker rm pspdb
}

while getopts t: flag; do
  case "${flag}" in
    t) type=${OPTARG} ;;
  esac
done

case $type in
  start)
    start_test_db
    echo "pspdb database started."
    ;;
  unit)
    start_test_db
    CGO_ENABLED=0 go test -v -p 1 -count=1 -covermode=count -coverprofile=coverage/c.out -run Unit ./...
    stop_test_db
    echo "pspdb database stopped."
    ;;
  integration)
    start_test_db
    CGO_ENABLED=0 go test -v -p 1 -count=1 -covermode=count -coverprofile=coverage/c.out -run Integration ./...
    stop_test_db
    echo "pspdb database stopped."
    ;;
  *)
    start_test_db
    CGO_ENABLED=0 go test -v -p 1 -count=1 -covermode=count -coverprofile=coverage/c.out ./...
    stop_test_db
    echo "pspdb database stopped."
    ;;
esac
