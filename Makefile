.PHONY build_migrator:
build_migrator:
	go build -o ./cmd/bin/migrator ./cmd/migrator/main.go

.PHONY run_migrator:
run_migrator:
	./cmd/bin/migrator -direction up

.PHONY test_database:
test_database:
	docker run -d --name door_locker_database --network=application_network -e POSTGRES_DB=door_locks -e POSTGRES_USER=user -e POSTGRES_PASSWORD=password -p 5432:5432 postgres;

.PHONY deploy_in_docker:
deploy_in_docker:
	docker network create application_network;
	docker run -d --name door_locker_database --rm --network=application_network -e POSTGRES_DB=door_locks -e POSTGRES_USER=user -e POSTGRES_PASSWORD=password -p 5432:5432 postgres;
	sleep 3;
	go build -o ./cmd/bin/migrator ./cmd/migrator/main.go;
	./cmd/bin/migrator -direction up;
	docker build -t backend:v0.0.1 .;
	docker run --name backend_application --rm --network application_network --env-file ./configs/.env -p 9090:9090 backend:v0.0.1;

.PHONY deploy_local:
deploy_local:
	docker run -d --name door_locker_database --rm -e POSTGRES_DB=door_locks -e POSTGRES_USER=user -e POSTGRES_PASSWORD=password -p 5432:5432 postgres;
	sleep 2;
	go build -o ./cmd/bin/migrator ./cmd/migrator/main.go;
	./cmd/bin/migrator -direction up;
	go build -o ./cmd/bin/application ./cmd/app/main.go;
	APP_ENVIRONMENT=development HTTP_SERVER_HOST=localhost HTTP_SERVER_PORT=9090 DB_HOST=localhost DB_PORT=5432 DB_USER=user DB_USER_PASSWORD=password DATABASE_NAME=door_locksAPP_ENVIRONMENT=development HTTP_SERVER_HOST=localhost HTTP_SERVER_PORT=9090 DB_HOST=localhost DB_PORT=5432 DB_USER=user DB_USER_PASSWORD=password DATABASE_NAME=door_locks ./cmd/bin/application;