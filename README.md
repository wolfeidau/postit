# postit [![Build Status](https://travis-ci.org/wolfeidau/postit.svg?branch=master)](https://travis-ci.org/wolfeidau/postit)

This project illustrates how you can use https://github.com/twitchtv/twirp running in lambda or a docker container. It leans on the great work done in https://github.com/apex/gateway to provide the http server abstraction in lambda therefore enabling deployment of generic golang web servers.

# building

Testing lambda locally.

* Install all our dependencies.

```
make setup
```

* Run local test environment for lambda.

```
make local-lambda
```

* Connect using the client and hit the service.

```
go run cmd/postit/main.go http://localhost:3000
```

# licence

Copyright 2018 Mark Wolfe and licensed under the Apache License, Version 2.0