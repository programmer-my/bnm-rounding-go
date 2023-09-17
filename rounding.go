package main

import (
	"fmt"
	"math"
	"strconv"
)

func minArrFloat64(arr []float64) (float64, error) {
	if len(arr) == 0 {
		return 0, fmt.Errorf("input array must not be empty")
	}

	smallest := math.Inf(1)

	for _, num := range arr {
		num := float64(num)
		if math.Min(num, smallest) == num {
			smallest = num
		}
	}

	return smallest, nil
}

// Rounds an amount (in ringgit) according to BNM rounding mechanism
// Reference: https://www.bnm.gov.my/misc/-/asset_publisher/2BOPbOBfILtL/content/about-the-rounding-mechanism
func BnmRoundStr(amountStr string) (string, error) {
	amountFloat, convErr := strconv.ParseFloat(amountStr, 64)

	if convErr != nil {
		return "", fmt.Errorf("failed to parse amount: %s", convErr)
	}

	x := math.Floor(float64(amountFloat*10)) / 10
	y := x + 0.05
	z := y + 0.05

	diffX := math.Abs(x - amountFloat)
	diffY := math.Abs(y - amountFloat)
	diffZ := math.Abs(z - amountFloat)

	diffs := []float64{diffX, diffY, diffZ}
	smallestDiff, err := minArrFloat64(diffs)

	if err != nil {
		return "", fmt.Errorf("unexpected error: %s", err)
	}

	if smallestDiff == diffX {
		return fmt.Sprintf("%.2f", x), nil
	} else if smallestDiff == diffY {
		return fmt.Sprintf("%.2f", y), nil
	} else { // diffZ
		return fmt.Sprintf("%.2f", z), nil
	}
}
