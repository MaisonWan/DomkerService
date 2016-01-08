package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type configValue struct {
	Key   string
	Value string
}

type Config struct {
	configData map[string]interface{}
}

// 构造函数
func GetConfig() *Config {
	config := Config{}
	if result, err := loadConfig(); err == nil {
		config.configData = result
	} else {
		fmt.Println("GetConfig Error", err)
	}
	return &config
}

func loadConfig() (result map[string]interface{}, err error) {
	configData, readErr := ioutil.ReadFile("notification/config/config.json")
	if readErr != nil {
		fmt.Println("Can not open config file", readErr)
		return
	}
	err = json.Unmarshal(configData, &result)
	if err != nil {
		return result, err
	}
	return result, nil
}

func (c *Config) GetEmailUsername() string {
	return c.configData["email.username"].(string)
}

func (c *Config) GetEmailPassword() string {
	return c.configData["email.password"].(string)
}

func (c *Config) GetEmailHost() string {
	return c.configData["email.host"].(string)
}

func (c *Config) GetEmailReceiver() string {
	return c.configData["email.receiver"].(string)
}

func (c *Config) GetDoubanUser() string {
	return c.configData["douban.user"].(string)
}
