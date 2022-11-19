var dfs func(currC []int, index, target int)

func combinationSum(candidates []int, target int) [][]int {
	var result [][]int
    var currC []int // current combination
    
    if len(candidates) == 0 {
		return result
	}
    
    dfs = func(currC []int, index, target int) {
        if target < 0 {
            return
        }

        if target == 0 {
            copyC := make([]int, len(currC))
            copy(copyC, currC)
            result = append(result, copyC)
        }

        for i := index; i < len(candidates); i++ {
            currC = append(currC, candidates[i])
            dfs(currC, i, target - candidates[i])
            currC = currC[:len(currC)-1]
        }
    }


	dfs(currC, 0, target)

	return result
}