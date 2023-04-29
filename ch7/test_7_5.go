package main

import (
	"flag"
	"fmt"
)

type Celsius float64
type Fahrenheit float64

type Value interface {
	String() string
	Set(string) error
}

const (
	AbsoluteZeroC Celsius = -273.15
	Freezing      Celsius = 0
	Boiling       Celsius = 100
)

func CToF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

func FtoC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}

func (c Celsius) String() string {
	return fmt.Sprintf("%g 摄氏度", c)
}

type celsiusFlag struct{ Celsius }

func (f *celsiusFlag) Set(s string) error {
	var unit string
	var value float64
	//Sscanf
	fmt.Sscanf(s, "%f%s", &value, &unit)
	switch unit {
	case "C":
		f.Celsius = Celsius(value)
		return nil
	case "F":
		f.Celsius = FtoC(Fahrenheit(value))
		return nil
	}
	return fmt.Errorf("invalid temperature %q", s)
}

func CelsiusFlag(name string, value Celsius, usage string) *Celsius {
	f := celsiusFlag{value}
	//f 实现了Set方法
	flag.CommandLine.Var(&f, name, usage)
	return &f.Celsius
}

var temp = CelsiusFlag("temp", 20.0, "the temperature")

func main() {
	//var a Celsius = 1.2
	//fmt.Println(a)

	flag.Parse()
	fmt.Println(*temp)
}
