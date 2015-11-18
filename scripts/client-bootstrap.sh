cd /home/vagrant/sync
export GOPATH=/home/vagrant/sync/src
export ENDPOINT=src/github.com/pchuck/endpoint
export HOST=192.168.1.11
export PORT=8081
yum install -y go
echo "running client.."
go run $ENDPOINT/client-gin-concurrent.go http://$HOST:$PORT/v0/read 100

