// Напишите приложение, рекурсивно вычисляющее заданное из стандартного ввода число Фибоначчи.

package fibonacchi

// recursion version of Fib
func FibonacciR(n uint8) uint64 {
	if n == 0 || n == 1 {
		return uint64(n)
	}

	return FibonacciR(n-2) + FibonacciR(n-1)
}

// version of Fib with maps
func FibonacciM(n uint8) uint64 {
	// инициализируем map, задаем первые 2 элемента последовательности
	var result map[uint8]uint64 = make(map[uint8]uint64, n+1)
	result[0] = 0
	result[1] = 1
	// остальное расчитываем в цикле
	for i := uint8(2); i <= n; i += 1 {
		result[i] = result[i-1] + result[i-2]
	}
	// читаем из мапы (полезно только когда n=0, иначе нужен только последний элемент мапы, но она у нас не сортирована..)
	return result[n]
}
