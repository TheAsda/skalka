package config

// Description of a flow that can be executed
type FlowConfig struct {
	// Config version that is meant to exist for backwards compatibility
	Version float32
	// Global envs that will be provided for every step
	Env Env
	// Required tools for the flow
	Requirements []Requirement
	// Collection of workflows that will be executed concurrently
	Jobs Jobs
}

// Required tool that will be checked for before running jobs
type Requirement struct {
	// Command that must be executed and end with code 0
	Cmd string
}

// The name of a job
type JobName = string

// Collection of names and workflows
type Jobs = map[JobName]Job

// Configuration actions and configurations for current workflow
type Job struct {
	// Defines whether changes made to the file system must be reverted on error
	FlushOnError bool
	// Envs that will be provided to the steps of current job
	Env Env
	// Commands that should be executed in a row
	Steps []Step
}

// One action that should be executed after previous Step completed
type Step struct {
	// If defined step is treated as question contained and Run must not be provided
	StepWithOptions `yaml:",inline"`
	// If defined step runs plugin and must not have Run
	StepPlugin `yaml:",inline"`
	// Name of the action
	Name string
	// Environment variables for the step
	Env Env
	// Command that should be executed
	Run string
}

// Step contains options that will be executed depends on variable value
type StepWithOptions struct {
	// Variable name where question answer is stored
	Variable string
	// Question that will be asked before options pick
	Question string
	// Collection of options that user can choose from
	Options Options
}

// Collection of option value as key and description as value
type Options = map[string]StepOption

// Description of option
type StepOption struct {
	// Displayed name of the option
	Name string
	// If this option is selected by default, works only if variable is not defined on previous steps
	IsDefault bool
	// Command that should be executed
	Run string
	// Defines that option should skip execution
	Skip bool
}

// Step that uses plugin for processing
type StepPlugin struct {
	// The name of a plugin that is chosen for execution
	Uses string
}
