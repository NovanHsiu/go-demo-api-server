package common

import (
	"encoding/json"
	"os"
)

type ConfigCommon struct {
	Port       string `json:"port"`
	SslPort    string `json:"ssl_port"`
	TlsCrtPath string `json:"tls_crt_path"`
	TlsKeyPath string `json:"tls_key_path"`
}

type ConfigFile struct {
	StaticFileDir string `json:"static_file_dir"`
}

type ConfigDB struct {
	Type     string `json:"type"`
	Host     string `json:"host"`
	Port     string `json:"port"`
	Dbname   string `json:"dbname"`
	User     string `json:"user"`
	Passwd   string `json:"passwd"`
	Sslmode  string `json:"sslmode"`
	Timezone string `json:"timezone"`
}

// Config parameters name mapping to config.json must be capitalize
type Config struct {
	File   ConfigFile
	Common ConfigCommon
	DB     ConfigDB
}

// GetConfig get config by viper
func GetConfig() Config {
	config := Config{}
	dataByte, err := os.ReadFile(GetExecutionDir() + "/configs/config.json")
	if err != nil {
		panic(err)
	}
	json.Unmarshal(dataByte, &config)
	return config
}
