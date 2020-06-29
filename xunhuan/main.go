package main

func main() {
	// var str = "你好，golang"

	// for k, v := range str {
	// 	fmt.Printf("key=%v, var=%c\n", k, v)
	// }

	// key=0, var=你
	// key=3, var=好
	// key=6, var=，
	// key=9, var=g
	// key=10, var=o
	// key=11, var=l
	// key=12, var=a
	// key=13, var=n
	// key=14, var=g

	// for ..i.. 和 for...range...
	// var arr = []string{"php", "java", "node", "golang"}
	// for i := 0; i < len(arr); i++ {
	// 	fmt.Println(arr[i])
	// }

	// php
	// java
	// node
	// golang

	// var arr = []string{"php", "java", "node", "golang"}
	// for _, str1 := range arr {
	// 	fmt.Println(str1)
	// }

	// switch...case 和 if...else
	// 判断文件类型
	// var ext = ".htmljf"

	// if ext == ".html" {
	// 	fmt.Println("text/html")
	// } else if ext == ".css" {
	// 	fmt.Println("text/css")
	// } else if ext == ".js" {
	// 	fmt.Println("text/javascript")
	// } else {
	// 	fmt.Println("找不到后缀")
	// }

	// switch ext {                    // 也可以是 switch ext := ".html"; ext {
	// case ".html":
	// 	fmt.Println("text/html")
	// 	break
	// case ".css":
	// 	fmt.Println("text/css")
	// 	break
	// case ".js":
	// 	fmt.Println("text/javascript")
	// 	break
	// default:
	// 	fmt.Println("找不到此后缀")
	// }

	// 一个分支可以有多个值，多个case 值中间使用英文逗号分隔
	// 判断一个数是不是偶数

	// var n = 5
	// switch n {
	// case 1, 3, 5, 7, 9:
	// 	fmt.Println("此数是奇数")
	// case 0, 2, 4, 6, 8:
	// 	fmt.Println("此数是偶数")
	// }

	// switch什么都不需要   case后可跟表达式

	// switch 的穿透 fallthrough,只会穿透紧挨着的下一层
	// var age = 30
	// switch {
	// case age < 18:
	// 	fmt.Println("好好学习")
	// case age >= 18 && age < 24:
	// 	fmt.Println("娶白富美")
	// case age >= 24 && age < 40:
	// 	fmt.Println("努力花钱")
	// 	fallthrough
	// case age >= 40 && age < 60:
	// 	fmt.Println("注意身体")
	// default:
	// 	fmt.Println("安详晚年")
	// }

}
