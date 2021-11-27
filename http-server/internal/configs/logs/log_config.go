package logs

// LogConfigModel ,
type LogConfigModel struct {
	Path         string `yaml:"path"`
	FileName     string `yaml:"file_name"`
	Level        string `yaml:"level"`
	MaxAge       int    `yaml:"max_age"`
	RotationTime int    `yaml:"rotation_time"`
}
