// С клавиатуры вводится трехзначное число. Выведите цифры, соответствующие количество сотен, десятков и единиц в этом числе.

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func scanLNumStr(msg string, strLen int) string {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print(msg)
		str1, _ := reader.ReadString('\n')

		// remove newline
		str1 = strings.Replace(str1, "\n", "", -1)

		// check if string is number & check lenght
		if _, e := strconv.Atoi(str1); e != nil || len(str1) != strLen {
			fmt.Printf("Требуется %d-значное число!\n", strLen)
			continue
		}

		return str1
	}
}

func main() {
	numStr := scanLNumStr("Введите трехзначное число: ", 3)

	// т.к. цифры попадают в ASCII, обойдемся без UTF-8, иначе пришлось бы юзать string([]rune(numStr)[0])
	fmt.Printf("сотни:   %s\n", string(numStr[0]))
	fmt.Printf("десятки: %s\n", string(numStr[1]))
	fmt.Printf("единицы: %s\n", string(numStr[2]))
}
