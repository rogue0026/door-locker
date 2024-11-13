.PHONY build_app:
build:
	go build -o ./cmd/bin/application ./cmd/app/main.go

.PHONY run_app:
run:
	./cmd/bin/application -cfg ./configs/local.yaml

.PHONY clean_app:
clean:
	rm -rf ./cmd/bin/

.PHONY build_migrator:
build_migrator:
	go build -o ./cmd/bin/migrator ./cmd/migrator/main.go

.PHONY run_migrator:
run_migrator:
	./cmd/bin/migrator -cfg ./configs/local.yaml -direction down