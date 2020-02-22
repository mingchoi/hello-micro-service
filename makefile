default:
	echo "Hello world"

build_service_entity_reputation:
	docker build -t mingchoi/service-entity-reputation -f service-entity-reputation/Dockerfile ./


build_service_api_reputation:
	docker build -t mingchoi/service-api-reputation -f service-api-reputation/Dockerfile ./

up:
	docker-compose up

all:
	make build_service_entity_reputation
	make build_service_api_reputation
	make up
