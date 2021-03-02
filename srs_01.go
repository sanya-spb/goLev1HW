// Познакомьтесь с алгоритмом сортировки вставками. Напишите приложение, которое принимает на вход набор целых чисел и отдаёт на выходе его же в отсортированном виде. Сохраните код, он понадобится нам в дальнейших уроках.

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// сортировка вставками (на wikipedia даны 3 варианта, взял второй от Н. Вирта)
func mySort(A []int64) {
	for i := 1; i < len(A); i++ {
		x, j := A[i], i
		for ; j >= 1 && A[j-1] > x; j-- { // на wikipedia тут ошибка: j > 1 и тогда первый элемент не сортируется..
			A[j] = A[j-1]
		}
		A[j] = x
	}
}

func main() {
	var counter []int64
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		num, err := strconv.ParseInt(scanner.Text(), 10, 64)
		if err != nil {
			panic(err)
		}
		counter = append(counter, num)
	}

	// сначала сделал так, потом еще раз перечитал условие, ok, сделаем вставками..
	// sort.Slice(counter, func(i, j int) bool { return counter[i] < counter[j] })

	mySort(counter)

	// не стал выводить в цикле, так читабельней..
	fmt.Printf("%v\n", counter)
}
