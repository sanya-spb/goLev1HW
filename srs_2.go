// Напишите программу, вычисляющую диаметр и длину окружности по заданной площади круга. Площадь круга должна вводиться пользователем с клавиатуры.

package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func scanFloat64(msg string) float64 {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print(msg)
		str1, _ := reader.ReadString('\n')

		// remove newline
		str1 = strings.Replace(str1, "\n", "", -1)

		// replace comma to dot
		str1 = strings.Replace(str1, ",", ".", -1)

		// convert string variable to float variable
		num, e := strconv.ParseFloat(str1, 64)
		if e != nil {
			println("Требуется вещественное число!")
			continue
		}
		return num
	}
}

func main() {
	S := scanFloat64("Площадь круга: ")

	r := math.Sqrt(S / math.Pi)
	d := 2 * r
	l := 2 * math.Pi * r

	fmt.Printf("радиус: r = √(S/π) = %0.2f\n", r)
	fmt.Printf("диаметр: d = 2*r = %0.2f\n", d)
	fmt.Printf("длина окружности: l = 2πr = %0.2f\n", l)
}
