package main

import "fmt"

func main() {
	// * 基本構文
	fmt.Println("Hello World")

	// * 変数宣言
	// var 変数名 型
	// int
	var num1 int = 10;
	var num2 int = 20;

	fmt.Println((num1 * num2))

	// * String
	var str string = "Goの基本文法を学ぶ"
	fmt.Println(str)

	// * bool
	var flg bool = true
	if(flg) {fmt.Println("bool値:確認")}

	// * float32: 32bit 約7桁の精度
	// * float64:　　64bit 約15桁の精度
	var r float32 = 3.14;
	fmt.Println((r))


	fmt.Println(("hoge"))

	var result float32;
	result = (3 * 5 ) / 2
	fmt.Println(result); 


	// * 条件分岐
	// 閏年判定
	// 1. 4で割り切れる
	// 2. 100で割り切れる場合は閏年ではない
	// 3. 400で割り切れる場合は閏年
	var num int = 2000;
	if(num % 4 == 0 && num % 400 ==0 || num % 100 != 0) {
		fmt.Println(("閏年"))
	} else {
		fmt.Println(("閏年ではない"))
	}
	
	// * switch文
	var n int = 1;
	var day string;
	switch(n) {
	case 1:
		day = "月"
		break;
	case 2:
		day = "火"
		break;
	default:
		day = "土"
	}
	fmt.Println((day))

	for i:=0; i < 100; i++ {
		if(i%15==0) {
			fmt.Println("FizzBuzz");
			} else if(i%3==0) {
			fmt.Println("Fizz");
			} else if(i % 5 ==0){
				fmt.Println("Buzz");
		} else {
			fmt.Println((i));
		}
	}

	var weight float32 = 48.0;
	var height float32 = 1.530;

	result = getBmi(weight, height);
	fmt.Println((result))

	// 配列
	scores := [3]int{2, 3, 5};
	for i, score := range scores {
		fmt.Println((i))
		fmt.Println((score))
	} 

}

func getBmi(val1 float32, val2 float32) float32{
	return val1 / (val2 * val2);
}