// Доработать калькулятор: больше операций и валидации данных.

package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// Разбор строки ввода от пользователя
func getExpr(msg string) (string, string, string) {
	reader := bufio.NewReader(os.Stdin)
	// соберем регулярочку.. люблю я их..
	regexpExpr := regexp.MustCompile(`^\s*([\+\-]?\d+[\.\,]?\d*)\s*([+\-*\/%]|\/\/|concat|rnd)\s*([\+\-]?\d+[\.\,]?\d*)\s*$`)
	// попросим ввести математическое выражение (настойчиво..)
	for {
		fmt.Print(msg)
		str1, _ := reader.ReadString('\n')
		if parsedStr := regexpExpr.FindAllStringSubmatch(str1, -1); parsedStr != nil {
			// log.Printf("%q\n", parsedStr) //debug
			return parsedStr[0][1], parsedStr[0][3], parsedStr[0][2]
		}
		fmt.Println("Что-то не так с выражением, попробуйте еще раз..")
	}
}

// собственно, строку в число + сколько знаков у нас после запятой
func str2float(str string) (num float64, frac int) {
	str = strings.Replace(str, ",", ".", 1)
	num, _ = strconv.ParseFloat(str, 64)
	// log.Printf("len=%d; index=%d\n", len(str), strings.Index(str, ".")) //debug
	if indx := strings.Index(str, "."); indx != -1 {
		return num, len(str) - indx - 1
	} else {
		return num, 0
	}
}

// угадай-ка кто больше..
func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func myCalc(op1 string, op2 string, op string) {
	a, a_frac := str2float(op1)
	b, b_frac := str2float(op2)
	// fmt.Printf("a=%f; frac=%d\n", a, a_frac) //debug
	// fmt.Printf("b=%f; frac=%d\n", b, b_frac) //debug
	// fmt.Printf("op=%s\n", op) //debug

	switch op {
	case "+": // разминка, точность берем из наиболее "неточного" числа
		fmt.Printf("сложение, точность: %d зн.посл.зап.\nрезультат: %0."+fmt.Sprint(Max(a_frac, b_frac))+"f\n", Max(a_frac, b_frac), a+b)
	case "-": // тоже самое при вычитании
		fmt.Printf("вычитание, точность: %d зн.посл.зап.\nрезультат: %0."+fmt.Sprint(Max(a_frac, b_frac))+"f\n", Max(a_frac, b_frac), a-b)
	case "*": // при умножении "неточность" увеличивается
		fmt.Printf("умножение, точность: %d зн.посл.зап.\nрезультат: %0."+fmt.Sprint(a_frac+b_frac)+"f\n", a_frac+b_frac, a*b)
	case "/": // какое деление?.. покажем оба варианта, если числа целые и только один вариант, если "не очень целые"
		if a_frac+b_frac == 0 {
			fmt.Printf("целочисленное деление, точность: %d зн.посл.зап.\nрезультат: %d\n", 0, int(a/b))
		}
		fmt.Printf("обычное деление, точность: максимальная\nрезультат: %f\n", a/b)
	case "%": // тут будет результат только если числа целые
		fmt.Println("остаток от деления")
		if a_frac+b_frac != 0 {
			fmt.Println("результат можно посчитать только для целых чисел!")
			break
		}
		fmt.Printf("результат: %d\n", int(a)%int(b))
	case "//": // двухсимвольный оператор, и деление с приведением типов до вычисления
		fmt.Printf("целочисленное деление\nрезультат: %d\n", int(a)/int(b))
	case "concat": // склеим.. но от первого числа возьмем только целую часть.
		fmt.Printf("конкатенация\nрезультат: %s\n", fmt.Sprintf("%d%0."+fmt.Sprint(b_frac)+"f", int(a), b))
	default: // ветки default по идее у нас быть не может, т.к. регулярка не пропустит ничего лишнего
		// но тут будет 1 случай с оператором "rnd", заодно проверим редактор на генерацию отступов..
		fmt.Println("поиграем!")
		if b_frac != 0 {
			fmt.Println("второй операнд нужен целый, но можно и отбросить дробную часть ;)")
			b = float64(int(b))
			b_frac = 0
		}
		switch int(b) % 5 {
		case 0:
			fmt.Println("результат: камень")
		case 1:
			fmt.Println("результат: ножницы")
		case 2:
			fmt.Println("результат: бумага")
		default: // полный бред и доширак-программирование :)
			rand.Seed(int64(a))
			opList := []string{"+", "-", "/", "*", "%", "//"}
			myCalc(op1, op2, opList[rand.Intn(len(opList))])
		}
	}
}

func main() {
	myCalc(getExpr("Введите выражение (например 12.3 + 4): "))
}
