package config

// Logger is the configuration for the logger.
type Logger struct {
	Level        string `yaml:"level"`
	Prefix       int    `yaml:"prefix"`
	Director     string `yaml:"director"`
	ShowLine     bool   `yaml:"show_line"`      // 是否显示行号
	LogInConsole bool   `yaml:"log_in_console"` // 是显示打印的路径
}
