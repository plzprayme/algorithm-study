func findUnsortedSubarray(nums []int) int {
    unsorted := copy(nums)
    sort.Ints(nums)
    
    left, right := -1, -1
    for i, v := range nums {
        if v == unsorted[i] {
            continue
        }
        
        if left == -1 {
            left = i
        } else {
            right = i
        }
    }

    answer := right - left + 1
    if answer == 1 {
        return 0
    }
    return answer    
}

func copy(origin []int) []int {
    copied := make([]int, 0)
    for _, v := range origin {
        copied = append(copied, v)
    }
    return copied
}