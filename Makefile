build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build --ldflags '-extldflags "-static" -s -w' -o demo

run:
	make build
	docker build -t iot-demo .
	docker run -d --restart=always --name iot-demo -p 8060:8080 iot-demo