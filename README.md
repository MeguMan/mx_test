## Тестовое задание для стажера в юнит Merchant Experience
### Quick start:
    docker-compose up   
### API methods:  
#### /offers POST
Path это путь к файлу на яндекс диске, по умолчанию в конфиге указан 
мой токен, получить свой можно по 
ссылке https://yandex.ru/dev/disk/poligon/.
Данные в файле должны обязательно находиться на первом листе.

В ответе мы получаем id по которому мы можем узнать состояние работы над нашим запросом.

Request body example:
```json
{
    "seller_id": 23,
    "path": "avito-test/table.xlsx"
}
```
Request response example:
```json
{
    "key": "655ec1d62bbc49399397ba13ee78f3b0"
}
```

#### /offers GET
Пустой pattern соответствует всем значениям, по паттерну "теле" находятся и "телефоны", и "телевизоры".
Если seller_id или offer_id не были указаны, то тогда будут возвращены соответственно все значения.


Request body example:
```json
{
    "seller_id": 23,
    "pattern": "моло"
}
```
Request response example:
```json
[
    {
        "offer_id": 6,
        "name": "молоко",
        "price": 20,
        "quantity": 123,
        "seller_id": 23
    },
    {
        "offer_id": 7,
        "name": "молоко вкусное",
        "price": 123,
        "quantity": 234,
        "seller_id": 23
    }
]
```

#### /offers/status/{id} GET
Показывает статус запроса по id, который мы получили из POST запроса.
Все статусы хранятся в памяти.

Request response example:
```json
task is finished
{"CreatedRows":4,"UpdatedRows":0,"DeletedRows":1,"ErrorRows":3}
```
