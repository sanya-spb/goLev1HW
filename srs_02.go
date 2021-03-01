// Задание для продвинутых (необязательное). Написать приложение, которое ищет все простые числа от 0 до N включительно. Число N должно быть задано из стандартного потока ввода. P.S. в math есть подсказка как это сделать

package main

import (
	"fmt"
	"math"
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

// O(√(n))
func isPrime(n int) bool {
	// чтоб не заниматься расчетами каждую итерацию цикла
	max := int(math.Sqrt(float64(n)))
	for i := 2; i <= max; i++ {
		// проверка на делимость без остатка
		if n%i == 0 {
			// ИМХО для наглядности и информативности..
			fmt.Printf("\n%d делится на %d.\n", n, i)
			return false
		}
	}
	fmt.Println() // да, смешал логику и форматирование.. самому не нравится
	return true
}

func main() {
	num, _ := scanNum("Введите натуральное число: ")

	if isPrime(num) {
		fmt.Printf("число %d простое!\n", num)
	} else {
		fmt.Printf("число %d не простое!\n", num)
	}
}
