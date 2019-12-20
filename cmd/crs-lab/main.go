package main

import (
	"fmt"

	"github.com/cadaverine/crs-lab/matrices"
)

func main() {
	m1 := matrices.CreateMatrix(3, 10)
	m1.FillRandom(49.9, 30)

	crs1 := m1.Collapse()

	fmt.Printf("\nCreated:\n\n%v\n", m1)
	fmt.Printf("\nCollapsed(CRS):\n\n%v\n", crs1)

	m2 := matrices.CreateMatrix(3, 10)
	m2.FillRandom(49.9, 50)

	crs2 := m2.Collapse()

	fmt.Printf("\nCreated:\n\n%v\n", m2)
	fmt.Printf("\nCollapsed(CRS):\n\n%v\n", crs2)

	sum, _ := matrices.Sum(crs1, crs2)
	fmt.Printf("\nSum:\n\n%v\n", sum)
	fmt.Printf("\nExpanded sum:\n\n%v\n", sum.Expand())
}
