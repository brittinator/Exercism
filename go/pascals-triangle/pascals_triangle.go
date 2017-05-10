package pascal

const testVersion = 1

// Triangle returns Pascal's triangle given the numRows
func Triangle(numRows int) [][]int {
	triangle := make([][]int, numRows)
	// guard
	if numRows == 0 {
		return triangle
	}
	triangle[0] = []int{1}
	for i := 1; i < numRows; i++ {
		row := make([]int, i+1)
		triangle[i] = row
		for col := 0; col < len(row); col++ {
			var edge string
			if col == 0 {
				edge = "left"
			} else if col == len(row)-1 {
				edge = "right"
			}
			triangle[i][col] = sum(i, col, edge, triangle)
		}
	}
	// fmt.Printf("triangle is: \n %v", triangle)
	return triangle
}

// sum takes the position of where the sum will wind up, as well as if it's on the edge
func sum(row, col int, edge string, triangle [][]int) int {
	if edge != "" {
		if edge == "left" {
			return triangle[row-1][col]

		}
		return triangle[row-1][col-1]
	}
	return triangle[row-1][col-1] + triangle[row-1][col]
}
