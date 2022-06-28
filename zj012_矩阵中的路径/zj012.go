package zj012_矩阵中的路径

type Pos struct {
	X int
	Y int
}

func exist(board [][]byte, word string) bool {
	if len(board) == 0 {
		return false
	}
	words := []byte(word)

	m := len(board)
	n := len(board[0])
	starts := findStart(board, words[0])
	if len(starts) == 0 {
		return false
	}

	for _, startPos := range starts {
		x, y := startPos.X, startPos.Y
		used := make(map[*Pos]bool)
		used[startPos] = true
		for idx, w := range words {
			if idx == 0 {
				continue
			}
			if x < m-1 && y < n-1 {
				if board[x+1][y] == w || board[x][y+1] == w {
					used[startPos] = true
				}
			}
		}
	}
	return true
}

func findStart(board [][]byte, c byte) []*Pos {
	m := len(board)
	n := len(board[0])
	sp := make([]*Pos, 0)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			sp = append(sp, &Pos{i, j})
		}
	}
	return sp
}
