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
	docker run -d --name door_locker_database --network=door_locker_network -e POSTGRES_DB=door_locks -e POSTGRES_USER=user -e POSTGRES_PASSWORD=password -p 5432:5432 postgres;

.PHONY drop_test_database:
drop_test_database:
	docker stop test_db;
	docker rm test_db;



.PHONY all:
all: test_database
	echo "Ждем пока развернется контейнер с сервером БД";
	sleep 2;
	go build -o ./cmd/bin/migrator ./cmd/migrator/main.go;
	./cmd/bin/migrator -direction up;
	go build -o ./cmd/bin/application ./cmd/app/main.go;
	./cmd/bin/application -cfg ./configs/local.yaml;


.PHONY build_image:
build_image:
	docker build -t door-locker-img:v.0.0.1 .;