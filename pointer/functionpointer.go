package main

import "fmt"

func main(){
	fmt.Println(f() == f()) // "false"
}
func f() *int {
v := 1
return &v
}
