package zj012_矩阵中的路径

//深度优先 dsp
func exist(board [][]byte, word string) bool {
	if len(board) == 0 {
		return false
	}
	words := []byte(word)
	m := len(board)
	n := len(board[0])
	isexist := false
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			isexist = dsp(board, m, n, words, i, j, 0)
			if isexist {
				return true
			}
		}
	}
	return isexist
}

func dsp(board [][]byte, m, n int, words []byte, i, j int, k int) bool {
	if i < 0 || i >= m || j < 0 || j >= n || k >= len(words) || board[i][j] != words[k] {
		return false
	}
	if len(words)-1 == k {
		return true
	}
	t := board[i][j]
	board[i][j] = ' '
	b := dsp(board, m, n, words, i+1, j, k+1) || dsp(board, m, n, words, i-1, j, k+1) ||
		dsp(board, m, n, words, i, j+1, k+1) || dsp(board, m, n, words, i, j-1, k+1)
	board[i][j] = t
	return b
}
