package main

// Вклад в банке составляет x рублей. Ежегодно он увеличивается на p процентов,
// после чего дробная часть копеек отбрасывается. Каждый год сумма вклада становится больше.
// Определите, через сколько лет вклад составит не менее y рублей.
// Входные данные:
// Программа получает на вход три натуральных числа: x, p, y.
// Выходные данные:
// Программа должна вывести одно целое число.

import "fmt"

func main() {
	var x, p, y, year int
	fmt.Scan(&x, &p, &y)
	year = 0

	for x < y {
		x += int(x*p) / 100
		year += 1
	}
	fmt.Println(year)
}
