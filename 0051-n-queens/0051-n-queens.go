func solveNQueens(n int) [][]string {
    solutions := [][]string{}

    board := make([][]rune, n)
    for i := range n {
        board[i] = make([]rune, n)
        for j := range n {
            board[i][j] = '.'
        }
    }

    // state menagement
    vertical := make([]bool, n)
    diagonal := make([]bool, n * 2)
    anti_diagonal := make([]bool, n * 2)

    var put func(row int) 
    put = func(row int) {
        if row == n {
            solution := make([]string, n)
            for i := range n {
                solution[i] = string(board[i])
            }

            solutions = append(solutions, solution)
            return
        }

        for col := range n {
            if vertical[col] || diagonal[row + col] || anti_diagonal[row - col + n] {
                continue
            }

            board[row][col] = 'Q'
            vertical[col] = true
            diagonal[row + col] = true
            anti_diagonal[row - col + n] = true

            put(row + 1)

            anti_diagonal[row - col + n] = false
            diagonal[row + col] = false
            vertical[col] = false
            board[row][col] = '.'
        }
    }

    put(0)

    return solutions
}

