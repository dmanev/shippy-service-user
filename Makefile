# user-service/Makefile

gobuild:
	CGO_ENABLED=0 GOOS=linux
	go build -o shippy-service-user -a -installsuffix cgo  main.go repository.go

build:
	protoc -I. --go_out=plugins=micro:. \
	 proto/user/user.proto
	docker build -t shippy-service-user .

run:
	docker run -p 50052:50051 -e MICRO_SERVER_ADDRESS=:50051 shippy-service-user
