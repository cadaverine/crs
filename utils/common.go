package utils

import "fmt"

func StringifyFloatSlice(s []float64, width, precision int) string {
	var result string
	format := fmt.Sprintf("%%%v.%vf", width, precision)

	for _, item := range s {
		result += fmt.Sprintf(format, item)
	}

	return result
}

func StringifyIntSlice(s []int, width int) string {
	var result string
	format := fmt.Sprintf("%%%vd", width)

	for _, item := range s {
		result += fmt.Sprintf(format, item)
	}

	return result
}
