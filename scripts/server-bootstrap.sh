cd /home/vagrant/sync
export GOPATH=/home/vagrant/sync/src
export GOBIN=/home/vagrant/sync/bin
export ENDPOINT=src/github.com/pchuck/endpoint
yum install -y go git
echo "fetching gin framework"
go get github.com/gin-gonic/gin
echo "building server.."
go build $ENDPOINT/server-gin.go
echo "installing server.."
go install $ENDPOINT/server-gin.go
echo "running server.."
#go run $ENDPOINT/server-gin.go >& /tmp/go.out
bin/server-gin >& /tmp/go.out &

