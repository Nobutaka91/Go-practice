package main

import (
	"io"
	"log"
	"os"
)

func main() {
	file, err := os.Create("log.txt") // ファイル生成
	if err != nil {
		log.Fatalln(err) // エラーのログを出力した後にプログラムを強制終了する
	}
	defer file.Close() // deferで関数の終了時するタイミングで実行される
	flags := log.Lshortfile // ログにファイル名と行番号を含める
	warnLogger := log.New(io.MultiWriter(file, os.Stderr), "WARN: ", flags)
	errorLogger := log.New(io.MultiWriter(file, os.Stderr), "ERROR: ", flags)

	warnLogger.Println("warning A")

	errorLogger.Fatalln("critical error") // Fatalln: プログラムを強制終了する
}



