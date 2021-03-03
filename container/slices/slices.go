package main

import "fmt"

func updateSlice(s []int) {
	s[0] = 100
}

func main() {
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	fmt.Println("arr[2:6] =", arr[2:6])
	fmt.Println("arr[:6] =", arr[:6])
	s1 := arr[2:]
	fmt.Println("arr[2:] =", s1)
	s2 := arr[:]
	fmt.Println("arr[:] =", s2)

	fmt.Println("After updateSlice")
	updateSlice(s1)
	updateSlice(s2)
	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(arr)

	fmt.Println("ReSlice")
	s2 = s2[:5]
	fmt.Println(s2)
	s2 = s2[3:]
	fmt.Println(s2)

	fmt.Println("Extending slice")
	// slice可以向后扩展 不可以向前扩展
	fmt.Println("arr=", arr)
	s1 = arr[2:6]
	s2 = s1[3:5]
	fmt.Println("s1=", s1)
	fmt.Printf("s1=%v, len(s1)=%d, cap(s1)=%d\n", s1, len(s1), cap(s1))
	fmt.Println("s2=", s2)
	fmt.Printf("s2=%v, len(s2)=%d, cap(s2)=%d\n", s2, len(s2), cap(s2))

	s3 := append(s2, 10)
	s4 := append(s3, 11)
	s5 := append(s4, 12)
	// s4 s5不再是对arr的view。添加元素是如果超越cap，系统会重新分配更大的底层数组
	fmt.Println("s3, s4, s5 =", s3, s4, s5)
	fmt.Printf("s4=%v, len(s4)=%d, cap(s4)=%d\n", s4, len(s4), cap(s4))
	fmt.Println("arr =", arr)

}
