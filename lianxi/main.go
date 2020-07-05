package main

import "fmt"

func main() {

	// 1.请求一个数组里面元素的和以及这些元素的平均值，分别用for和for...range实现
	// var arr1 = [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	// num := 0
	// for i := 0; i < len(arr1); i++ {
	// 	num += i
	// }
	// fmt.Printf("平均值：%.2f", float64(num)/float64(len(arr1)))

	// 2. 请求出一个数组的最大值，并得到对应的下标

	// var arr1 = [...]int{1, 2, 3, 4, 5, 666, -1, 700, 8, 9, 100}
	// max := arr1[0]
	// index := 0
	// for i := 0; i < len(arr1); i++ {
	// 	if arr1[i] > max {
	// 		max = arr1[i]
	// 		index = i
	// 	}
	// }
	// fmt.Printf("最大值：%v  下标：%v", max, index)

	// 3、从数组{1，3，5，7，8}中找出和为8的元素的下标(0,3) 和 (1,2)

	var arr = [...]int{1, 3, 5, 7, 8}
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i]+arr[j] == 8 {
				fmt.Printf("(%v,%v)", i, j)
			}
		}
	}
}
