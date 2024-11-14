# Preparing
# =======================================================
FROM golang:1.22.5-alpine3.20 AS base

RUN mkdir /door-locker

COPY . /door-locker

WORKDIR /door-locker

RUN go build -o /door-locker/application ./cmd/app/main.go

# Production
# =======================================================
FROM scratch AS prod

COPY --from=base /door-locker/application /door-locker/application

EXPOSE 9090

CMD ["./door-locker/application"]