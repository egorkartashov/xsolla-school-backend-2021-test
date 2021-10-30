# GraphQL API

## GET /products

Ответы на все запросы возвращаются в следующем формате:
```json
{
  "data": {...}, //результат выполнения запроса
  "errors": [
    //ошибки выполнения запроса
  ]
}
```

> :memo: Cимволом `*` помечены обязательные аргументы, при запросе их указывать не надо

### Типы данных

Тип product со следующими полями:
* `id` (uuid в виде строки) - идентификатор продукта
* `sku` (string) - SKU продукта
* `name` (string) - название продукта
* `type` (string) - тип продукта
* `priceInCents` (int) - стоимость продукта в центах

Примеры запросов:

Пример №1

Запрос: `GET {{server-url}}/api/products?query={product(sku: "123"){id,sku,name,type,priceInCents}`

Ответ:
```json
{
  "data": {
    "productBySku": {
      "id": "29a7e62d-2948-45b4-8e38-1c0e6c73d520",
      "name": "Hello world",
      "priceInCents": 0,
      "sku": "123",
      "type": "Message"
    }
  }
}
```

Пример №2

Запрос: `GET {{server-url}}/api/products?query=mutation+_{updateBySku(sku: "123",name="Hello, World!!!", priceInCents: 10){sku,name,type,priceInCents}`

Ответ:
```json
{
  "data": {
    "updateBySku": {
      "name": "Hello, World!!!",
      "priceInCents": 10,
      "sku": "123",
      "type": "Message"
    }
  }
}
```


### Запросы на получение данных (Query)

`product(id*: ""){...}`
Получить продукт по его идентификатору

`productBySku(sku*: ""){...}`
Получить продукт по его SKU

`productsList(offset: 0, limit: 10){...}`
Получить список продуктов. 
`offset` - сколько первых значений нужно пропустить, `limit` - сколько значений максимум можно вернуть

### Запросы на изменение данных (Mutation)

`create(sku*: "", name*: "", type*: "", priceInCents*: 0){...}`
Получить продукт по его идентификатору. Возвращает созданный продукт, включая его ID

`update(id*: "", sku: "", name: "", type: "", priceInCents: 0){...}` 
Изменить продукт по его ID. Изменяет только указанные поля и возвращает измененный продукт

`updateBySku(sku*: "", name: "", type: "", priceInCents: 0){...}`
Обновить продукт по его SKU. Изменяет только указанные поля и возвращает измененный продукт

`delete(id*: "")`
Удалить продукт по его ID. Возвращает сообщение "deleted product with ID=<ID продукта>"

`deleteBySku(sku*: "")`
Удалить проудкт по его SKU. Возвращает сообщение "deleted product with SKU=<SKU продукта>"

