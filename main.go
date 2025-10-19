package main

import (
	"fmt"
	"go-basics/calculator"
	"os"

	"github.com/joho/godotenv"
)


func main() {
	godotenv.Load() // .envファイルを読み込む
	fmt.Println(os.Getenv("GO_ENV")) // 環境変数を取得して表示
	fmt.Println(calculator.Offset) // calculatorパッケージの大文字始まりの変数にアクセス
	fmt.Println(calculator.Sum(1, 2)) // calculatorパッケージのSum関数を使用
	fmt.Println(calculator.Multiply(3, 4)) // calculatorパッケージのMultiply関数を使用
}
