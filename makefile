# go-endpoint
#

hello:
	go run hello.go

server:
	go run server.go

client:
	curl localhost:4000/time

client_error:
	curl localhost:4000
