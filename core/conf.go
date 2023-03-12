package core

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"gvb-server/config"
	"gvb-server/global"
	"io/fs"
	"io/ioutil"
	"log"
)

const ConfigFile = "settings.yaml"

// InitConf is the function to initialize the configuration in yaml file.
func InitConf() {
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

func SetYaml() error {
	bytes, err := yaml.Marshal(global.Config) // 将配置文件转换为yaml格式
	if err != nil {
		global.Log.Error(err)
		return err
	}
	err = ioutil.WriteFile(ConfigFile, bytes, fs.ModePerm) // 将配置文件写入到yaml文件中
	if err != nil {
		global.Log.Error(err)
		return err
	}
	global.Log.Infof("配置文件修改成功")
	return nil
}
