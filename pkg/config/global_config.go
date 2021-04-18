package config

import (
	"github.com/TheAsda/skalka/pkg/settings"
	"os"
)

type GlobalConfig struct {
	Settings settings.Settings
	Env      Env
	Workdir  string
}

func NewGlobalConfig(settings settings.Settings, env Env, workdir string) *GlobalConfig {
	globalEnv := os.Environ()
	return &GlobalConfig{Settings: settings, Env: append(globalEnv, env...), Workdir: workdir}
}
