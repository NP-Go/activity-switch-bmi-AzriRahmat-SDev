package main

import (
	"fmt"
	"math"
)

func main() {
	var userWeight float64
	var userHeight float64
	var userBmi float64

	//Insert Code Here
	fmt.Println("Please enter your weight in KG")
	fmt.Scanln(&userWeight)
	fmt.Println("Please enter your height in M")
	fmt.Scanln(&userHeight)

	switch userBmi = userWeight / math.Pow(userHeight, 2); {
	case userBmi < 18.5:
		fmt.Println("You are under weight")
	case userBmi >= 18.5 && userBmi < 24.9:
		fmt.Println("Healthy Weight, your BMI is: ", userBmi)
	case userBmi >= 25 && userBmi < 29.9:
		fmt.Println("Overweight, your BMI is: ", userBmi)
	case userBmi >= 30 && userBmi < 34.9:
		fmt.Println("Obese, your BMI is: ", userBmi)
	case userBmi >= 35 && userBmi < 39.9:
		fmt.Println("Severely Obese")
	default:
		fmt.Println("Morbidly Obese")
	}
}
