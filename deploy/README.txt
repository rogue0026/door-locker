КОМАНДЫ ВВОДИТЬ БЕЗ КАВЫЧЕК!!!!
Создаем новый контейнер PostgreSQL, убедитесь, что порт 5432 хостовой машины не используется каким-либо приложением
    "docker run -d --name door_locker_database_container -e POSTGRES_DB=door_locker -e POSTGRES_USER=paul -e POSTGRES_PASSWORD=password -p 5432:5432 postgres"
# Копируем в созданный контейнер файл, содержащий скрипты на языке SQL
    "docker cp ./deploy/queries.sql door_locker_database_container:/app_utils"
Теперь необходимо подключиться к командной оболочке контейнера
Для этого выполните команду:
    "docker exec -it door_locker_database_container psql -d door_locker -U paul -W"
После выполнения вышеуказанной команды будет предложено ввести пароль. Введите password
Если все сделано правильно, то будет запущена утилита psql для работы с сервером БД PostgreSQL, приглашение к вводу должно выглядеть вот так: "door_locker=#"
Далее, выполнить следующую команду: "\i /app_utils"
В результате будут выполнены скрипты из файла queries.sql, который лежит в директории deploy
Тестовая БД будет готова наполнена данными и готова к работе


