# Tasker
Сервис позволяет создавать задачи с разными статусами
# Как запусить?
Для старта потребуется установленный `docker compose`
1. `.env` - необходимо создать файл `.env` в корне проекта, на основе примера `.env.example`
2. `make run` - запускает процесс сборки образа и запуска контейнеров
3. `make migrate` - при первом старте так же необходимо провести миграции
4. `make stop` - остонавливает работу приложения
