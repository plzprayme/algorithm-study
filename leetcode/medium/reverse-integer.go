func reverse(x int) int {

    str := []rune(strconv.Itoa(x))
    var i, j int
    if str[0] == 45 {
        i, j = 1, len(str) - 1
    } else {
        i, j = 0, len(str) - 1    
    }
    
    for ; i < j; i, j = i + 1, j - 1 {
        str[i], str[j] = str[j], str[i]
    }
    
    
    answer, _ := strconv.Atoi(string(str))
    
    if answer > math.MaxInt32 - 1 || answer < math.MinInt32 {
        return 0
    }
    return answer
}
