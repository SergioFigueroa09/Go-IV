package main
import "fmt"
//intercambia realiza el movimiento de valores de a hacia b y viceversa.
func intercambia(a *int, b *int){
	aux := *a
	*a = *b
	*b = aux
}

func main(){
	var a int
	var b int
	fmt.Scanln(&a)
	fmt.Scanln(&b)
	intercambia(&a,&b)
	fmt.Println(a)
	fmt.Println(b)
}