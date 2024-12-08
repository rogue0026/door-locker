# Stage 1: Build
FROM golang:1.22.5-bullseye AS builder

# Установка рабочей директории внутри контейнера
WORKDIR /app

# Копируем go.mod и go.sum файлы
COPY go.mod go.sum ./

# Устанавливаем зависимости
RUN go mod download

# Копируем весь код в рабочую директорию
COPY . .

# Сборка приложения
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app ./cmd/app/main.go

# Stage 2: Run
FROM alpine:latest

# Установка сертификатов (необходимо для HTTPS-запросов)
# RUN apk --no-cache add ca-certificates

# Установка рабочей директории внутри контейнера
WORKDIR /root/

# Копируем скомпилированное приложение из стадии сборки
COPY --from=builder /app/app .

# Указываем порт, который будет использовать приложение
EXPOSE 9090

# Команда для запуска приложения
CMD ["./app"]
