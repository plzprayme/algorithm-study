func numberOfSteps(num int) int {
    count := 0
    for num > 0 {
        count++
        if num % 2 == 0 {
            num /= 2 
            continue
        }
        num -= 1
    }
    return count
}
