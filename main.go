package main

import (
	"fmt"
	"math"
)

const IMTPower = 2

func main() {
	fmt.Println("BMI-Check")
	userWeight, userHeight := getUserInput()
	IMT := calculateIMT(userWeight, userHeight)
	outputResult(IMT)
	switch {
	case IMT < 16:
		fmt.Println("You are strong underweight")
	case IMT < 25:
		fmt.Println("You are underweight")
	case IMT < 30:
		fmt.Println("You are overweight")
	default:
		fmt.Println("you have a degree of obesity")
	}
}

func outputResult(imt float64) {
	result := fmt.Sprintf("Your BMI: %.0f", imt)
	fmt.Println(result)
}

func calculateIMT(userWeight float64, userHeight float64) float64 {
	IMT := userWeight / math.Pow(userHeight/100, IMTPower)
	return IMT
}

func getUserInput() (float64, float64) {
	var userHeight float64
	var userWeight float64
	fmt.Print("Enter your height in centimeters: ")
	fmt.Scan(&userHeight)
	fmt.Print("Enter your weight: ")
	fmt.Scan(&userWeight)
	return userWeight, userHeight
}
