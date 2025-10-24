package numbers

import (
	"fmt"
	"strconv"
)

func FormatToTwoDecimalPlaces(value float64) float64 {
	formattedValue := fmt.Sprintf("%.2f", value)
	convertedValue, _ := strconv.ParseFloat(formattedValue, 64)

	return convertedValue
}
