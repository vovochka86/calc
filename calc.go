package main

import (
	"fmt"
	"math"
	"errors"
)

// Mathematical operations defined as constants using iota
const (
	Add = iota
	Subtract
	Multiply
	Divide
	Modulo
	Power
	Sqrt
)

type Operation int

// Receiver function that performs the mathematical operation on one or two operands
func (op Operation) calculate(lhs, rhs float64) (float64, error) {
	switch op {
	case Add:
		return lhs + rhs, nil
	case Subtract:
		return lhs - rhs, nil
	case Multiply:
		return lhs * rhs, nil
	case Divide:
		if rhs == 0 {
			return 0, errors.New("division by zero")
		}
		return lhs / rhs, nil
	case Modulo:
		if rhs == 0 {
			return 0, errors.New("modulo by zero")
		}
		return math.Mod(lhs, rhs), nil
	case Power:
		return math.Pow(lhs, rhs), nil
	case Sqrt:
		if lhs < 0 {
			return 0, errors.New("square root of negative number")
		}
		return math.Sqrt(lhs), nil
	}
	return 0, errors.New("unhandled operation")
}

// Function to check and scan valid float64 input
func scanFloatInput(prompt string) float64 {
	var num float64
	for {
		fmt.Print(prompt)
		_, err := fmt.Scan(&num)
		if err == nil {
			break
		}
		fmt.Println("Invalid input. Please enter a valid number.")
	}
	return num
}

// Function to check and scan valid operation input
func scanOperationInput() Operation {
	var opChoice int
	for {
		fmt.Println("Choose an operation: (0 = Add, 1 = Subtract, 2 = Multiply, 3 = Divide, 4 = Modulo, 5 = Power, 6 = Sqrt)")
		_, err := fmt.Scan(&opChoice)
		if err == nil && opChoice >= 0 && opChoice <= 6 {
			break
		}
		fmt.Println("Invalid choice. Please select a valid operation.")
	}
	return Operation(opChoice)
}

func main() {
	// Get the first operand from user input
	lhs := scanFloatInput("Enter first number (for Sqrt, only this number will be used): ")

	// Get the operation from user input
	op := scanOperationInput()

	// If the operation is not square root, take the second number
	var rhs float64
	if op != Sqrt {
		rhs = scanFloatInput("Enter second number: ")
	}

	// Perform the calculation and handle potential errors
	result, err := op.calculate(lhs, rhs)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("The result is: %.2f\n", result)
	}
}
