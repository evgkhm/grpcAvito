# Microservice for working with user balance

# Stack
+ gRPC
+ PostgreSQL
+ golang-migrate/migrate
+ sqlx
+ logrus
+ Docker

# Getting Started
1. `git clone https://github.com/evgkhm/grpcAvito`
2. `cd grpcAvito`
3. `docker-compose up --build`

# Implemented
1. Метод начисления средств на баланс
2. Метод резервирования средств
3. Метод признания выручки
4. Метод получения баланса пользователя
5. Доп. задания:
 - Доп. задание 1 (метод для получения месячного отчета) в csv файл
 - Сценарий разрезервирования денег

# API
| Endpoint               | gRPC method |                          Description |
|------------------------|:-----------:|-------------------------------------:|
| /get_balance/{id}      |     GET     | Метод получения баланса пользователя |
| /create                |    POST     |          Метод создания пользователя |
| /sum                   |    POST     |   Метод начисления средств на баланс |
| /reservation           |    POST     |         Метод резервирования средств |
| /revenue               |    POST     |              Метод признания выручки |
| /delete_reservation    |    POST     |               Метод удаления резерва |
| /report/{year}/{month} |     GET     |          Получение месячного отсчета |

# Example
Получение текущего баланса
![img.png](img.png)

# For database connecting
1. Из командной строки выполнить docker ps
2. Найти CONTAINER ID
3. Выполнить docker exec -it CONTAINER ID /bin/bash
4. Подключение к БД psql -U admin -d users

# For proto-files generating
protoc -I ./proto --go_out ./proto --go_opt paths=source_relative --go-grpc_out ./proto --go-grpc_opt paths=source_relative --grpc-gateway_out ./proto --grpc-gateway_opt paths=source_relative ./proto/server.proto