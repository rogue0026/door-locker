# команда для поднятия тестовой БД
docker run -d -e POSTGRES_DB=door_locker -e POSTGRES_USER=paul -e POSTGRES_PASSWORD=dfcz123 -p 5432:5432 postgres