package main

import "fmt"
type Celsius float64
type Fahrenheit float64
const (
AbsoluteZeroC Celsius = -273.15
FreezingC Celsius = 0
BoilingC Celsius = 100
)
func main(){
	fmt.Println(CToF(12.1))
	fmt.Println(FToC(20))
	var cel Celsius = 100
	var fa Fahrenheit = 100
	//fmt.Println(cel == fa)
}
func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9)}