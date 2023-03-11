package config

import "strconv"

// System is the configuration for the system.
type System struct {
	SrvName string `yaml:"srv_name"`
	Host    string `yaml:"host"`
	Port    int    `yaml:"port"`
	Env     string `yaml:"env"`
}

// GetSrvName 返回服务名称
func (s *System) GetSrvName() string {
	return s.SrvName
}

// Addr 返回系统地址
func (s *System) Addr() string {
	return s.Host + ":" + strconv.Itoa(s.Port)
}
