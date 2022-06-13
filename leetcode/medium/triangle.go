func minimumTotal(triangle [][]int) int {
    maxR := len(triangle)
    
    dp := make([][]int, 0, maxR)
    for i := 0; i < maxR; i++ {
        dp = append(dp, make([]int, i + 1))
    }
    dp[0][0] = triangle[0][0]
    
    for r := 1; r < maxR; r++ {
        for c := 0; c < len(dp[r]); c++ {
            if c == 0 {
                dp[r][c] += dp[r-1][c]
            } else if c == len(dp[r]) - 1 {
                dp[r][c] += dp[r-1][c-1]
            } else {
                dp[r][c] += min(dp[r-1][c], dp[r-1][c-1])
            }
            dp[r][c] += triangle[r][c]
        }
    }
    
    return min(dp[maxR-1]...)
}

func min(arr ...int) int {
    min := math.MaxInt32
    for _, v := range arr {
        if min > v {
            min = v
        }
    }
    return min
}
