package main

import (
	"flag"
	"fmt"
)

func main() {
	c := flag.Float64("c", 0, "Temperature in Celcius")
	f := flag.Float64("f", 0, "Temperature in Fahrenheit")
	k := flag.Float64("k", 0, "Temperature in Kelvin")
	flag.Parse()
	fmt.Println(CtoF(Celcius(*c)))
	fmt.Println(CtoK(Celcius(*c)))
	fmt.Println(FtoC(Fahrenheit(*f)))
	fmt.Println(KtoC(Kelvin(*k)))
}
