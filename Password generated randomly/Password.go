package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"time"

	"github.com/atotto/clipboard"
)

func main() {
	var length int
	fmt.Print("请输入密码长度 (直接回车默认为 16): ")
	// 处理回车默认值
	_, err := fmt.Scanln(&length)
	if err != nil || length <= 0 {
		length = 16
	}

	// 排除易混淆字符: l, 1, I, o, 0, O
	charset := "abcdefghijkmnpqrstuvwxyzABCDEFGHJKLMNPQRSTUVWXYZ23456789!@#$%^&*"

	password, err := generatePassword(length, charset)
	if err != nil {
		fmt.Println("生成失败:", err)
		return
	}

	// 稳健的展示方式
	fmt.Println("\n========================================")
	fmt.Printf("生成的密码: %s\n", password)
	fmt.Println("========================================")

	// 复制到剪贴板
	err = clipboard.WriteAll(password)
	if err != nil {
		fmt.Println("复制到剪贴板失败:", err)
	} else {
		fmt.Println("✅ 密码已自动复制到剪贴板！")
		fmt.Println("⚠️  出于安全考虑，剪贴板将在 60 秒后自动清空...")
		
		go func(p string) {
			time.Sleep(60 * time.Second)
			current, _ := clipboard.ReadAll()
			if current == p {
				clipboard.WriteAll("")
				fmt.Println("\n[系统消息] 剪贴板已安全清空。")
			}
		}(password)
	}

	fmt.Println("\n按回车键退出程序...")
	fmt.Scanln()
}

func generatePassword(length int, charset string) (string, error) {
	password := make([]byte, length)
	for i := range password {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		if err != nil {
			return "", err
		}
		password[i] = charset[num.Int64()]
	}
	return string(password), nil
}