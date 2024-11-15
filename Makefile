.PHONY build_app:
build_app:
	go build -o ./cmd/bin/application ./cmd/app/main.go

.PHONY run_app:
run_app:
	./cmd/bin/application

.PHONY clean_app:
clean_app:
	rm -rf ./cmd/bin/

.PHONY build_migrator:
build_migrator:
	go build -o ./cmd/bin/migrator ./cmd/migrator/main.go

.PHONY run_migrator:
run_migrator:
	./cmd/bin/migrator -direction up

.PHONY test_database:
test_database:
	docker run -d --name door_locker_database --network=application_network -e POSTGRES_DB=door_locks -e POSTGRES_USER=user -e POSTGRES_PASSWORD=password -p 5432:5432 postgres;

.PHONY drop_test_database:
drop_test_database:
	docker stop test_db;
	docker rm test_db;

.PHONY all:
all:
	docker network create application_network;
	docker run -d --name door_locker_database --network=application_network -e POSTGRES_DB=door_locks -e POSTGRES_USER=user -e POSTGRES_PASSWORD=password -p 5432:5432 postgres;
	echo "Ждем развертывания базы данных";
	sleep 2;
	echo "Собираем мигратор и накатываем миграции на базу данных";
	go build -o ./cmd/bin/migrator ./cmd/migrator/main.go;
	./cmd/bin/migrator -direction up;
	echo "Начинаем сборку образа backend-приложнения";
	docker build -t backend:v0.0.1 .;
	echo "Запускаем backend-приложение";
	docker run --name backend_application --network application_network --env-file ./configs/.env -p 9090:9090 backend:v0.0.1;
