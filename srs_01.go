// Напишите приложение, рекурсивно вычисляющее заданное из стандартного ввода число Фибоначчи.

package main

import (
	"fmt"
)

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

	return fibonacchi(n-2) + fibonacchi(n-1)
}

func main() {
	num, _ := scanNum("Введите натуральное число: ")

	fmt.Printf("fibonacchi(%d) = %d\n", num, fibonacchi(num))
}
