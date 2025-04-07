package zlog

type LogConfig struct {
	LogPath    string `json:"logPath" yaml:"logPath"`
	LogLevel   string `json:"logLevel" yaml:"logLevel"`
	MaxSize    int    `json:"maxSize" yaml:"maxSize"`
	MaxBackups int    `json:"maxBackups" yaml:"maxBackups"`
	MaxAge     int    `json:"maxAge" yaml:"maxAge"`
	Compress   bool   `json:"compress" yaml:"compress"`
}
