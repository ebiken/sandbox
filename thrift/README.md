# /sandbox/thrift

This directory includes test code for trying out thrift + Golang.

# References

* Blog: Implement Golang RPC Service Based on Apache Thrift
	* https://xuri.me/2016/06/27/implement-golang-rpc-service-based-on-apache-thrift.html


# Building sample server/client app.

1. Create "mythrift.thrift"
2. Generate gen-go/mythrift
```
$ thrift -r --gen go mythrift.thrift
```
3. Implement Server
```
mkdir thrift-server
vi thrift-server/main.go
```
4. Implement Client
```
mkdir thrift-client
vi thrift-client/main.go
```
5. Build and run Server
```
$ go install github.com/ebiken/sandbox/thrift/thrift-server
$ thrift-server
thrift server in 127.0.0.1:9090
```
6. Run thrift client
```
$ go run github.com/ebiken/sandbox/thrift/thrift-client/main.go
```

