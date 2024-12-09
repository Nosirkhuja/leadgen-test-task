# Building Service

## Описание
Сервис управления информацией о строениях с использованием Golang, PostgreSQL и Gin Framework.

## Требования
- Go 1.21+
- Docker
- Docker Compose
- Make

## Быстрый старт

### Запуск PostgreSQL
```bash
make postgres
```

### Генерация Swagger документации
```bash
make swagger
```

### Сборка приложения
```bash
make build
```

### Запуск приложения локально
```bash
make run
```

### Запуск тестов
```bash
make test
```

## Управление базой данных
- Запуск базы данных: `make postgres`
- Остановка базы данных: `make down`
- Пересоздание базы данных: `make recreate-db`

## Конфигурация
Настройки подключения к базе данных находятся в `config/config.yaml`:
```yaml
database:
  host: localhost
  port: 5432
  user: buildinguser
  password: buildingpass
  name: buildingdb
```

## Эндпоинты
- `POST /api/v1/buildings`: Создание строения
- `GET /api/v1/buildings`: Получение списка строений с фильтрацией
- `/docs/index.html`: Swagger документация

## Примеры запросов

### Создание строения
```bash
curl -X POST http://localhost:8080/api/v1/buildings \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Офисное здание",
    "city": "Москва",
    "year": 2020,
    "floor_count": 10
  }'
```

### Получение списка строений
```bash
curl "http://localhost:8080/api/v1/buildings?city=Москва&year=2020"
```