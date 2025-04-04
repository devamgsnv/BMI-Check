package main

import (
	"fmt"
	"math"
)

const IMTPower = 2

func main() {
	for {
		fmt.Println("BMI-Check")
		userWeight, userHeight := getUserInput()
		IMT := calculateIMT(userWeight, userHeight)
		outputResult(IMT)
		isRepeatCalculation := checkRepeatCalculation()
		if !isRepeatCalculation {
			break
		}
	}
}

func outputResult(imt float64) {
	result := fmt.Sprintf("Your BMI: %.0f", imt)
	fmt.Println(result)
	switch {
	case imt < 16:
		fmt.Println("You are strong underweight")
	case imt < 25:
		fmt.Println("You are underweight")
	case imt < 30:
		fmt.Println("You are overweight")
	default:
		fmt.Println("you have a degree of obesity")
	}
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

func checkRepeatCalculation() bool {
	var userChoice string
	fmt.Print("Do you want to make another calculation (y/n): ")
	fmt.Scan(&userChoice)
	if userChoice == "y" || userChoice == "Y" {
		return true
	}
	return false
}
