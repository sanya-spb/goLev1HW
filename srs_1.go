// Напишите программу для вычисления площади прямоугольника. Длины сторон прямоугольника должны вводиться пользователем с клавиатуры.

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func scanFloat32(msg string) float32 {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print(msg)
		str1, _ := reader.ReadString('\n')

		// remove newline
		str1 = strings.Replace(str1, "\n", "", -1)

		// replace comma to dot
		str1 = strings.Replace(str1, ",", ".", -1)

		// convert string variable to float variable
		num, e := strconv.ParseFloat(str1, 32)
		if e != nil {
			println("Требуется вещественное число!")
			continue
		}
		return float32(num)
	}
}

func main() {
	a := scanFloat32("Сторона a: ")
	b := scanFloat32("Сторона b: ")

	fmt.Printf("Площадь прямоугольника: %0.2f * %0.2f = %0.2f\n", a, b, a*b)
}
