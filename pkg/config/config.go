package config

type Config struct {
	Version      *int
	Env          Env
	Requirements []Requirement
	Jobs         Jobs
}

type Env = []string

type Requirement struct {
	Cmd  *string
	Path *string
}

type Jobs = map[string][]Job

type Job struct {
	FlushOnError *bool
	Steps        []Step
}

type Step struct {
	StepWithOptions
	StepPlugin
	Name string
	Env  Env
	Run  string
}

type StepWithOptions struct {
	Variables string
	Question  string
	Options   map[string][]StepOption
}

type StepPlugin struct {
	Uses string
}

type StepOption struct {
	Name      string
	IsDefault bool
	Run       string
	Skip      bool
}
