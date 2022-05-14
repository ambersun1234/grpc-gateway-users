# grpc-gateway-users
This project shows the simple usage of [gRPC](https://grpc.io/) and [grpc-gateway](https://grpc-ecosystem.github.io/grpc-gateway/)

# Run
```
$ go run server.go
```

# Test
You can use [curl](https://curl.se/) to perform simple HTTP/1.1 request to the target server\
First install curl
```shell
$ sudo apt install curl -y
```

Then send a request
```shell
$ curl http://localhost:7777/api/users/1
{"user_id":"1","user_name":"test_user_name","first_name":"test_first_name","last_name":"test_last_name","email":"test@test.com"}
$ 
```

# License
This project is licensed under MIT License, see the [LICENSE](./LICENSE) file for more detail
