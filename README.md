1. С какими интерфейсами мы уже сталкивались в предыдущих уроках? Обратите внимание на уроки, в которых мы читали из стандартного ввода и писали в стандартный вывод.

часто использовал такую конструкцию:
> reader := bufio.NewReader(os.Stdin)

на входе у нее:
```go
type Reader interface {
    Read(p []byte) (n int, err error)
}
```

для вывода:
> func Println(a ...interface{}) (n int, err error)

в пакете fmt во всех функциях параметры в виде интерфейсов ожидаются:
```go
func Errorf(format string, a ...interface{}) error
func Fprint(w io.Writer, a ...interface{}) (n int, err error)
func Fprintf(w io.Writer, format string, a ...interface{}) (n int, err error)
func Fprintln(w io.Writer, a ...interface{}) (n int, err error)
func Fscan(r io.Reader, a ...interface{}) (n int, err error)
func Fscanf(r io.Reader, format string, a ...interface{}) (n int, err error)
func Fscanln(r io.Reader, a ...interface{}) (n int, err error)
func Print(a ...interface{}) (n int, err error)
func Printf(format string, a ...interface{}) (n int, err error)
func Println(a ...interface{}) (n int, err error)
func Scan(a ...interface{}) (n int, err error)
func Scanf(format string, a ...interface{}) (n int, err error)
func Scanln(a ...interface{}) (n int, err error)
func Sprint(a ...interface{}) string
func Sprintf(format string, a ...interface{}) string
func Sprintln(a ...interface{}) string
func Sscan(str string, a ...interface{}) (n int, err error)
func Sscanf(str string, format string, a ...interface{}) (n int, err error)
func Sscanln(str string, a ...interface{}) (n int, err error)
```

в общем интерфейсы в Go это очень часто, там где можно сэкономить на дублировании кода, в голове еще не все улеглось, но похоже это взгляд на ООП чуть со стороны.

пустой интерфейс - очень удобно (привет Pyton'у с его динамической типизацией..):

```go
func main() {
	var i interface{}
	i = 42
	i = "hello"
}
```

очень изящно реализован пакет с сортировкой (sort), из-за абстрагирования от типов через интерфейсы - не надо дублировать код под свои типы, над лишь определить методы и сортировка из коробки!

ну и как же без самого распространенного интерфейса Error (10 строчек в исходниках, если убрать комментарии!)

```go
type error interface {
    Error() string
}
```

2. Посмотрите примеры кода в своём портфолио. Везде ли ошибки обрабатываются грамотно? Хотите ли вы переписать какие-либо функции? Проверьте себя: Вам должны быть знакомы следующие ключевые слова Go: interface. Вам должны быть знакомы следующие функции: panic, recover.

У меня в основном использовались бесконечные циклы, с уходом на новую итерацию в случае чего. Мне этот подход не нравится.
Все же надо обрабатывать ошибки, ловить и запускать исключения, и логичнее программа будет и не так топорно.