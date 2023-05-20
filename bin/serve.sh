if [ "${SERVER_MODE}" == "grpc" ];
	then
		go run ../grpc.go
	else
	 	go run ../main.go
fi