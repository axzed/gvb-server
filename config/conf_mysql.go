package config

import "strconv"

// Mysql is the configuration for the mysql database.
type Mysql struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Config   string `yaml:"config"` // 高级配置 例如: charset
	DB       string `yaml:"db"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	LogLevel string `yaml:"log_level"` // 日志等级 debug为输出全部SQL, dev, release线上环境
}

// Dsn is the function to get the dsn of the mysql database.
func (m *Mysql) Dsn() string {
	return m.User + ":" + m.Password + "@tcp(" + m.Host + ":" + strconv.Itoa(m.Port) + ")/" + m.DB + "?charset=utf8mb4&parseTime=True&loc=Local"
}
