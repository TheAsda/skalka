package main

import (
	"fmt"
	"github.com/TheAsda/skalka/internal/parser"
	"github.com/TheAsda/skalka/pkg/config"
	"gopkg.in/yaml.v2"
	"io"
)

type ConsoleWriter struct {
	scope string
}

func (c ConsoleWriter) Write(p []byte) (n int, err error) {
	fmt.Printf("[%s]:%s", c.scope, p)
	return len(p), nil
}

type ConsoleWriterFactory struct {
}

func (c ConsoleWriterFactory) GetLogger(name string) io.Writer {
	return ConsoleWriter{scope: name}
}

const (
	yml = `
version: 0
jobs:
  go:
    steps:
    - name: Check go version
      run: go version
  nodejs:
    steps:
    - name: Check node version
      run: node --version
    - name: Check npm version
      run: npm --version
`
)

type MockFileReader struct {
}

func (m MockFileReader) Read(path string) ([]byte, error) {
	return []byte(yml), nil
}

func (m MockFileReader) ReadString(path string) (string, error) {
	panic("implement me")
}

func main() {
	conf := config.FlowConfig{
		Version:      0,
		Env:          []string{},
		Requirements: []config.Requirement{},
		Jobs: map[string]config.Job{
			"nodejs": {
				FlushOnError: false,
				Steps: []config.Step{
					{
						Name: "Check node version",
						Run:  "node --version",
					},
					{
						Name: "Check npm version",
						Run:  "npm --version",
					},
				}},
			"go": {
				FlushOnError: false,
				Steps: []config.Step{
					{
						Name: "Check go version",
						Run:  "go version",
					},
				},
			},
		},
	}
	str, err := yaml.Marshal(conf)
	if err != nil {
		fmt.Printf("Err: %s", err.Error())
		return
	}

	fmt.Printf("%s", str)

	p := parser.NewParser(MockFileReader{})

	flowConfig, err := p.ParseFlowConfig("")
	if err != nil {
		fmt.Printf("Err: %s", err.Error())
		return
	}
	fmt.Printf("%+v", flowConfig)
}
