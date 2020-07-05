package main

func main() {

	// 数组的长度是类型的一部分
	/*
		var arr1 [3]int
		var arr2 [4]int
		var strArr [3]string

		fmt.Printf("arr1:%T arr2:%T strArr:%T", arr1, arr2, strArr)
	*/

	// 数组的初始化 第一种方法
	/*var arr1 [3]int

	arr1[0] = 23
	arr1[1] = 10
	arr1[2] = 100
	fmt.Println(arr1)
	*/

	// 第二种,初始化直接赋值
	// var arr1 = [3]int{1, 2, 3} // arr1 := [3]int{1, 2, 3}
	// fmt.Println(arr1)

	// 第四种，自行推断
	/*var arr1 = [...]int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(arr1)
	*/
	// 数组
	/*arr1 := [...]string{"php", "java", "golang", "js"}
	arr1[3] = "Python"
	fmt.Println(arr1)
	*/

	// 第四种方法
	/*arr := [...]int{0: 1, 1: 10, 2: 100}
	fmt.Println(arr)
	*/

	// 数组循环遍历
	// var arr1 = [3]int{23, 34, 5}
	/*for i := 0; i < len(arr1); i++ {
		fmt.Println(arr1[i])
	}
	*/
	// arr1 := [...]string{"php", "java", "golang", "js"}
	// for k, v := range arr1 {
	// 	fmt.Printf("key:%v  value:%v\n", k, v)
	// }
}
