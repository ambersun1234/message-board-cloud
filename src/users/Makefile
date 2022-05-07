all: generate-users-proto generate-users-proto-gateway

path = ./proto/users
target = ./proto/users.proto
target_yaml = ./proto/users.yaml

generate-users-proto: $(target)
	@mkdir -p $(path)
	@protoc -I./proto --go_out=plugins=grpc:$(path) \
		$<

generate-users-proto-gateway: $(target) $(target_yaml)
	@mkdir -p $(path)
	@protoc -I./proto --grpc-gateway_out $(path) \
			--grpc-gateway_opt logtostderr=true \
			--grpc-gateway_opt grpc_api_configuration=$(target_yaml) \
			$(target)

.PHONY: clean

clean:
	@rm -rf $(path)