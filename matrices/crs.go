package matrices

import (
	"errors"

	"github.com/cadaverine/crs-lab/utils"
)

type CRS struct {
	Values   []float64
	Columns  []int
	Pointers []int
	RowsNum  int
	ColsNum  int
}

func CreateCRS(m *Matrix) *CRS {
	return m.Collapse()
}

func (c *CRS) Expand() *Matrix {
	matrix := CreateMatrix(c.RowsNum, c.ColsNum)

	for i := 0; i < c.RowsNum; i++ {
		for k := c.Pointers[i]; k < c.Pointers[i+1]; k++ {
			matrix[i][c.Columns[k]] = c.Values[k]
		}
	}

	return &matrix
}

func (c CRS) String() string {
	var result string

	result += "Values:   " + utils.StringifyFloatSlice(c.Values, 6, 1) + "\n"
	result += "Columns:  " + utils.StringifyIntSlice(c.Columns, 6) + "\n"
	result += "Pointers: " + utils.StringifyIntSlice(c.Pointers, 6) + "\n"

	return result
}

func Sum(c1 *CRS, c2 *CRS) (*CRS, error) {
	if c1.RowsNum != c2.RowsNum || c1.ColsNum != c2.ColsNum {
		return nil, errors.New("Matrices sizes must be equal")
	}

	values := make([]float64, 0)
	columns := make([]int, 0)
	pointers := make([]int, 0)

	a := 0
	b := 0
	c := 0

	for i := 0; i < c1.RowsNum; i++ {
		pointers = append(pointers, c)

		na := c1.Pointers[i+1]
		nb := c2.Pointers[i+1]

		for a < na || b < nb {
			if a < na && b < nb {
				if c1.Columns[a] == c2.Columns[b] {
					value := c1.Values[a] + c2.Values[b]
					values = append(values, value)
					columns = append(columns, c1.Columns[a])
					a++
					b++
					c++
				} else if c1.Columns[a] < c2.Columns[b] {
					values = append(values, c1.Values[a])
					columns = append(columns, c1.Columns[a])
					a++
					c++
				} else {
					values = append(values, c2.Values[b])
					columns = append(columns, c2.Columns[b])
					b++
					c++
				}
			} else if a < na {
				values = append(values, c1.Values[a])
				columns = append(columns, c1.Columns[a])
				a++
				c++
			} else if b < nb {
				values = append(values, c2.Values[b])
				columns = append(columns, c2.Columns[b])
				b++
				c++
			}
		}
	}

	pointers = append(pointers, len(values))

	return &CRS{
		Values:   values,
		Columns:  columns,
		Pointers: pointers,
		RowsNum:  c1.RowsNum,
		ColsNum:  c1.ColsNum,
	}, nil
}
