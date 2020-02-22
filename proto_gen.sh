protoc --go_out=plugins=grpc,paths=source_relative:./proto-go/proto-entity-reputation \
--proto_path=proto/entity/reputation/ \
proto/entity/reputation/*.proto