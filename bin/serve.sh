#!/bin/bash
if [ "${SERVER_MODE}" = "grpc" ]; then
		go run ${APP_DIR}/grpc.go
else
	 	go run ${APP_DIR}/main.go
fi