package main

import (
	"testing"
)

// Test various operations using the Operation type
func TestCalculator(t *testing.T) {
	tests := []struct {
		op      Operation
		lhs     float64
		rhs     float64
		want    float64
		wantErr bool
	}{
		{Add, 10, 5, 15, false},
		{Subtract, 10, 5, 5, false},
		{Multiply, 10, 5, 50, false},
		{Divide, 10, 5, 2, false},
		{Divide, 10, 0, 0, true}, // division by zero
		{Modulo, 10, 3, 1, false},
		{Modulo, 10, 0, 0, true}, // modulo by zero
		{Power, 2, 3, 8, false},
		{Sqrt, 16, 0, 4, false},
		{Sqrt, -16, 0, 0, true}, // square root of negative number
	}

	for _, tt := range tests {
		got, err := tt.op.calculate(tt.lhs, tt.rhs)
		if (err != nil) != tt.wantErr {
			t.Errorf("Calculate() error = %v, wantErr %v", err, tt.wantErr)
			continue
		}
		if got != tt.want {
			t.Errorf("Calculate() got = %v, want %v", got, tt.want)
		}
	}
}
