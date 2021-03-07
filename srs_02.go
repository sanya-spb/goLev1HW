// Напишите приложение, рекурсивно вычисляющее заданное из стандартного ввода число Фибоначчи.
// Оптимизируйте приложение за счёт сохранения предыдущих результатов в мапе.

package main

import (
	"fmt"
)

var Fibonacchi map[int]int

func scanNum(msg string) (int, error) {
	var a int
	for {
		fmt.Print(msg)
		n, err := fmt.Scan(&a)
		if err != nil || n != 1 || a < 1 {
			fmt.Println("Требуется натуральное число 1,2,.. !")
			continue
		}
		return a, err
	}
}

func fibonacchi(n int) int {
	if n == 0 || n == 1 {
		return n
	}
	result := fibonacchi(n-2) + fibonacchi(n-1)
	Fibonacchi[n] = result
	return result
}

func main() {
	num, _ := scanNum("Введите натуральное число: ")

	Fibonacchi = make(map[int]int)
	fmt.Printf("fibonacchi(%d) = %d\n", num, fibonacchi(num))
	fmt.Printf("eval fibonacchi(%d) = %v\n", num, Fibonacchi)
}
