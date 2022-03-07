# e-wallet

## Установка
* `go mod tidy`
* Мигрировать все *.up.sql из папки /migrations
* Запустить seeder для добавление пользователей: `Post /seed-users`

## Использование
### Исполнимый файл
`/cmd/main.go`

### Методы
Проверить существует ли аккаунт электронного кошелька.
* `Post /wallet/check`:
```
header:"X-UserId" type:"int"    validate:"required,gt=0"
header:"X-Digest" type:"string" validate:"required"
json:"email"      type:"string" validate:"required,email"
```

Пополнение электронного кошелька.
* `Post /wallet/new-transaction`:
```
header:"X-UserId" type:"int"    validate:"required,gt=0"
header:"X-Digest" type:"string" validate:"required"
json:"amount"     type:"string" validate:"required,gt=0
```

Получить общее количество и суммы операций пополнения за текущий месяц.
* `Post /wallet/monthly-transaction-info`:
```
header:"X-UserId" type:"int"    validate:"required,gt=0"
header:"X-Digest" type:"string" validate:"required"
```

Получить баланс электронного кошелька.
* `Post /wallet/balance`:
```
header:"X-UserId" type:"int"    validate:"required,gt=0"
header:"X-Digest" type:"string" validate:"required"
```

- X-UserId - ID пользователя в БД
- X-Digest - hmac-sha1 хеш с ключом Key пользователя в БД

### Вспомогательные методы
Получить строку X-Digest для пользователя с идентификатором X-UserId по его ключу Key
* `Post /get-digest`:
```
header:"X-UserId" type:"int"  validate:"required,gt=0"
```
