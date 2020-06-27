package main

func main() {

	// 整形與整形
	// var a int8 = 20
	// var b int16 = 40
	// fmt.Println(int16(a) + b)

	// 浮點型與浮點型
	// var a float32 = 20
	// var b float64 = 40
	// fmt.Println(float64(a) + b)

	// 整形和浮點型轉換
	// var a float32 = 20.33
	// var b int = 40
	// fmt.Println(a + float32(b))

	// 注意： 轉換的時候建議從低位轉換成高位， 高位轉換成低位的時候轉換不成功就會溢出，和我們想的結果不一樣

	// var a int8 = 20
	// var b int16 = 140 // 轉換成int16

	// 其他類型轉換成string類型
	// 注意： Sprintf 使用中需要注意轉換的格式 int 為%d float 為%f bool 為%t byte 為%c
	// var i int = 20
	// var f float64 = 12.456
	// var t bool = true
	// var b byte = 'a'

	// str1 := fmt.Sprintf("%d", i)
	// fmt.Printf("值： %v 類型：%T\n", str1, str1)    // 值： 20 類型：string

	// str2 := fmt.Sprintf("%f", f)
	// fmt.Printf("值： %v 類型： %T\n", str2, str2)   // 值： 12.456000 類型： string

	// str3 := fmt.Sprintf("%t", t)
	// fmt.Printf("值： %v  類型：%T\n", str3, str3)   // 值： true  類型：string

	// str4 := fmt.Sprintf("%c", b)
	// fmt.Printf("值： %v 類型：%T\n", str4, str4)    // 值： a 類型：string

	// 通過 strconv

}
