package main

import "fmt"

func main(){
	var three, four = returnThreeAndFour()
	fmt.Println(three,four)
}
func returnThreeAndFour() (int,int) {
	return 3, 4
}