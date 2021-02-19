// С клавиатуры вводится трехзначное число. Выведите цифры, соответствующие количество сотен, десятков и единиц в этом числе.

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func scanLNum(msg string, strLen int) int {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print(msg)
		str1, _ := reader.ReadString('\n')

		// remove newline
		str1 = strings.Replace(str1, "\n", "", -1)

		// check lenght
		if len(str1) != strLen {
			fmt.Printf("Требуется %d-значное число!\n", strLen)
			continue
		}

		// convert string variable to int variable
		num, e := strconv.Atoi(str1)
		if e != nil {
			fmt.Printf("Требуется %d-значное число!\n", strLen)
			continue
		}
		return num
	}
}

func main() {
	num := scanLNum("Введите трехзначное число: ", 3)

	fmt.Printf("сотни:   %d\n", int(num/100))
	fmt.Printf("десятки: %d\n", int((num-int(num/100)*100)/10))
	fmt.Printf("единицы: %d\n", num%10)
}
