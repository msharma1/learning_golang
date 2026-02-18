package main

func CtoF(c Celcius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }
func FtoC(f Fahrenheit) Celcius { return Celcius(Celcius((f - 32) * 5 / 9)) }
func CtoK(c Celcius) Kelvin     { return Kelvin(c + 273.15) }
func KtoC(k Kelvin) Celcius     { return Celcius(k - 273.15) }
