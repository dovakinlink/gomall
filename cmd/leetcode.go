package main

import "fmt"

//nums1 := []int{1, 3, 4, 6, 7}
//nums2 := []int{2, 9, 10, 15}
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	i1, i2, m, n := 0, 0, len(nums1), len(nums2)
	mk := map[int]int{}
	for i := 0; i < m+n; i++ {
		if i1 >= m {
			mk[i] = nums2[i2]
			i2++
			continue
		}
		if i2 >= n {
			mk[i] = nums1[i1]
			i1++
			continue
		}
		if nums1[i1] < nums2[i2] {
			mk[i] = nums1[i1]
			i1++
		} else {
			mk[i] = nums2[i2]
			i2++
		}
	}
	if (m+n)%2 == 0 {
		k1 := (m + n) / 2
		k2 := k1
		k1 -= 1
		f := float64(mk[k1]) + float64(mk[k2])
		return f / 2
	} else {
		return float64(mk[(m+n)/2])
	}
}

func main() {
	nums1 := []int{1, 3, 4, 6, 7, 17, 18}
	nums2 := []int{2, 9, 10, 15, 16}
	f := findMedianSortedArrays(nums1, nums2)
	fmt.Println(f)
}
