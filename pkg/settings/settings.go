package settings

type Settings struct {
	LogLevel LogLevelType `yaml:"log_level"`
	Executor struct {
		Command   string
		Arguments string
	}
}
