package main

func main() {

	// break,continue,goto

	// for i := 1; i < 10; i++ {
	// 	if i == 2 {
	// 		break
	// 	}
	// 	fmt.Println(i)
	// }
	// fmt.Println("继续执行")

	// for i := 0; i < 2; i++ {
	// 	for j := 0; j < 3; j++ {
	// 		fmt.Printf("i=%v j=%v\n", i, j)
	// 		if j == 1 {
	// 			break
	// 		}
	// 	}
	// }

	// 再多重循环中，可以用标号 label 标出想 break 的循环。语法labelxxx

	// label1:
	// 	for i := 0; i < 2; i++ {
	// 		for j := 0; j < 3; j++ {
	// 			fmt.Printf("i=%v j=%v\n", i, j)
	// 			if j == 1 {
	// 				break label1
	// 			}
	// 		}
	// 	}

	// continue 跳过当前
	// for i := 1; i <= 10; i++ {
	// 	if i == 3 {
	// 		continue
	// 	}
	// 	fmt.Println(i)
	// }
	// label2:
	// 	for i := 0; i < 2; i++ {
	// 		for j := 0; j < 4; j++ {
	// 			if j == 2 {
	// 				continue label2
	// 			}
	// 			fmt.Printf("i=%v j=%v\n", i, j)
	// 		}
	// 	}

	// goto 跳转到指定标签

	// 	var n = 30
	// 	if n > 18 {
	// 		fmt.Printf("成年人\n")
	// 		goto label3
	// 	}

	// 	fmt.Println("aaa")
	// 	fmt.Println("bbb")
	// label3:
	// 	fmt.Println("ccc")
	// 	fmt.Println("ddd")

}
