package main

func main() {

	//1、golang中定義字符， 字符屬於int類型

	// var a = 'a'
	// fmt.Printf("值： %v  類型:%T", a, a) //輸出碼值97 int32

	// 2、原樣輸出
	// var a = 'a'
	// fmt.Printf("值： %c 類型：%T", a, a) // a  int32

	// Go 語言中字符有兩种
	// 1. uint8類型， 或者叫 byte型， 代表了ASCII碼的一個字符
	// 2. rune 類型， 代表了一個UTF-8字符。

	// 3、定義一個字符串輸出字符串裡面的字符
	// var str = "this"
	// fmt.Printf("值：%v 原樣輸出：%c  類型：%T", str[2], str[2], str[2]) //值：105 原樣輸出：i  類型：uint8

	// unsafe.Sizeof() 沒法查看string類型數據所占用的存儲空間

	// 4、循環輸出字符串裡面的數字
	// byte
	// s := "你好 golang"   //228(ä)189(½)160( )229(å)165(¥)189(½)32( )103(g)111(o)108(l)97(a)110(n)103(g)
	// s := "golang"        //103(g)111(o)108(l)97(a)110(n)103(g)
	// for i := 0; i < len(s); i++ {
	// 	fmt.Printf("%v(%c)", s[i], s[i])
	// }

	// range
	// s := "你好 golang"
	// for _, v := range s {
	// 	fmt.Printf("%v(%c)", v, v) //20320(你)22909(好)32( )103(g)111(o)108(l)97(a)110(n)103(g)
	// }

	// 5、修改字符串
	// s1 := "big"
	//強制類型轉換,適用於沒有漢字的字符串中
	// s1[0] = 'p'
	// fmt.Printf(s1) // 報錯，字符串不可直接轉換
	// byteStr := []byte(s1)
	// byteStr[0] = 'p'
	// fmt.Println(string(byteStr))

	// s1 := "你好 golang"
	// runeStr := []rune(s1)
	// runeStr[0] = 'p'
	// fmt.Println(string(runeStr))

}
