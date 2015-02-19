# go-endpoint
#

LEARN=src/github.com/pchuck/learn
ENDPOINT=src/github.com/pchuck/endpoint
DBHOST=localhost:27017

# tmux a complete dev environment
create_env:
	scripts/go-endp-tmux.sh work go-endp

# gin server
get_gin:
	go get github.com/gin-gonic/gin

build_server_gin:
	go build $(ENDPOINT)/server-gin.go

install_server_gin:
	go install $(ENDPOINT)/server-gin.go

run_server_gin:
	go run $(ENDPOINT)/server-gin.go

client_gin:
	go run $(ENDPOINT)/client-gin.go http://localhost:8081/v0/read

client_gin_concurrent:
	go run $(ENDPOINT)/client-gin-concurrent.go http://localhost:8081/v0/read 100

curl_gin:
	curl localhost:8081/v0/read 

# mongodb
get_mgo:
	go get gopkg.in/mgo.v2

start_mdb:
	mongod --dbpath ~/.mdb-db &

client_mdb:
	go run $(ENDPOINT)/client-mdb.go localhost:27017 test people


# misc
tour:
	go run $(LEARN)/tour.go

concurrency:
	go run $(LEARN)/goroutines.go

console:
	go run $(LEARN)/console.go

# simple server
simple_server:
	go run $(LEARN)/server.go

simple_client:
	curl localhost:4000/time

simple_client_error:
	curl localhost:4000
