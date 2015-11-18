
go-endpoint
===========

restful endpoint for a mock service, in golang, deployed via vagrant


## running via vagrant (microservice and client)
```
vagrant up
```

## running locally

* install go - https://golang.org/doc/install or http://golang.org/dl/
* set the environment and run the microservice/client

```

  % . scripts/go_env.sh

  % make run_server_gin &

  % make client_gin_concurrent

```
