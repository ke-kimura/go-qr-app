package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	//コマンドラインから値を受け取る
	flag.Parse()
	arg := flag.Arg(0)
	msg := fmt.Sprintf("Hello %s\n", arg)

	//ファイルの作成
	f, err := os.Create("./hello.txt")
	if err != nil {
		fmt.Printf("failed to create file \n:%v", err)
		return
	}
	defer f.Close()

	//ファイルの書き込み

	_, err = f.WriteString(msg)
	if err != nil {
		fmt.Printf("failed to write message to file \n:%v", err)
	}

	fmt.Println("Done")
}
