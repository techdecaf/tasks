package cmd

import (
	"fmt"

	"github.com/techdecaf/templates"
)

// SetEnvFrom map[string]string
func SetEnvFrom(variables map[string]string) {
	for key, val := range variables {
		err := tasks.TemplateVars.Set(templates.Variable{
			Key:         key,
			Value:       val,
			OverrideEnv: true,
		})
		if err != nil {
			log.Fatal(fmt.Sprintf("could not set variable %s = %s", key, val), err)
		}
	}
}
