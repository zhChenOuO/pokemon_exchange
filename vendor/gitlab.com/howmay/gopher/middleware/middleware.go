package middleware

// Config ...
type Config struct {
	Prometheus
	App
}

// Prometheus 用來設定 prometheus
type Prometheus struct {
	SubSystemName string `yaml:"sub_system_name"`
	Bind          string `yaml:"bind"`
	MetricsPath   string `yaml:"metrics_path"`
}

//App app
type App struct {
	VersionHeaderName string `yaml:"version_header_name"`
	SupportedVersion  string `yaml:"supported_version"`
}
