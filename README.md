Request count tracking API
==========================
##### Version: v1.0

Based on given token, you can track request counts.  

Counter service generates a unique token and tracks request counts from the same user (token).  
Envoy service accepts a request at any path. On /auth, the envoy service returns a token provided by the counter service.  

If a request is made with an Authorization header and valid token, the envoy service returns a request count for the request path.  

Example series of requests:  

  GET /auth   	returns { "token": "token123" }  

  GET / 	      returns { "count": 1 } // using token123  

  GET /       	returns { "count": 2 } // using token123  

  GET /my/path 	returns { "count": 1 } // using token123  

  GET /auth   	returns { “token”: “token456” }  

  GET /       	returns { “count”: 1 } // using token456  

  GET / 	      returns { “count”: 3 } // using token123  


Prerequests
-----------
Go 1.9 or later  
Redis  

Go instalation guide https://golang.org/doc/install  
Redis installation guide https://redis.io/topics/quickstart  

Initial local setup for Linux  
-----------------------------

Clone repository to your $GOPATH/src/ path
```
git clone https://github.com/vistiria/tracker
```

Start redis  
```
redis-server
```

Run counter service  
(Counter service will run on port 6381 by default, if you want to run it on different port, you have to set   COUNTER_ADDR environment variable, e.g. export COUNTER_ADDR=127.0.0.1:6373.  
Redis is running on port 6379 by default. If you run Redis on different port you have to set REDIS_ADDR   environment variable, e.g. export REDIS_ADDR=127.0.0.1:6372)  
```
# go to counter directory
cd $GOPATH/src/tracker/counter

# install dependencies
go get -t -v ./...

# compile
go build

# run counter service
./counter
# or
go run main.go
```

Run envoy service  
(Envoy service will run on port 6380 by default, if you want to run it on different port, you have to set ENVOY_ADDR environment variable, e.g. export ENVOY_ADDR=127.0.0.1:6374.
Counter service will run on port 6381 by default, if you want to run it on different port, you have to set   COUNTER_ADDR environment variable, e.g. export COUNTER_ADDR=127.0.0.1:6373.)  

```
# go to envoy directory
cd $GOPATH/src/tracker/envoy

# install dependencies
go get -t -v ./...

# compile
go build

# run envoy service
./envoy
# or
go run  main.go
```

To compile the protocol buffer definition install protobuf.  
protobuf installation guide: https://github.com/golang/protobuf  
```
cd $GOPATH/src/tracker/manager_grpc
protoc --go_out=plugins=grpc:./rpc/ manager_rpc.proto
```


Endpoints:
----------
GET /auth  
GET /<any path>   // the Authorization header with the token value is required

You have to add the Authorization header with the token value to the second request.  
To get token use /auth endpoint  

Example curl commands  
```
curl "http://localhost:6380/auth" -v
# {"token":"c6f6fd4b-b347-4aa1-b22d-c2a786eb82fd"}

curl "http://localhost:6380/my/path" -v -H "Authorization: c6f6fd4b-b347-4aa1-b22d-c2a786eb82fd"
# {"count":1}
```

Tests
-----

You can run a simple integration test sending 4003 requests using 3 envoy services.  
Make sure counter service and Redis are running.  
```
#run envoy services on ports 6382, 6383, 6384 (in $GOPATH/src/tracker/envoy path)
ENVOY_ADDR=127.0.0.1:6382 go run main.go
ENVOY_ADDR=127.0.0.1:6383 go run main.go
ENVOY_ADDR=127.0.0.1:6384 go run main.go

# install dependencies
pip install tornado

# run python script
cd $GOPATH/src/tracker
python send_requests.py
```

Run counter unit tests  
```
cd $GOPATH/src/tracker/counter/service
go test
```

Discussion:
------------
I would use disc memory database instead of Redis if data have to be persistent.  

Vertical Scaling Discussion:  
Envoy services can scale easily because they are stateless. Scaling of counter servers is more complex. Each counter server has its own Redis instance, so I would recommend using Redis partitioning approach, so every instance will only contain a subset of keys. It can be done by:  

1. Adding load balancing logic in the client (client side partitioning)  
Redis doesn't have to be partitioned, because the client directly select the server with the right Redis instance.
Envoy should implement hash partitioning (hashing load balancing) algorithm based on token value. This method enables envoy to repeatedly direct requests with the same token to the same server.  

2. Proxy load balancer  
Hash partitioning logic is on Redis site. Envoy services send requests to a proxy that is able to speak the Redis protocol, instead of sending requests directly to the right Redis instance. The proxy will make sure to forward request to the right Redis instance according to the configured partitioning schema.  
Cons: Proxies would have temporary copies of the RPC request and response, so they would need more resources to operate. Proxy model increases latency.  

3. External load balancer  
Client is asking external LB for server address.  
* this approach can work similarly to client site partitioning, but load balancing algorithm is implemented on external load balancer instead on client.  
* partitioning logic could be on Redis site. In this case external LB would speak to Redis instances, and return proper server address to the envoy when asked.  
