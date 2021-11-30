package utils

import (
	"fmt"
)

// Config parameters name mapping to config.toml must be capitalize
type Config struct {
	File   map[string]string
	Common map[string]string
	DB     map[string]string
}

// GetConfig get config by viper
func GetConfig() Config {
	var config Config
	jsonMap, err := ReadJSONConfig("./configs/config.json")
	if err != nil {
		panic(fmt.Errorf("read configs/config.json error: %s", err))
	}
	// reset sub config type
	reqItemList := []string{"common", "db", "file"}
	for _, item := range reqItemList {
		if jsonMap[item] != nil {
			subConfig := jsonMap[item].(map[string]interface{})
			newSubConfig := make(map[string]string)
			for key := range subConfig {
				newSubConfig[key] = subConfig[key].(string)
			}
			jsonMap[item] = newSubConfig
		} else {
			if HasString(reqItemList, item) {
				panic("config.json missing require key: " + item)
			} else {
				jsonMap[item] = make(map[string]string)
			}
		}
	}
	// set form config
	config = Config{
		Common: jsonMap[reqItemList[0]].(map[string]string),
		DB:     jsonMap[reqItemList[1]].(map[string]string),
		File:   jsonMap[reqItemList[2]].(map[string]string),
	}
	return config
}
