package calculator

import "fmt"

var offset float64 = 1 // 小文字始まりはパッケージ内のみで使用できる
var Offset float64 = 1 //　大文字始まりは他のパッケージからも使用できる

func Sum(a float64, b float64) float64 {
	fmt.Println("multiply: ", multiply(a, b)) // 同じパッケージ内の小文字始まりの関数も使用できる
	fmt.Println("Multiply: ", Multiply(a, b)) // 同じパッケージ内の大文字始まりの関数も使用できる
	return a + b + offset
}