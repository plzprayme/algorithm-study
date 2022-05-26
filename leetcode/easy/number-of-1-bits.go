func hammingWeight(num uint32) int {
    var count uint32
    for num > 0 {
        count += num % 2
        num = num / 2
    }
    return int(count)
}
