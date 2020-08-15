# fiber-example

https://gofiber.io/

## Run

```bash
$ make dev
```

## API

- Post

```bash
curl --request POST 'localhost:3000/api/v1/todo' \
--header 'Content-Type: application/json' \
--data-raw '{
    "id": 1,
    "name": "Todo 1"
}'
```

- Get

```bash
curl --request GET 'localhost:3000/api/v1/todo/1'
```

- Put

```bash
curl --request PUT 'localhost:3000/api/v1/todo/69' \
--header 'Content-Type: application/json' \
--data-raw '{
    "name": "Todo 1"
}'
```

- Patch

```bash
curl --request PATCH 'localhost:3000/api/v1/todo/96' \
--header 'Content-Type: application/json' \
--data-raw '{
    "name": "Todo 1"
}'
```

- Delete

```bash
curl --request DELETE 'localhost:3000/api/v1/todo/99' \
--header 'Content-Type: application/json' \
--data-raw '{
    "name": "Todo 1"
}'
```
