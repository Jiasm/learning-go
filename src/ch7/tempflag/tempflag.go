package main

import (
	. "ch7/tempflag/tempconv"
	"fmt"
	"flag"
)

type celsiusFlag struct { Celsius }

var temp = CelsiusFlag("temp", 20.0, "the temperature")

func main() {
	flag.Parse()
	fmt.Println(*temp)
}

func (f *celsiusFlag) Set(s string) error {
	var unit string
	var value float64

	fmt.Sscanf(s, "%f%s", &value, &unit)

	switch unit {
	case "C", "摄氏度":
		f.Celsius = Celsius(value)
		return nil
	case "F", "华氏度":
		f.Celsius = F2C(Fahrenheit(value))
		return nil
	}

	return fmt.Errorf("invalid temperature %q", s)
}

func CelsiusFlag(name string, value Celsius, usage string) *Celsius {
	f := celsiusFlag{value}
	flag.CommandLine.Var(&f, name, usage)

	return &f.Celsius
}