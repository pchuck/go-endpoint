# go-endpoint
#

console:
	go run console.go

# gin server
get-gin:
	go get github.com/gin-gonic/gin

server-gin:
	go run server-gin.go

client-gin-perf:
	go run client-gin-perf.go http://localhost:8081/v0/read 2

# misc
tour:
	go run tour.go

concurrency:
	go run goroutines.go

# simple server
server:
	go run server.go

client:
	curl localhost:4000/time

client_error:
	curl localhost:4000
