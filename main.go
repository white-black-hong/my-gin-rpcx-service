package main

import (
	"fmt"
	config "my-service/global"
)

var (
	configYml string
)

func init() {
	// 读取配置
	configYml = "config.yaml"
	config.Setup(configYml)
	fmt.Println("配置读取成功")
}

func main() {

}
