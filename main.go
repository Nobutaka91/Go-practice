package main

import (
	"fmt"
)

const secret = "abc" // 定数secretを宣言

type Os int // Os型を定義

const ( // iotaを使った定数宣言
	Windows Os = iota + 1
	MacOS
	Linux
)

var ( // 複数の変数をまとめて宣言
	i int
	s string
	b bool
)

func main() {


	i := 2
	ui := uint16(2) // uint16型の変数uiを宣言し、2を代入
	fmt.Println(i)
	fmt.Printf("i: %v %T\n", i, i ) // 変数iの値と型を表示
	fmt.Printf("i: %[1]v %[1]T ui: %[2]v %[2]T\n", i, ui) // 変数iとuiの値と型を表示

	f := 1.23456
	s := "hello"
	b := true
	fmt.Printf("f: %[1]v %[1]T\n", f) // 変数fの値と型を表示
	fmt.Printf("s: %[1]v %[1]T\n", s) // 変数sの値と型を表示
	fmt.Printf("b: %[1]v %[1]T\n", b) // 変数bの値と型を表示
	
	pi, title := 3.14, "Go"
	fmt.Printf("pi: %v title: %v\n", pi, title) // 変数piとtitleの値を表示

	x := 10
	y := 1.23
	z := float64(x) + y
	fmt.Println(z) // 変数xとyの値を表示

	var ui1 uint16 // uint16型の変数ui1を宣言
	fmt.Printf("memory address of ui1: %p\n", &ui1) // 変数ui1のメモリアドレスを表示
	var ui2 uint16 
	fmt.Printf("memory address of ui2: %p\n", &ui2)
	var p1 *uint16 // ポインタ型の変数p1を宣言
	fmt.Printf("value of p1: %v\n", p1) // 変数p1の値を表示（nilになる）
	p1 = &ui1 // 変数p1に変数ui1のアドレスを代入
	fmt.Printf("value of p1: %v\n", p1) // 変数p1の値を表示（ui1のアドレスになる）

	// shadowing
	ok, result := true, "A"
	fmt.Printf("memory address of result: %p\n", &result) // 変数resultのメモリアドレスを表示
	if ok {
		result := "B"
		fmt.Printf("memory address of result: %p\n", &result) // ブロック内の変数resultのメモリアドレスを表示
		println(result) // ブロック内のresultを表示（"B"）
	} else {
		result := "C"
		println(result) 
	}
	println(result) // ブロック外のresultを表示（"A"）

}
