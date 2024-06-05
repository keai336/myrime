package main

import (
	"fmt"
	script "myscript/myscript/goon"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
)

// 反查
func main() {

	filePath := "C:\\Users\\keai3\\AppData\\Roaming\\Rime\\radical_pinyin.dict.yaml" // 替换为你的文本文件路径

	// 获取当前工作目录
	//fmt.Println("绝对路径:", filePath)
	result := script.ParseFile(filePath)
	//fmt.Println(result)
	var input string
	go func() {
		for {
			fmt.Scanln(&input)
			if rs, ok := result[input]; ok {
				fmt.Println(input, ":", strings.Join(rs, "\t"))
			} else {
				fmt.Println("no result for it")
			}

		}

	}()
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c
	fmt.Println("thanks", time.Now())
	os.Exit(1)
}
