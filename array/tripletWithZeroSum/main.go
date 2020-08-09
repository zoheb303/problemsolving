package main

import (
	"fmt"
	"sort"
)

func tripletSum(a []int) [][]int {

	var res2d [][]int
	sort.Ints(a)
	fmt.Println(a)
	n := len(a) - 1
	for i := 0; i <= n-2; i++ {
		if i > 0 && a[i] == a[i-1] {
			continue
		}

		left := i + 1
		right := n
		fmt.Println(left, i, right)
		for left < right {
			var res []int

			if a[left]+a[i]+a[right] < 0 {

				left++

			} else if a[left]+a[i]+a[right] > 0 {

				right--
			} else if a[left]+a[i]+a[right] == 0 {

				res = append(res, a[i], a[left], a[right])

				left++
				right--
				for left < right && a[left] == a[left-1] {
					left++
				}
				for left < right && a[right] == a[right+1] {
					right--
				}
			}
			if len(res) != 0 {
				res2d = append(res2d, res)
			}
		}

	}
	return res2d
}

func main() {

	var arr = []int{-2, 0, 0, 2, 2}
	//var arr = []int{0, 0, 0}

	sort.Ints(arr)

	a := tripletSum(arr)

	fmt.Println(a)

}
