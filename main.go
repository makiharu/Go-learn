package main

import "fmt"

func main() {
	// 基本構文
	fmt.Println("Hello World")

	// 変数宣言
	// var 変数名 型
	var num1 int = 10;
	var num2 int = 20;

	fmt.Println((num1 * num2))

	// String
	var str string = "Goの基本文法を学ぶ"
	fmt.Println(str)

	// bool
	var flg bool = true
	if(flg) {fmt.Println("bool値:確認")}

	// float32: 32bit 約7桁の精度
	// float64:　　64bit 約15桁の精度
	var r float32 = 3.14;
	fmt.Println((r))


}