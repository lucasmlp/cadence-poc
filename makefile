include .env
export $(shell sed 's/=.*//' .env)

GOPATH=$(shell go env GOPATH)

deps:
	@ echo
	@ echo "Downloading dependencies..."
	@ echo
	@ go get -v ./...

update-deps:
	@ echo
	@ echo "Updating dependencies..."
	@ echo
	@ go get -u ./...

cadence-containers:
	@ echo
	@ echo "Starting Cassandra Cadence and Cadence Web..."
	@ echo
	@ docker-compose up -d cassandra cadence cadence-web

cadence-worker:
	@ echo
	@ echo "Starting the Cadence Worker..."
	@ echo
	@ go run ./worker/main.go

cadence-replayer:
	@ echo
	@ echo "Starting the Cadence Replayer..."
	@ echo
	@ go run ./replay/main.go

hello-world:
	@ go run ./trigger/main.go HelloWorld

activity:
	@ go run ./trigger/main.go Activity

waiting-signal:
	@ go run ./trigger/main.go WaitingSignal

version:
	@ go run ./trigger/main.go Version

version2:
	@ go run ./trigger/main.go Version2