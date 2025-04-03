package zlog

type LogConfig struct {
	LogPath    string `json:"logPath"`
	LogLevel   string `json:"logLevel"`
	MaxSize    int    `json:"maxSize"`
	MaxBackups int    `json:"maxBackups"`
	MaxAge     int    `json:"maxAge"`
	Compress   bool   `json:"compress"`
}
