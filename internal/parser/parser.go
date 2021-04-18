package parser

import (
	"github.com/TheAsda/skalka/internal/fs"
	"github.com/TheAsda/skalka/pkg/config"
	"gopkg.in/yaml.v2"
)

type Parser struct {
	reader fs.PathReader
}

func NewParser(reader fs.PathReader) *Parser {
	return &Parser{reader: reader}
}

func (p *Parser) ParseFlowConfig(path string) (*config.FlowConfig, error) {
	data, err := p.reader.Read(path)
	if err != nil {
		return nil, err
	}
	var flowConfig config.FlowConfig
	err = yaml.Unmarshal(data, &flowConfig)
	if err != nil {
		return nil, err
	}
	return &flowConfig, nil
}
