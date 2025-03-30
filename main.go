package main

import (
	"Fank/cmd"
	_ "Fank/internal/logger" // 确保导入 logger 包
)

func main() {
	cmd.Start()
}
