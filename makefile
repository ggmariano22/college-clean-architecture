CONTAINER_NAME=go-college

docker-install:
	make docker-build
	make docker-up

docker-build:
	docker-compose build

docker-up:
	docker-compose up

docker-logs:
	docker logs --follow ${CONTAINER_NAME}

docker-down:
	docker-compose down