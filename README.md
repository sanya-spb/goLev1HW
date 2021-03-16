> Разработайте пакет для чтения конфигурации типичного веб-приложения через флаги или переменные окружения.
> Пример конфигурации можно посмотреть [здесь](https://gist.github.com/rumyantseva/26bee59a04d416c55e0e7e8155717d59). По желанию вы можете задать другие имена полям, сгруппировать их или добавить собственные поля.
> Помимо чтения конфигурации приложение также должно валидировать её - например, все URL’ы должны соответствовать ожидаемым форматам.
> Работу с конфигурацией необходимо вынести в отдельный пакет (не в пакет main).

#### REQUIREMENT
go get github.com/komkom/toml


#### RESULT
Вывод программы:
```json
$ SERVER_PORT=8181 MY_URL= go run main.go --debug
{
        "Debug": true,
        "My_url": "",
        "Database": {
                "Host": "127.0.0.1",
                "Port": 54321,
                "User": "test",
                "Pass": "pwd123",
                "Ssl": true
        },
        "Server": {
                "Bind": [
                        "10.0.0.1",
                        "127.0.0.1"
                ],
                "Port": 8181,
                "Log_level": 3
        }
}
Ok
```

В результате были применены default параметры, на них наложены параметры из конфиг-файла, далее поверх легли из ENV, и самые приоритетные из params, затем исправлен Log_level. 
Т.к. установлен Debug, то в stdout выведен "причесанный" JSON конфига.

валидацию URL сделал минимальную, т.к. думаю что основная проверка должна быть при подключении по этому URLу

#### BUGS

1. flag provided but not defined

Хотел сделать красиво: 
```go
result.Debug = *flag.Bool("debug", getEnvBool("DEBUG", result.Debug), "Output verbose debug information")
```
но параметр --debug в упор не ставится, пришлось разбить на 2 части:
```go
	result.Debug = getEnvBool("DEBUG", result.Debug)
	flag.BoolVar(&result.Debug, "debug", result.Debug, "Output verbose debug information")
```

#### TODO

1. как сделать взаимозаменяемую структуру в Go

```conf
<...>
# Database connection config
[database]
    # host = "127.0.0.1"# DB host (default: 127.0.0.1)
    port = 54321        # DB port (default: 5432)
    user = "test"       # DB user (required)
    pass = "pwd123"     # DB Port (required)
    ssl = true          # use ssl (optional)
<...>
```
либо
```conf
<...>
# Database connection config
[database]
    url = jdbc:postgresql://localhost/test?user=fred&password=secret&ssl=true
<...>
```

понятно, что можно в структуру добавить еще один string и далее что-нить придумать..
а как сделать, чтоб в зависимости от конфига - была либо та либо другая структура (знаний пока не хватает..)