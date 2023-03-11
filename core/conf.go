package core

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"gvb-server/config"
	"gvb-server/global"
	"io/ioutil"
	"log"
)

// InitConf is the function to initialize the configuration in yaml file.
func InitConf() {
	const ConfigFile = "settings.yaml"
	c := &config.Config{}
	yamlConf, err := ioutil.ReadFile(ConfigFile)
	if err != nil {
		panic(fmt.Errorf("get yamlConf err: %s", err))
	}
	err = yaml.Unmarshal(yamlConf, c)
	if err != nil {
		log.Fatalf("config Init Unmarshal: %v", err)
	}
	log.Println("config yamlFile load Init success.")
	// 将配置文件赋值给全局变量
	global.Config = c
}
