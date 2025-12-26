# URL Shortener

Сервис для сокращения URL-адресов на основе gRPC и PostgreSQL

## Архитектура

- **gRPC сервер** на порту 8080
- **PostgreSQL** база данных на порту 5432
- **Docker Compose** для запуска всех компонентов

## Требования к запуску 

- Docker и Docker Compose
- Go 1.25.0

## Установка и запуск

### Запуск сервера
```bash
# Запустить контейнеры
./run.sh up

# Пересобрать и запустить
./run.sh build

# Остановить контейнеры
./run.sh down
```

## Использование

### Создание короткой ссылки
```bash
./run.sh create 
```

Пример:
```bash
./run.sh create https://example.com
```

### Получение оригинальной ссылки
```bash
./run.sh get 
```

Пример:
```bash
./run.sh get aBcDeF
```

### Проверка содержимого базы данных
```bash
./run.sh check-db
```