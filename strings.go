package main

func main() {

	// const (
	// 	n1, n2 = iota + 1, iota + 2
	// 	n3, n4
	// )
	// fmt.Println(n1, n2, n3, n4)

	//求長度
	// var str1 = "你好"
	// fmt.Println(len(str1)) //一個漢字三字符
	// var str2 = "abcd"
	// fmt.Println(len(str2))

	//拼接字符串 + 或者 Sprintf
	// str1 := "你好"
	// str2 := "golang"
	// // // str3 := str1 + str2
	// // // fmt.Println(str3)
	// str3 := fmt.Sprintf("%v---%v", str1, str2)
	// fmt.Println(str3)

	//切片
	// var str1 = "123-456-789"
	// arr := strings.Split(str1, "-")
	// fmt.Println(arr) //[123 456 789]

	//string.Join join 操作 表示把切片鏈接成字符串
	// var str1 = "123-456-789"
	// arr := strings.Split(str1, "-")
	// str2 := strings.Join(arr, "*")
	// fmt.Println(str2)

	// arr := []string{"php", "java", "golang"}
	// fmt.Println(arr) //[php java golang]
	// str3 := strings.Join(arr, "-")
	// fmt.Println(str3) //php-java-golang

	//strings.Contains(s, substr)判斷是否包含
	// str1 := "this is str"
	// str2 := "this"
	// flag := strings.Contains(str1, str2)
	// fmt.Println(flag) //true

	// strings.HasPrefix(s, prefix),strings.HasSuffix(s, suffix) 前綴、後綴判斷
	// str1 := "this is str"
	// str2 := "this"
	// flag := strings.HasPrefix(str1, str2)
	// fmt.Println(flag) //true

	// strings.Index(s, substr),strings.LastIndex(s, substr) 子串出現的位置，查找不到返回-1，找到返回下標，從0開始的。
	// 	str1 := "this is str"
	// 	str2 := "is"
	// 	// num := strings.Index(str1, str2)
	// 	// print(num) //this 2
	// 	num := strings.LastIndex(str1, str2) //從左往右最後出現的位置
	// 	print(num)                           //this 5

}
