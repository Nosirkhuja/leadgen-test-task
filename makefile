# Makefile для управления проектом Building Service

# Запуск PostgreSQL в Docker
postgres:
	docker-compose up -d

# Остановка PostgreSQL
down:
	docker-compose down

# Сборка приложения
build:
	go build -o building-service cmd/main.go

# Запуск приложения локально
run:
	go run cmd/main.go

# Генерация Swagger документации
swagger:
	swag init -g cmd/main.go

# Запуск тестов
test:
	go test ./...

# Полная очистка (останавливает Docker, удаляет собранные файлы)
clean:
	docker-compose down
	rm -f building-service
	go clean

# Пересоздание базы данных
recreate-db: down
	docker-compose up -d

.PHONY: postgres down build run swagger test clean recreate-db
