# go-endpoint
#

LEARN=src/github.com/pchuck/learn
ENDPOINT=src/github.com/pchuck/endpoint
DBHOST=localhost:27017
PORT=8081

## environment
create_env:
	tmuxinator start go-endpoint


## local development/deployment

# go/gin server
get_gin:
	go get github.com/gin-gonic/gin

build_server_gin: # creates server-gin
	go build $(ENDPOINT)/server-gin.go

install_server_gin: # installs in bin/
	go install $(ENDPOINT)/server-gin.go

run_server_gin:
	go run $(ENDPOINT)/server-gin.go

# go clients
client_gin:
	go run $(ENDPOINT)/client-gin.go http://localhost:$(PORT)/v0/read

client_gin_concurrent:
	go run $(ENDPOINT)/client-gin-concurrent.go http://localhost:$(PORT)/v0/read 100

curl_gin:
	curl localhost:$(PORT)/v0/read 

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


## remote vagrant deployment

VAGRANT_PROVIDER=virtualbox

# add a box
add_centos:
	vagrant box add centos/7 --provider=$(VAGRANT_PROVIDER)

# update boxes
update_boxes:
	vagrant box update 

# list available boxes
list:
	vagrant box list

# virtualbox extension updater
vbguest:
	vagrant plugin install vagrant-vbguest

# initialize a vagrant environment
init:
	vagrant init

# start all instances
up:
	vagrant up

# show machine states
status:
	vagrant status

# restart/update the web instance
# !!! broken in multimachine?
reload:
	vagrant reload --provision

# connect to instances by name
ssh_web:
	vagrant ssh web

ssh_console:
	vagrant ssh console

# also see: halt, suspend
destroy:
	vagrant destroy
