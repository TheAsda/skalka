package config

type Config struct {
	version      int
	env          Env
	requirements []Requirement
	jobs         Jobs
}

type Env = []string

type Requirement struct {
	cmd  string
	path string
}

type Jobs = map[string]Job

type Job struct {
	flushOnError bool
	steps        []Step
}

type Step struct {
}
