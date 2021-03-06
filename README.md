# postit [![Build Status](https://travis-ci.org/wolfeidau/postit.svg?branch=master)](https://travis-ci.org/wolfeidau/postit)

This project illustrates how you can use https://github.com/twitchtv/twirp running in lambda. It leans on the great work done in https://github.com/apex/gateway to provide the http server abstraction in lambda therefore enabling deployment of generic golang web servers.

# why?

The reason I have combined these projects together is to build a robust RPC service which can be deployed on lambda and provide much needed visibility of the enclosed RPC calls.

# building

Testing lambda locally.

Install all our dependencies.

```
make setup
```

Run local test environment for lambda.

```
make local-lambda
```

Connect using the client and hit the service.

```
go run cmd/postit/main.go http://localhost:3000
```

# tracing

To enable tracing of RPC requests there is also integration to AWS xray.

![Demo of xray](doc/images/xraydemo.png)

# todo

- [ ] Add RPC meta data to the xray trace
- [ ] Add a dynamodb or S3 call to persist and retrieve posts
- [ ] Add middleware to support including cognito identity information from lambda

# licence

Copyright 2018 Mark Wolfe and licensed under the Apache License, Version 2.0