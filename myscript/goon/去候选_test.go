package script

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"regexp"
	"strings"
	"testing"

	"github.com/syndtr/goleveldb/leveldb"
)

type CommandResult struct {
	Output string
	Error  error
}

func executeBatFile(command string) CommandResult {
	// 准备执行的命令

	cmd := exec.Command("cmd.exe", "/C", command)

	// 执行命令并捕获输出
	output, err := cmd.CombinedOutput()

	// 返回结果结构体
	return CommandResult{
		Output: string(output),
		Error:  err,
	}
}

func TestIt(t *testing.T) {
	dbPath := "C:\\Users\\keai3\\AppData\\Roaming\\Rime\\rime_ice.userdb"

	// 打开数据库

	// 输入字符串，前一部分和后一部分
	var input string
	r := regexp.MustCompile("^[^\t]+\t[^\t]+$") // 前一部分是 "A"，后一部分是 "I"
	fmt.Println("输入")
	for {
		//fmt.Scanln(&input)
		input = "zhen xin	振新"
		if r.MatchString(input) {
			break
		}
		fmt.Println("错误输入,重新输入")
	}
	rs := executeBatFile("D:\\tools\\Rime\\weasel-0.16.0\\stop_service.bat")
	if rs.Error != nil {
		panic(rs.Error)

	}
	fmt.Println(rs.Output)
	// 分割输入字符串
	db, err := leveldb.OpenFile(dbPath, nil)
	if err != nil {
		log.Fatalf("Failed to open LevelDB: %v", err)
	}
	defer db.Close()
	parts := strings.Split(input, "\t")
	if len(parts) != 2 {
		log.Fatalf("Invalid input format. Expected format: <prefix>\t<suffix>")
	}
	prefix := parts[0]
	suffix := parts[1]

	// 创建一个迭代器来遍历数据库中的键值对
	iter := db.NewIterator(nil, nil)
	defer iter.Release()

	// 创建一个读取器来读取用户输入
	reader := bufio.NewReader(os.Stdin)

	for iter.Next() {
		key := iter.Key()
		value := iter.Value()

		// 分割键值对中的键和候选词
		keyParts := bytes.SplitN(key, []byte("\t"), 2)
		if len(keyParts) != 2 {
			continue
		}
		keyPrefix := string(keyParts[0])
		keySuffix := string(keyParts[1])

		// 前一部分完全匹配，后一部分模糊匹配
		if keyPrefix == prefix && strings.Contains(keySuffix, suffix) {
			// 打印匹配的键值对
			fmt.Printf("Matched Key: %s, Value: %s\n", key, value)
			fmt.Print("Do you want to delete this key-value pair? (y/n): ")

			// 读取用户输入
			input, err := reader.ReadString('\n')
			if err != nil {
				log.Fatalf("Failed to read input: %v", err)
			}
			input = strings.TrimSpace(input)

			// 检查用户输入
			if input == "y" || input == "Y" {
				// 删除匹配的键值对
				err := db.Delete(key, nil)
				if err != nil {
					log.Fatalf("Failed to delete key: %v", err)
				}
				fmt.Println("Key-value pair deleted.")
			} else {
				fmt.Println("Key-value pair not deleted.")
			}
		}
	}

	// 检查迭代器是否有错误
	if err := iter.Error(); err != nil {
		log.Fatalf("Iterator error: %v", err)
	}
	executeBatFile("D:\\tools\\Rime\\weasel-0.16.0\\start_service.bat")
}

func TestBat(t *testing.T) {
	// 执行的bat文件路径
	batFilePath := "D:\\tools\\Rime\\weasel-0.16.0\\stop_service.bat"

	// 准备执行bat文件的命令
	cmd := exec.Command(batFilePath)

	// 执行命令并捕获输出
	output, err := cmd.CombinedOutput()
	if err != nil {
		// 打印错误信息
		fmt.Println("Error executing batch script:", err)
		return
	}

	// 打印执行结果
	fmt.Println("Batch script output:", string(output))
}
