.PHONY build_app:
build_app:
	go build -o ./cmd/bin/application ./cmd/app/main.go

.PHONY run_app:
run_app:
	./cmd/bin/application -cfg ./configs/local.yaml

.PHONY clean_app:
clean_app:
	rm -rf ./cmd/bin/

.PHONY build_migrator:
build_migrator:
	go build -o ./cmd/bin/migrator ./cmd/migrator/main.go

.PHONY run_migrator:
run_migrator:
	./cmd/bin/migrator -cfg ./configs/local.yaml -direction up

.PHONY test_database:
test_database:
	docker run -d --name test_db -e POSTGRES_DB=door_locker -e POSTGRES_USER=paul -e POSTGRES_PASSWORD=password -p 5432:5432 postgres;

.PHONY drop_test_database:
drop_test_database:
	docker stop test_db;
	docker rm test_db;

.PHONY all:
all: test_database
	echo "Ждем пока развернется контейнер";
	sleep 2;
	go build -o ./cmd/bin/migrator ./cmd/migrator/main.go;
	./cmd/bin/migrator -cfg ./configs/local.yaml -direction up;
	go build -o ./cmd/bin/application ./cmd/app/main.go;
	./cmd/bin/application -cfg ./configs/local.yaml;