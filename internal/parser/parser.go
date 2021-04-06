package parser

import (
	"fmt"
	"github.com/TheAsda/skalka/pkg/config"
	"regexp"
	"strings"
)

type Parser struct {
	isLinted bool
	config   config.Config
}

func NewParser(config config.Config) *Parser {
	return &Parser{config: config}
}

func (p Parser) Lint() error {
	err := p.lintVersion()

	if err != nil {
		return err
	}

	err = p.lintEnv(p.config.Env)

	if err != nil {
		return err
	}

	err = p.lintRequirements()

	p.isLinted = true
	return nil
}

func (p Parser) lintVersion() error {
	if p.config.Version != nil {
		return nil
	}
	return NewError("Version is not specified")
}

func (p Parser) lintEnv(env config.Env) error {
	envRegExp := regexp.MustCompile("\\w+:\\w+")
	for _, item := range env {
		item = strings.ReplaceAll(item, " ", "")
		if !envRegExp.MatchString(item) {
			return NewError(fmt.Sprintf("Wrong env format `%s`", item))
		}
	}
	return nil
}

func (p Parser) lintRequirements() error {
	for _, item := range p.config.Requirements {
		if item.Cmd == nil {
			return NewError(fmt.Sprintf("Cmd must be specified `%+v`", item))
		}
	}
	return nil
}

func (p Parser) lintJobs() error {
	nameRegExp := regexp.MustCompile("\\w+")

	for key, value := range p.config.Jobs {
		if !nameRegExp.MatchString(key) {
			return NewError(fmt.Sprintf(""))
		}
		for _, item := range value {
			err := p.lintJob(item)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (p Parser) lintJob(job config.Job) error {
	return nil
}

func (p Parser) GetConfig() (*config.Config, error) {
	if !p.isLinted {
		err := p.Lint()
		if err != nil {
			return nil, err
		}
	}
	return &p.config, nil
}
