include .env

run:
	sudo docker compose up -d --build

migrate:
	sudo docker exec -it ${APP_CONTAINER_NAME} goose up

