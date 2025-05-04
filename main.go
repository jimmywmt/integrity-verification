package main

import (
	"fmt"
	"log"
	"os"

	"github.com/jimmywmt/integrity-verification/tools"
)

// ===== 密鑰參數（與主程式相同）=====
// 編譯前請使用sed命令替換這兩行
const Password = "{{PASSWORD}}"
const SaltHex = "{{SALT_HEX}}"

func main() {
	println("=== 產生 token.chk 對應檔案 ===")

	if len(os.Args) < 2 {
		fmt.Println("用法：generate-token <binary_path>")
		os.Exit(1)
	}

	// 路徑請放在 /data 目錄下 (請在docker run 時掛載 /data 目錄)
	binaryPath := os.Args[1]
	token := tools.FileToken(binaryPath, Password, SaltHex)
	// 將 token 寫入 /data/token.chk (請在docker run 時掛載 /data 目錄)
	outputPath := "/data/token.chk"
	if err := os.WriteFile(outputPath, []byte(token), 0644); err != nil {
		log.Fatalf("寫入 %s 失敗: %v", outputPath, err)
	}

	fmt.Println("✅ 成功產生 token.chk 對應檔案:", outputPath)
}
