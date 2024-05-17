package fileutil

import (
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

// WriteYAML 将 data 结构体写入到指定的YAML文件中
func WriteYAML(filePath string, data any) error {
	// 确保目录存在
	dir := filepath.Dir(filePath)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		if err := os.MkdirAll(dir, 0755); err != nil {
			return err
		}
	}

	// 将data编码为YAML
	body, err := yaml.Marshal(data)
	if err != nil {
		return err
	}

	// 写入文件
	return os.WriteFile(filePath, body, 0644)
}

// ReadYAML 从指定的YAML文件中读取data结构体
func ReadYAML(filePath string) (any, error) {
	var data any

	// 读取文件内容
	body, err := os.ReadFile(filePath)
	if err != nil {
		return body, err
	}

	// 将YAML解码为data结构体
	err = yaml.Unmarshal(body, &data)
	if err != nil {
		return data, err
	}

	return data, nil
}

// HomeConfigPath 获取当前用户的 .config 目录
// $HOME/.config/
func HomeConfigPath() string {
	home, err := os.UserHomeDir()
	if err != nil {
		return ""
	}
	return filepath.Join(home, ".config")
}
