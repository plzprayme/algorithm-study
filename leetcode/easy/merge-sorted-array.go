func merge(nums1 []int, m int, nums2 []int, n int)  {
    merged := make([]int, 0, len(nums1))
    
    for i1, i2 := 0, 0; i1 < m || i2 < n; {
        if i1 == m {
            merged = append(merged, nums2[i2])
            i2++
            continue
        } 
        
        if i2 == n {
            merged = append(merged, nums1[i1])
            i1++
            continue
        }
        
        if nums1[i1] < nums2[i2] {
            merged = append(merged, nums1[i1])
            i1++
        } else {
            merged = append(merged, nums2[i2])
            i2++
        }
    }
    
    for i, v := range merged {
        nums1[i] = v
    }
}
