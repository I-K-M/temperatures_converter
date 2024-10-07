package main

import (
	"fmt"
	"strings"
)

var temperature float32

func getTemperatureType() string {
	var temperatureType string
	fmt.Printf("Is your temperature in Celsius or Fahrenheit? Please type in C or F.\n")
	fmt.Scanln(&temperatureType)
	for strings.ToUpper(temperatureType) != "C" && strings.ToUpper(temperatureType) != "F" {
		fmt.Printf("-------------\n!-Entered value not understood-!\n\n")
		fmt.Printf("Is your temperature in Celsius or Fahrenheit? Please type in C or F.\n")
		fmt.Scanln(&temperatureType)
	}
	return strings.ToUpper(temperatureType)
}

func convert(temperatureType string) {
	if strings.ToUpper(temperatureType) == "C" {
		fmt.Printf("Enter a temperature in Celsius:")
		fmt.Scanln(&temperature)
		fmt.Printf("----------------\nThis temperature is %.2f° in Fahrenheit\n", (temperature*1.8)+32)
	} else if strings.ToUpper(temperatureType) == "F" {
		fmt.Printf("Enter a temperature in Fahrenheit:")
		fmt.Scanln(&temperature)
		fmt.Printf("----------------\nThis temperature is %.2f° in Celsius\n", (temperature-32)/1.8)
	}
}

func main() {
	var goOn string
	fmt.Println("--------------\nUNITS CONVERTER PROGRAM | WELCOME\n ----------------")
	for {
		temperatureType := getTemperatureType()
		convert(temperatureType)
		fmt.Printf("Do you want to convert a new temperature? Please type YES or NO.\n")
		fmt.Scanln(&goOn)
		if strings.ToUpper(goOn) == "NO" {
			fmt.Printf("--------------\n-OK, good bye, have a nice day-\n-----------------\n")
			break
		}
	}
}
