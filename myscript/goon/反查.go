package script

import (
	"bufio"
	"fmt"
	"github.com/mozillazg/go-pinyin"
	"os"
	"strings"
)

func ParseFile(filePath string) map[string][]string {
	// 创建一个空的映射用于存储字和部分的拼音
	result := make(map[string][]string)

	// 打开文件
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
		return result
	}
	defer file.Close()

	// 使用带缓冲的读取器逐行读取文件内容
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		// 忽略以#开头的行和空行
		if strings.TrimSpace(line) != "" && !strings.HasPrefix(line, "#") {
			// 按制表符或空格分割每行，获取字和部分的拼音
			parts := strings.Split(line, "\t")
			if len(parts) == 2 {
				character := parts[0]
				pinyin := parts[1]
				// 将字和部分的拼音添加到映射中
				result[character] = append(result[character], pinyin)
			}
		}
	}

	// 检查扫描过程中是否有错误
	if err := scanner.Err(); err != nil {
		fmt.Println("Error scanning file:", err)
	}

	return result
}

func pinyinit(in string) string {
	in = strings.Replace(in, " ", "", -1)
	a := pinyin.NewArgs()
	outsl := pinyin.Pinyin(in, a)
	if len(outsl) == 0 {
		return "no result for it"
	}
	outs := []string{}
	for _, v := range outsl {
		outs = append(outs, v[0])
	}
	out := strings.Join(outs, " ")
	return out
	// [[zhong] [guo] [ren]]
}
