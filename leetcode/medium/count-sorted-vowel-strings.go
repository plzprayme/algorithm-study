func countVowelStrings(n int) int {
    a, b, c, d, e := make([]int, n), make([]int, n), make([]int, n), make([]int, n), make([]int, n)
    a[0], b[0], c[0], d[0], e[0] = 1, 1, 1, 1, 1
    
    for i := 1; i < n; i++ {
        a[i] = a[i - 1]
        b[i] = a[i - 1] + b[i - 1]
        c[i] = a[i - 1] + b[i - 1] + c[i - 1]
        d[i] = a[i - 1] + b[i - 1] + c[i - 1] + d[i - 1]
        e[i] = a[i - 1] + b[i - 1] + c[i - 1] + d[i - 1] + e[i - 1]
    }
    
    return a[n - 1] + b[n - 1] + c[n - 1] + d[n - 1] + e[n - 1]
}
