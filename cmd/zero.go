package cmd

import (
	"fmt"
	"strconv"
)

func Add(first, second string) (result string) {
	num1, err := strconv.ParseFloat(first, 64)
	if err != nil {
		fmt.Println("Error: First value in invalid.")
		return
	}

	num2, err := strconv.ParseFloat(second, 64)
	if err != nil {
		fmt.Println("Error: Second value in invalid.")
		return
	}

	return fmt.Sprintf("%f", num1+num2)
}

func Subtract(from, subtract string) (result string) {
	num1, err := strconv.ParseFloat(from, 64)
	if err != nil {
		fmt.Println("Error: From value in invalid.")
		return
	}

	num2, err := strconv.ParseFloat(subtract, 64)
	if err != nil {
		fmt.Println("Error: Subtract value in invalid.")
		return
	}

	return fmt.Sprintf("%f", num1-num2)
}

func Multiply(first, second string, shouldRoundUp bool) (result string) {
	num1, err := strconv.ParseFloat(first, 64)
	if err != nil {
		fmt.Println("Error: First value is not a decimal.")
		return
	}

	num2, err := strconv.ParseFloat(second, 64)
	if err != nil {
		fmt.Println("Error: Second value is not a decimal.")
		return
	}

	if shouldRoundUp {
		return fmt.Sprintf("%.2f", num1*num2)
	}

	return fmt.Sprintf("%f", num1*num2)
}

func Divide(divide, by string, shouldRoundUp bool) (e error, result string) {
	num1, err := strconv.ParseFloat(divide, 64)
	if err != nil {
		return fmt.Errorf("first value is not a number"), ""
	}

	num2, err := strconv.ParseFloat(by, 64)
	if err != nil {
		return fmt.Errorf("second value is not a number"), ""
	}

	if shouldRoundUp {
		return nil, fmt.Sprintf("%.2f", num1/num2)
	}

	return nil, fmt.Sprintf("%f", num1/num2)
}
