## Url-shortener

### Запуск базы данных
docker-compose up -d postgres

### Остановка базы данных
docker-compose down

### Запуск сервера 
go run cmd/server/main.go

### Создать короткую ссылку
go run cmd/client/main.go create https://google.com

### Получить оригинальный URL
go run cmd/client/main.go get abc123

### Просмотр всех данных в БД
docker-compose exec postgres psql -U user -d urlshortener -c "SELECT * FROM urlshortener;"

#### Порт сервера 
8080

#### Порт БД
5432