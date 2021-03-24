## Урок 10. Основы тестирования Go-приложений

1. Выберите три любых приложения, написанных в рамках курса, и добавьте к ним тесты. Обратите внимание: тестируемый код должен быть вынесен из пакета main в отдельный пакет или пакеты. Среды выбранных приложений обязательно должно быть хотя бы одно, в котором реализовано два разных варианта одного и того же алгоритма (например, в задаче на сортировку слайсов). Для таких приложений добавьте бенчмарки. Также хотя бы для одного приложения тестирование должно производиться на основе табличных тестов. Добавьте хотя бы один пример использования тестируемых функций с помощью механизма example.
2. Познакомьтесь подробнее с библиотекой testify и попробуйте написание тестов с её помощью. Сравните этот способ написания тестов с вариантом, когда вы пишите тесты на чистом Go без testify. Какой подход вам нравится больше и почему?

### 1.
Вывод программы:
```bash
$ make
GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build \
        -ldflags "-s -w -X github.com/sanya-spb/goLev1HW/utils/version.version=v1.0.0 \
        -X github.com/sanya-spb/goLev1HW/utils/version.commit=git-38904cb \
        -X github.com/sanya-spb/goLev1HW/utils/version.buildTime=2021-03-24_10:27:41 \
        -X github.com/sanya-spb/goLev1HW/utils/version.copyright="sanya-spb"" \
        -o app_main main.go
$ make test
go test -v github.com/sanya-spb/goLev1HW/utils/config/
=== RUN   TestIsIPv4Net
--- PASS: TestIsIPv4Net (0.00s)
=== RUN   TestIsURL
--- PASS: TestIsURL (0.00s)
=== RUN   TestTestConfig
--- PASS: TestTestConfig (0.00s)
=== RUN   ExampleIsIPv4Net_true
--- PASS: ExampleIsIPv4Net_true (0.00s)
=== RUN   ExampleIsIPv4Net_false
--- PASS: ExampleIsIPv4Net_false (0.00s)
=== RUN   ExampleIsURL_true
--- PASS: ExampleIsURL_true (0.00s)
=== RUN   ExampleIsURL_false
--- PASS: ExampleIsURL_false (0.00s)
PASS
ok      github.com/sanya-spb/goLev1HW/utils/config      (cached)
go test -v github.com/sanya-spb/goLev1HW/fibonacci/
=== RUN   TestFibonacciR
--- PASS: TestFibonacciR (2.15s)
=== RUN   TestFibonacciM
--- PASS: TestFibonacciM (0.00s)
=== RUN   ExampleFibonacciR
--- PASS: ExampleFibonacciR (0.00s)
=== RUN   ExampleFibonacciM
--- PASS: ExampleFibonacciM (0.00s)
PASS
ok      github.com/sanya-spb/goLev1HW/fibonacci (cached)
```

### 2.
testify - серьезный подход к тестированию, в чем-то проще чем на чистом Go, но это смотря как относиться к тестированию.

Если использовать упрощенный подход, то чистого Go более чем достаточно. Ну а при серьезном подходе к тестированию, думаю в перспективе testify выигрывает: меньше сил тратится на написание более тщательных тестов.
