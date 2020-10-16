package main

import "fmt"

//Fibo el numero de Fibonacci
func Fibo(x int) int{
	if x==0 {
		return 0
	} else if x==1{
		return 1
	} else {
		return Fibo(x-1)+Fibo(x-2)
	}
}

func main() {
	fmt.Println(Fibo(8))
}