build:
	echo "Build proto file"
	protoc -I. --go_out=plugins=micro:. \
		  proto/user/user.proto

	echo "Build docker image"
	docker build -t user-service .

run:
	echo "Run docker image"
	docker run -p 50053:50051 -e MICRO_SERVER_ADDRESS=:50051 user-service
