package taskfile

import (
	"github.com/techdecaf/templates"
	yaml "gopkg.in/yaml.v2"
)

func map2var(item yaml.MapItem, override bool) templates.Variable {
	return templates.Variable{
		Key:         item.Key.(string),
		Value:       item.Value.(string),
		OverrideEnv: override,
	}
}
