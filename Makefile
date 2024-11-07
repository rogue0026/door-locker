.PHONY build:
build:
	go build -o ./cmd/bin/application ./cmd/app/main.go

.PHONY run:
run:
	./cmd/bin/application -cfg ./configs/local.yaml

.PHONY all:
all: build run

.PHONY clean:
clean:
	rm -rf ./cmd/bin/