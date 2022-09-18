# Сервис мониторинга продуктов

gRPC-сервер с постоянным хранилищем MongoDB, реализующий 2 метода:

- Fetch(URL) - запросить внешний CSV-файл со списком продуктов по внешнему адресу.
CSV-файл имеет вид PRODUCT NAME;PRICE. Последняя цена каждого продукта сохраняется 
в базе с датой запроса. Также сохраняется количество изменений цены продукта.

- List(paging params, sorting params) - получить постраничный список продуктов (цена, количество изменений цены, дата последнего обновления) с сортировкой по выбранному полю.


### Стэк
- go 1.18
- mongo 

### Запуск
Для запуска необходимо указать переменные в конфиг-файле configs/config.yml

```
DB:
    uri: "mongodb://localhost:27017"
    login: "admin" 
    password: "admin"
    database: "monitor_of_the_products"

server:
    Port: 9000
```

Сборка и запуск
```
go build -o app cmd/server/server.go && ./app
```

Для mongo можно использовать Docker

```
docker run --rm -d --name monitor_of_the_products \
 -e MONGO_INITDB_ROOT_USERNAME=admin \
 -e MONGO_INITDB_ROOT_PASSWORD=admin \
 -p 27017:27017 mongo:latest
```

MongoDB Compass
```
mongodb://admin:admin@localhost:27017
```

Update proto:
```
protoc --go_out=. --go-grpc_out=. proto/grpc_product.proto
```
