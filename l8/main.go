package main

import "fmt"

func main() {
	var num int64 = 5 //наше число (0101 в двоичной системе)
	i := 0            //позиция бита начиная с 0
	value := 0        //значение на котрое меняем i-й бит

	if value == 1 {
		num |= 1 << i // Установка бита в 1
	} else {
		num &^= 1 << i // Установка бита в 0
	}
	fmt.Printf("%d\n", num)
}
