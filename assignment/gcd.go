package main

import "fmt"

func main(){
	gCd := gcd(20,100)
	fmt.Println((gCd))
}
func gcd(x, y int) int {
	for y != 0 {
	x, y = y, x%y
	}
	return x
	}