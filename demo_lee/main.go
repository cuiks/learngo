package main

import (
	"fmt"
	"math/rand"
)

type Solution struct {
	w     []int
	total int
	sums  []int
}

func Constructor(w []int) Solution {
	total := 0
	sums := []int{0}
	for _, val := range w {
		total += val
		sums = append(sums, total)
	}
	return Solution{
		w:     w,
		total: total,
		sums:  sums,
	}

}

func (this *Solution) PickIndex() int {
	if len(this.w) == 0 {
		return 0
	}
	hit := rand.Intn(this.total)
	left := 0
	right := len(this.sums) - 1
	mid := 0
	for left < right-1 {
		mid = (left + right) / 2
		if hit <= this.sums[mid] {
			right = mid
		} else {
			left = mid
		}
	}
	return left

}

func main() {
	w := []int{3, 14, 1, 7}
	obj := Constructor(w)
	for i:=0;i<10;i++{
		fmt.Println(obj.PickIndex())
	}

}
