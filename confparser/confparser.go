package confparser

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"strings"

	"gopkg.in/yaml.v2"
)

type Parser struct {
	fd     *os.File
	format string
}

func New(file string) *Parser {
	var format string = ""
	if strings.HasSuffix(file, ".yaml") || strings.HasSuffix(file, ".yml") {
		format = "yaml"
	} else if strings.HasSuffix(file, "json") {
		format = "json"
	} else {
		return nil
	}
	fd, err := os.OpenFile(file, os.O_RDWR, 0644)
	if err != nil {
		return nil
	}
	return &Parser{
		fd:     fd,
		format: format,
	}
}

func (p *Parser) Load(v interface{}) (err error) {
	p.fd.Seek(0, 0)

	content, err := ioutil.ReadAll(p.fd)
	if err != nil {
		return err
	}
	switch p.format {
	case "json":
		return json.Unmarshal(content, v)
	case "yaml":
		return yaml.Unmarshal(content, v)
	default:
		return errors.New("file format not suppport")
	}
}

func (p *Parser) Save(v interface{}) (err error) {
	p.fd.Seek(0, 0)

	var content []byte
	switch p.format {
	case "json":
		content, err = json.MarshalIndent(v, "", "  ")
		if err != nil {
			return err
		}
	case "yaml":
		content, err = yaml.Marshal(v)
		if err != nil {
			return err
		}
	default:
		return errors.New("file format not suppport")
	}

	p.fd.Truncate(0)
	_, err = p.fd.Write(content)

	return err
}
