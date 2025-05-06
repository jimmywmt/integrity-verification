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

	if len(os.Args) < 3 {
		fmt.Println("用法：generate-token <binary_path> <machine_identifier>")
		os.Exit(1)
	}

	binaryPath := os.Args[1]
	token := tools.FileToken(binaryPath, Password, SaltHex)
	outputPath := "./token.chk"
	if err := os.WriteFile(outputPath, []byte(token), 0644); err != nil {
		log.Fatalf("寫入 %s 失敗: %v", outputPath, err)
	}

	fmt.Println("✅ 成功產生 token.chk 對應檔案:", outputPath)

	license := tools.StringToken(os.Args[2], Password, SaltHex)
	outputPath = "./license.chk"
	if err := os.WriteFile(outputPath, []byte(license), 0644); err != nil {
		log.Fatalf("寫入 %s 失敗: %v", outputPath, err)
	}

	fmt.Println("✅ 成功產生 license.chk 對應檔案:", outputPath)
}
