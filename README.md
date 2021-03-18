## Урок 9. Работа с файловой системой

1. К приложению из практической части предыдущего урока добавьте возможность читать данные из файлов. Конфигурация может быть задана в форматах yaml или json. Также по желанию вы можете добавить и другие форматы.
2. Помимо чтения конфигурации приложение также должно валидировать её - например, все URL’ы должны соответствовать ожидаемым форматам.
3. Работу с конфигурацией необходимо вынести в отдельный пакет (не в пакет main).

### RESULT
Вывод программы:
```bash
$ git tag
v1.0.0
$ make
GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build \
        -ldflags "-s -w -X github.com/sanya-spb/goLev1HW/utils/version.version=v1.0.0 \
        -X github.com/sanya-spb/goLev1HW/utils/version.commit=git-3c9bb0b \
        -X github.com/sanya-spb/goLev1HW/utils/version.buildTime=2021-03-18_14:12:01 \
        -X github.com/sanya-spb/goLev1HW/utils/version.copyright="sanya-spb"" \
        -o app_main 
$ ./app_main --config=config.yaml --debug
{
        "debug": true,
        "my_url": "https://google.com",
        "database": {
                "host": "127.0.0.1",
                "port": 54321,
                "user": "test",
                "pass": "pwd123",
                "ssl": true
        },
        "server": {
                "bind": [
                        "10.0.0.1",
                        "127.0.0.1"
                ],
                "port": 8888,
                "log_level": 3
        }
}
version: {Version:v1.0.0 Commit:git-3c9bb0b BuildTime:2021-03-18_14:12:01 Copyright:sanya-spb}
```

### TODO

1. как сделать взаимозаменяемую структуру в Go

```yaml
#<...>
# Database connection config
database:
    host: "127.0.0.1"
    port: 54321
    user: "test"
    pass: "pwd123"
    ssl: true
#<...>
```
либо
```yaml
#<...>
# Database connection config
database:
    url: jdbc:postgresql://localhost/test?user=fred&password=secret&ssl=true
#<...>
```

понятно, что можно в структуру добавить еще один string и далее что-нить придумать..

а как сделать, чтоб в зависимости от конфига - была либо та либо другая структура (знаний пока не хватает..)