import "strconv"

func isPalindrome(x int) bool {
    xs := strconv.Itoa(x)
    
    for i := 0; i < len(xs) / 2; i++ {
        if xs[i] != xs[len(xs) - 1 - i] {
            return false
        } 
    }
    
    return true
}