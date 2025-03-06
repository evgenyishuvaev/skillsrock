include .env

run:
	sudo docker compose up -d --build

stop:
	sudo docker compose stop


migrate:
	sudo docker exec -it ${APP_CONTAINER_NAME} goose up

