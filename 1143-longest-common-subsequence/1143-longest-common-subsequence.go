func longestCommonSubsequence(text1 string, text2 string) int {
    // create 2D (tabulation) of the input
    cols , rows := len(text1), len(text2)
    subLength := make([][]int, cols+1)
    for col := range subLength {
        subLength[col] = make([]int, rows+1)
    }
    
    for col := cols-1; col >= 0; col-- {
        for row := rows-1; row >= 0; row-- {
            if text1[col] == text2[row] {
                subLength[col][row] += 1 + subLength[col+1][row+1]
            } else {
                if subLength[col+1][row] > subLength[col][row+1] {
                    subLength[col][row] = subLength[col+1][row]
                } else {
                    subLength[col][row] = subLength[col][row+1]
                }
            }
        }
    }
    
    return subLength[0][0]
}