package main

import (
	"fmt"
	"math"
	"reflect"
)

type (
	Feet       float64
	Centimeter float64
	//hour			float64
	Minutes      float64
	Seconds      float64
	Milliseconds float64
	Celsius      float64
	Fahrenheit   float64
	Kelvin       float64
	Radian       float64
	Degree       float64
	Kilograms    float64
	Pounds       float64
)

/*
type Converter struct(
	FeetToCentimeter	Centimeter
	CentimeterToFeet	Feet
	MinutesToSeconds	Seconds
	seconds2Minutes	Minutes
	//Minutes@millliseconds	Milliseconds
	seconds2milliseconds	Milliseconds
	milliseconds2seconds	Seconds
	//milliseconds2Minutes	Minutes
	CelsiuisToFahrenheit	Fahrenheit
	fahrenheit2Celsius	Celsius
	KelvinToCelsius	celcius
	KelvinToFahrenheit	Fahrenheit
	//Celsius@Kelvin	Kelvin
	//fahreheit@Kelvin	Kelvin
	RadianToDegree	Degree
	Degree@Radian	Radian
	kilogram2pounds	Pounds
	Pounds@Kilograms
)
*/

// Converter converts values between scientific units
type Converter struct{}

var cvr Converter

func main() {

	var cvr Converter

	// displays examples of some avaialble conversions
	fmt.Println("16cm in Feet is ", cvr.CentimeterToFeet(16), "ft")
	fmt.Println("9 Minutes in Seconds is", cvr.MinutesToSeconds(9), "Seconds")

	// dislays Kelvin to Fahrenheit conversion that internally contains intermidiate Celsius conversion
	fmt.Println("\nThe Kelvin to Fahrenheit method is made up of two other methods.")
	var kelvinTemperature Kelvin
	fmt.Printf("Input a temperature in Kelvin: ")
	fmt.Scan(&kelvinTemperature)
	fmt.Println("This is", cvr.KelvinToFahrenheit(kelvinTemperature), "degrees in Fahrenheit.")

	//User picks disired conversion
	cvrPointer := reflect.TypeOf(&cvr)
	fmt.Println("There are", cvrPointer.NumMethod(), "conversion options")

	selection := make(map[int]string)
	for i := 1; i <= cvrPointer.NumMethod(); i++ {
		selection[i] = cvrPointer.Method(i - 1).Name
	}
	fmt.Println(selection)

	var SN int
	var value float64
	fmt.Printf("Enter the serial number of your prefered conversion then the value to convert : ")
	fmt.Scan(&SN, &value)
	//  ???????  //
	fmt.Println(cvrPointer.Method(SN - 1))
	//  ???????  //
}

func (cvr Converter) CentimeterToFeet(c Centimeter) Feet {
	return Feet(c * 0.0328)
}

func (cvr Converter) FeetToCentimeter(f Feet) Centimeter {
	return Centimeter(f / 0.0328)
}

func (cvr Converter) MinutesToSeconds(m Minutes) Seconds {
	return Seconds(m * 60)
}

//Seconds2Minutes converts Seconds to Minutes
func (cvr Converter) Seconds2Minutes(s Seconds) Minutes {
	return Minutes(s / 60)
}

func (cvr Converter) CelsiuisToFahrenheit(c Celsius) Fahrenheit {
	return Fahrenheit((c * 9 / 5) + 32)
}

func (cvr Converter) KelvinToCelsius(k Kelvin) Celsius {
	return Celsius(k - 273.15)
}

func (cvr Converter) KelvinToFahrenheit(k Kelvin) Fahrenheit {
	c := cvr.KelvinToCelsius(k)
	f := cvr.CelsiuisToFahrenheit(c)
	return f
}

func (cvr Converter) RadianToDegree(r Radian) Degree {
	return Degree(r * 180 / math.Pi)
}

func (cvr Converter) KilogramsToPounds(kg Kilograms) Pounds {
	return Pounds(kg / 2.20462)
}
