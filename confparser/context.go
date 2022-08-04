package confparser

import (
	"encoding/json"
	"errors"
)

type InitFunc func(json.RawMessage) error
type SaveFunc func() json.RawMessage

type Option struct {
	Name string
	Init InitFunc
	Save SaveFunc
}

// Context config context
type Context struct {
	parser  *Parser
	Conf    map[string]json.RawMessage
	options map[string]Option
}

func NewContext() *Context {
	ctx := new(Context)
	ctx.Conf = make(map[string]json.RawMessage)
	ctx.options = make(map[string]Option)
	return ctx
}

func (c *Context) Register(opt *Option) error {
	_, found := c.options[opt.Name]
	if found {
		return errors.New("config " + opt.Name + " has been registered")
	}
	if opt.Init == nil {
		return errors.New("init func should not be empty")
	}

	if opt.Save == nil {
		return errors.New("save func should not be empty")
	}
	c.options[opt.Name] = *opt
	return nil
}

func (c *Context) Load(configfile string) error {
	if c.parser == nil {
		parser := NewParser(configfile)
		if parser == nil {
			return errors.New("open file failed")
		}
		c.parser = parser
	}

	err := c.parser.Load(&c.Conf)
	if err != nil {
		return err
	}

	for name, option := range c.options {
		value, ok := c.Conf[name]
		if ok {
			if option.Init != nil {
				err := option.Init(value)
				if err != nil {
					return err
				}
			}
		}
	}

	return nil
}

// Save save config
func (c *Context) Save() {
	for name, option := range c.options {
		if option.Save != nil {
			str := option.Save()
			c.Conf[name] = str
		}
	}
	if c.parser != nil {
		c.parser.Save(c.Conf)
	}
}
