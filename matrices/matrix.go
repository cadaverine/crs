package matrices

import (
	"math/rand"

	"github.com/cadaverine/crs-lab/utils"
)

type Matrix [][]float64

func CreateMatrix(rowsNum, colsNum int) Matrix {
	matrix := make([][]float64, rowsNum)

	for i := 0; i < rowsNum; i++ {
		matrix[i] = make([]float64, colsNum)
	}

	return matrix
}

func (m *Matrix) FillRandom(max, density float64) {
	for i := 0; i < len(*m); i++ {
		for j := 0; j < len((*m)[0]); j++ {
			if rand.Float64()*100.0 < density {
				(*m)[i][j] = rand.Float64() * max
			}
		}
	}
}

func (m *Matrix) Collapse() *CRS {
	values := make([]float64, 0)
	columns := make([]int, 0)
	pointers := make([]int, 1)

	counter := 0

	for _, row := range *m {
		for j, item := range row {
			if item != 0 {
				values = append(values, item)
				columns = append(columns, j)
				counter++
			}
		}

		pointers = append(pointers, counter)
	}

	return &CRS{values, columns, pointers, len(*m), len((*m)[0])}
}

func (m Matrix) String() string {
	var result string

	for _, row := range m {
		result += utils.StringifyFloatSlice(row, 6, 1) + "\n\n"
	}

	return result
}
