Для развертывания приложения выполнить следующие действия:
1. Поднять docker-контейнер с сервером PostgreSQL:
    docker run -d --name door_locker_database_container -e POSTGRES_DB=door_locker -e POSTGRES_USER=paul -e POSTGRES_PASSWORD=password -p 5432:5432 postgres
Убедись, что порт 5432 хостовой машины не используется другим приложением. Если используется, то поменять номер порта слева (-p 5432:5432) на любой другой.
2. Находясь в директории проекта выполни по очереди команды:
    make build_migrator;
    make run_migrator;
Эти команды создадут в тестовой базе данных необходимые таблицы, настроит связи между таблицами, добавит функции и пользовательские процедуры для правильной работы приложения.