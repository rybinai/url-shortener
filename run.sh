#!/bin/bash

case $1 in
    up) docker-compose up -d ;;
    down) docker-compose down ;;
    build) docker-compose up --build ;;
    create)
        go run cmd/client/main.go create "$2"
        read -p "Нажмите Enter"
        ;;
    get)
        go run cmd/client/main.go get "$2"
        read -p "Нажмите Enter"
        ;;
    check-db)
        docker-compose exec postgres psql -U user -d urlshortener \
            -c "SELECT * FROM urlshortener;"
        read -p "Нажмите Enter"
        ;;
    *)
        echo "Доступные команды"
        echo "./run.sh up - Запустить контейнер"
        echo "./run.sh down - Остановить контейнер"
        echo "./run.sh build - Пересобрать и запустить"
        echo "./run.sh create <URL> - Создать короткую ссылку"
        echo "./run.sh get <shortCode> - Получить оригинальную ссылку по коду"
        echo "./run.sh check-db - Показать все записи в базе"
        read -p "Нажмите Enter"
        ;;
esac