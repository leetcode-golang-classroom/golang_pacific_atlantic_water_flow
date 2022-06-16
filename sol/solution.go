package sol

type Pair struct {
	row, col int
}

func pacificAtlantic(heights [][]int) [][]int {
	ROW := len(heights)
	COL := len(heights[0])

	pacific := make(map[Pair]struct{})
	atlantic := make(map[Pair]struct{})

	// find water enable cell
	var dfs func(row int, col int, prevHight int, visit map[Pair]struct{})
	dfs = func(row int, col int, prevHight int, visit map[Pair]struct{}) {
		if row < 0 || row >= ROW || col < 0 || col >= COL || heights[row][col] < prevHight {
			return
		}
		cell := Pair{row: row, col: col}
		if _, visited := visit[cell]; visited {
			return
		}
		visit[cell] = struct{}{}
		dfs(row-1, col, heights[row][col], visit)
		dfs(row+1, col, heights[row][col], visit)
		dfs(row, col-1, heights[row][col], visit)
		dfs(row, col+1, heights[row][col], visit)
	}

	for row := range heights {
		dfs(row, 0, heights[row][0], pacific)
		dfs(row, COL-1, heights[row][COL-1], atlantic)
	}
	for col := range heights[0] {
		dfs(0, col, heights[0][col], pacific)
		dfs(ROW-1, col, heights[ROW-1][col], atlantic)
	}
	result := [][]int{}
	for row := 0; row < ROW; row++ {
		for col := 0; col < COL; col++ {
			cell := Pair{row: row, col: col}
			_, inPacific := pacific[cell]
			_, inAtlantic := atlantic[cell]
			if inAtlantic && inPacific {
				result = append(result, []int{row, col})
			}
		}
	}
	return result
}
