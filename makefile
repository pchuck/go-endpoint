# go-endpoint
#

console:
	go run console.go

server:
	go run server.go

tour:
	go run tour.go

client:
	curl localhost:4000/time

client_error:
	curl localhost:4000
