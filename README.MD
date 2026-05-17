## 📦 Структура таблицы `orders`

| Поле         | Тип        | Описание                     |
|--------------|------------|------------------------------|
| id           | int        | Уникальный ID заказа         |
| customer     | text       | Имя клиента                  |
| address      | text       | Адрес доставки               |
| is_delivered | boolean    | Доставлен ли заказ           |
| created_at   | timestamp  | Дата создания заказа         |

> Все поля обязательны. По умолчанию заказ считается недоставленным.

---

## 🧱 Как применить миграцию вручную

1. Убедись, что у тебя запущен PostgreSQL
2. В терминале введи:

```bash
psql -U postgres -d gotracker -f migrations/001_create_orders.sql
```

## 🚀 Доступные API эндпоинты

- `GET /orders` — получить список заказов
- `GET /orders/{id}` — получить заказ по ID
- `POST /orders` — создать заказ
- `PUT /orders/{id}` — обновить заказ
- `DELETE /orders/{id}` — удалить заказ

## 🌱 Переменные окружения

- `DATABASE_URL` — PostgreSQL DSN
- `PORT` — порт для HTTP-сервера (по умолчанию `8080`)
- `REDIS_ADDR` — адрес Redis
- `REDIS_TTL` — время жизни кеша в секундах
- `KAFKA_BROKER` — адрес Kafka-брокера
- `APP_ENV` — среда выполнения

## 🐳 Запуск через Docker Compose

1. Скопируйте `.env.example` в `.env` и заполните значения.
2. Постройте и запустите сервисы:

```bash
docker-compose up --build
```

3. API будет доступен на `http://localhost:8080`.

Если нужно только запустить бэкенд без Kafka, закомментируйте сервисы `zookeeper` и `kafka` в `docker-compose.yml`.

