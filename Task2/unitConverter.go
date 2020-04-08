package main

import (
	"fmt"
	"math"
	"reflect"
)

type (
	Feet         float64
	Centimeter   float64
	Minutes      float64
	Seconds      float64
	Hours        float64
	Milliseconds float64
	Celsius      float64
	Fahrenheit   float64
	Kelvin       float64
	Radian       float64
	Degree       float64
	Kilograms    float64
	Pounds       float64
)

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
	fmt.Println("\nThere are", cvrPointer.NumMethod(), "conversion options")

	selection := make(map[int]string)
	for i := 1; i <= cvrPointer.NumMethod(); i++ {
		selection[i] = cvrPointer.Method(i - 1).Name
	}
	fmt.Println(selection)

	var SN int
	var value float64
	fmt.Printf("Enter the serial number of your prefered conversion then the value to convert : ")
	fmt.Scan(&SN, &value)

	switch SN {
	case 1:
		fmt.Println("Coversion value is:", cvr.CelsiuisToFahrenheit(Celsius(value)))
	case 2:
		fmt.Println("Coversion value is:", cvr.CentimeterToFeet(Centimeter(value)))
	case 3:
		fmt.Println("Coversion value is:", cvr.FeetToCentimeter(Feet(value)))
	case 4:
		fmt.Println("Coversion value is:", cvr.HoursToMinutes(Hours(value)))
	case 5:
		fmt.Println("Coversion value is:", cvr.HoursToSeconds(Hours(value)))
	case 6:
		fmt.Println("Coversion value is:", cvr.KelvinToCelsius(Kelvin(value)))
	case 7:
		fmt.Println("Coversion value is:", cvr.KelvinToFahrenheit(Kelvin(value)))
	case 8:
		fmt.Println("Coversion value is:", cvr.KilogramsToPounds(Kilograms(value)))
	case 9:
		fmt.Println("Coversion value is:", cvr.MinutesToSeconds(Minutes(value)))
	case 10:
		fmt.Println("Coversion value is:", cvr.RadianToDegree(Radian(value)))
	case 11:
		fmt.Println("Coversion value is:", cvr.Seconds2Minutes(Seconds(value)))

	}
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

func (cvr Converter) Seconds2Minutes(s Seconds) Minutes {
	return Minutes(s / 60)
}

func (cvr Converter) HoursToMinutes(h Hours) Minutes {
	return Minutes(h * 60)
}

func (cvr Converter) HoursToSeconds(h Hours) Seconds {
	m := cvr.HoursToMinutes(h)
	s := cvr.MinutesToSeconds(m)
	return s
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
