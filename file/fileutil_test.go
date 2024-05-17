package fileutil

import (
	"fmt"
	"testing"
)

func TestCover(t *testing.T) {
	err := CoverFile("test.txt", ",]", "]")
	fmt.Println(err)
}

func TestYaml(t *testing.T) {
	// Config 示例配置结构体
	type Config struct {
		Redis struct {
			IP   string `yaml:"ip"`
			Port int    `yaml:"port"`
		} `yaml:"redis"`
	}
	// 示例配置
	config := Config{
		Redis: struct {
			IP   string `yaml:"ip"`
			Port int    `yaml:"port"`
		}{
			IP:   "1.1.1.1",
			Port: 6379,
		},
	}

	// 写入YAML文件
	err := WriteYAML(HomeConfigPath()+"/example.yaml", config)
	if err != nil {
		fmt.Println("Error writing YAML:", err)
		return
	}
	fmt.Println("YAML file written successfully.")

	// 读取YAML文件
	readConfig, err := ReadYAML("example.yaml")
	if err != nil {
		fmt.Println("Error reading YAML:", err)
		return
	}
	fmt.Printf("Read config: %+v\n", readConfig)

}
