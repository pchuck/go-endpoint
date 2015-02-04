# go-endpoint
#

LEARN=src/github.com/pchuck/learn
ENDPOINT=src/github.com/pchuck/endpoint

console:
	go run console.go

# gin server
get-gin:
	go get github.com/gin-gonic/gin

run-server-gin:
	go run $(ENDPOINT)/server-gin.go

build-server-gin:
	go build $(ENDPOINT)/server-gin.go

install-server-gin:
	go install $(ENDPOINT)/server-gin.go

client-gin:
	go run $(ENDPOINT)/client-gin.go http://localhost:8081/v0/read

client-gin-concurrent:
	go run $(ENDPOINT)/client-gin-concurrent.go http://localhost:8081/v0/read 100

curl-gin:
	curl localhost:8081/v0/read 

# misc
tour:
	go run $(LEARN)/tour.go

concurrency:
	go run $(LEARN)/goroutines.go

# simple server
server:
	go run $(LEARN)/server.go

client:
	curl localhost:4000/time

client_error:
	curl localhost:4000
