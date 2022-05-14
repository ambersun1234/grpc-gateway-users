all: generate-users-proto

path = ./proto/users
target = ./proto/users.proto

generate-users-proto: $(target)
	@mkdir -p $(path)
	@protoc -I./proto --go_out=$(path) \
    	--go-grpc_out=$(path) \
		--grpc-gateway_out=$(path) \
		--grpc-gateway_opt grpc_api_configuration=./proto/users.yaml \
		$<

.PHONY: clean

clean:
	@rm -rf $(path)