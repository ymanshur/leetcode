var numStep func(num int) int

func climbStairs(n int) int {
    numStepMap := make(map[int]int)

    numStep = func(num int) int {
        if _, ok := numStepMap[num]; ok {
            return numStepMap[num]
        }
        
        numStepMap[num] = numStep(num-1) + numStep(num-2)
        
        return numStepMap[num]
    }
    
    numStepMap[1], numStepMap[2], numStepMap[3] = 1, 2, 3
    return numStep(n)
    
}