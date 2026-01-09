package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type DivisionError struct {
	Dividend float64
	Divisor  float64
}

type ValidationError struct {
	Field string
	Value float64
}

func (d DivisionError) Error() string {
	return fmt.Sprintf("division error: cannot divide %.0f by %.0f\n", d.Dividend, d.Divisor)
}

func (v ValidationError) Error() string {
	return fmt.Sprintf("validation error: invalid %s value %.0f\n", v.Field, v.Value)
}

func safeDivide(divisor, dividend float64, mode string) (float64, error) {
	if divisor == 0 {
		return 0, DivisionError{
			Dividend: dividend,
			Divisor:  divisor,
		}
	}

	if dividend < 0.0 && mode == "strict" {
		return 0, ValidationError{
			Field: "dividend",
			Value: dividend,
		}
	}

	if divisor < 0.0 && mode == "strict" {
		return 0, ValidationError{
			Field: "divisor",
			Value: divisor,
		}
	}
	result := dividend / divisor
	return result, nil
}

func performCalculation(divisor, dividend float64, mode, calName string) (float64, error) {
	res, err := safeDivide(dividend, divisor, mode)

	if err != nil {
		return 0, fmt.Errorf("calculation failed in %s: %w\n", calName, err)
	}

	return res, nil
}

func main() {
	// Read input
	var operationInput string
	var settingsInput string
	fmt.Scanln(&operationInput)
	fmt.Scanln(&settingsInput)

	// Parse operation input (operation,num1,num2,precision)
	operationParts := strings.Split(operationInput, ",")
	operation := operationParts[0]
	num1, _ := strconv.ParseFloat(operationParts[1], 64)
	num2, _ := strconv.ParseFloat(operationParts[2], 64)
	precision, _ := strconv.Atoi(operationParts[3])

	// Parse settings input (calculator_name,mode)
	settingsParts := strings.Split(settingsInput, ",")
	calculatorName := settingsParts[0]
	mode := settingsParts[1]

	// TODO: Write your code below
	// 1. Create DivisionError and ValidationError structs with Error() methods
	// 2. Implement safeDivide function
	// 3. Implement performCalculation function
	// 4. Call performCalculation and handle results
	res, err := performCalculation(num1, num2, mode, calculatorName)
	if err == nil {
		fmt.Printf("Calculation successful: %.0f\n", res)
	} else {
		fmt.Printf("Calculation failed: %s\n", err)
	}
	// 5. Use errors.As to check for specific error types
	var divError DivisionError
	var valError ValidationError

	divDetected := errors.As(err, &divError)
	valDetected := errors.As(err, &valError)

	fmt.Printf("Checking for division error: %t\n", divDetected)
	fmt.Printf("Checking for validation error: %t\n", valDetected)

	if errors.As(err, &divError) {
		fmt.Println("Division Error Details:")
		fmt.Printf("Dividend: %.0f\n", divError.Dividend)
		fmt.Printf("Divisor: %.0f\n", divError.Divisor)
	}

	if errors.As(err, &valError) {
		fmt.Println("Validation Error Details:")
		fmt.Printf("Field: %s\n", valError.Field)
		fmt.Printf("Value: %.0f\n", valError.Value)
	}
	// 7. Display final summary
	fmt.Println("Calculator Summary:")
	fmt.Printf("Name: %s\n", calculatorName)
	fmt.Printf("Mode: %s\n", mode)
	fmt.Printf("Operation: %s\n", operation)
	fmt.Printf("Input: %.0f %s %.0f\n", num1, operation, num2)
	fmt.Printf("Precision: %d decimal places\n", precision)
	if err != nil {
		fmt.Println("Status: Failed")
	} else {

		fmt.Println("Status: Success")
	}
}
