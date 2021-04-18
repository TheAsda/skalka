package main

import (
	"fmt"
	"github.com/TheAsda/skalka/internal/executor"
	"github.com/TheAsda/skalka/internal/fs"
	"github.com/TheAsda/skalka/internal/parser"
	"github.com/TheAsda/skalka/internal/progress_logger"
	"github.com/TheAsda/skalka/internal/queue_processor"
	"github.com/TheAsda/skalka/internal/settings_compiler"
	"github.com/TheAsda/skalka/pkg/config"
	"os"
	"path"
)

func main() {
	// Initialize logger
	logger := progress_logger.NewDefaultProgressLogger(os.Stdout)
	// Get workdir
	workdir, err := os.Getwd()
	if err != nil {
		logger.Error("Cannot get working directory")
		return
	}
	settingsPath := path.Join(workdir, "config.yml")
	fileReader := fs.NewFileReader()
	// Initialize settings compiler
	sc := settings_compiler.NewSettingsCompiler(settingsPath, fileReader)
	// Read settings
	settings, err := sc.ReadSettings()
	if err != nil {
		logger.Error(fmt.Sprintf("Cannot read settings: %s", err.Error()))
	}
	// Set log level
	logger.SetLogLevel(settings)
	// Initialize parser
	p := parser.NewParser(fileReader)
	flowConfigPath := path.Join(workdir, "flow.yml")
	// Get flow config
	conf, err := p.ParseFlowConfig(flowConfigPath)
	if err != nil {
		logger.Error(fmt.Sprintf("Cannot parse flow config: %s", err.Error()))
		return
	}
	// Create global config
	gc := config.NewGlobalConfig(settings, conf.Env, workdir)
	// Initialize task runner
	taskRunner := queue_processor.NewTaskRunner(settings.Executor.Command, settings.Executor.Arguments, logger.GetStdout(), logger.GetStderr())
	// Initialize executor
	e := executor.NewExecutor(*gc, *logger, taskRunner)
	// Fill queues
	err = e.Fill(conf.Jobs)
	if err != nil {
		logger.Error(fmt.Sprintf("Cannot fill queues: %s", err.Error()))
	}
	// Run queues
	e.Start()
	// Wait for finish
	e.Wait()
}
