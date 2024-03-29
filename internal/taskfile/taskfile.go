package taskfile

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/techdecaf/golog"
	"github.com/techdecaf/templates"

	yaml "gopkg.in/yaml.v2"
)

var pwd, _ = os.Getwd()
var logger = golog.Log{
	Name: "tasks",
}

// Options type
type Options struct {
	LogLevel bool `yaml:"log"`
}

// Task type
type Task struct {
	Description string        `yaml:"description"`
	Pre         []string      `yaml:"pre"`
	Variables   yaml.MapSlice `yaml:"variables"`
	Commands    []string      `yaml:"commands"`
	Options     Options       `yaml:"options"`
	Dir         string        `yaml:"dir"`
}

// TaskFile type
type TaskFile struct {
	Options   Options         `yaml:"options"`
	Variables yaml.MapSlice   `yaml:"variables"`
	Tasks     map[string]Task `yaml:"tasks"`
	FilePath  string

	// private
	TemplateVars templates.Variables
}

// Init function
func (tasks *TaskFile) Init() (err error) {
	if tasks.FilePath == "" {
		// use current working directory to get taskfile
		tasks.FilePath = filepath.Join(pwd, "taskfile.yaml")
	}

	// ensure taskfile file exists
	if _, err := os.Stat(tasks.FilePath); err != nil {
		return err
	}

	taskfile, err := os.Open(tasks.FilePath)
	if err != nil {
		return err
	}

	defer taskfile.Close()

	byteValue, err := ioutil.ReadAll(taskfile)
	if err != nil {
		return err
	}

	if err := yaml.Unmarshal(byteValue, &tasks); err != nil {
		return err
	}

	for _, item := range tasks.Variables {
		tasks.TemplateVars.List = append(tasks.TemplateVars.List, map2var(item, false))
	}
	if err := tasks.TemplateVars.Init(); err != nil {
		return err
	}

	// if err := tasks.toJSON(); err != nil {
	// 	return err
	// }

	return nil
}

// Execute a command using all variables resolved in the taskfile
func (tasks *TaskFile) Execute(cmd, name, dir string) (out string, err error) {
	command, err := templates.Expand(cmd, tasks.TemplateVars.Functions)
	if err != nil {
		return "", err
	}

	if tasks.Options.LogLevel {
		logger.Info(name, command)
	}

	return templates.Run(templates.CommandOptions{
		Cmd:       command,
		Dir:       dir,
		UseStdOut: true,
	})
}

// Run task
func (tasks *TaskFile) Run(key string) error {
	task := tasks.Tasks[key]

	// run pre tasks
	if len(task.Pre) > 0 {
		for _, i := range task.Pre {
			if err := tasks.Run(i); err != nil {
				return err
			}
		}
	}

	// resolve variables
	if len(task.Variables) > 0 {
		for _, v := range task.Variables {
			tasks.TemplateVars.Set(map2var(v, true))
		}
	}

	// run commands
	for _, cmd := range task.Commands {
		res, err := tasks.Execute(cmd, key, task.Dir)
		if err != nil {
			return err
		}

		if res != "" {
			logger.Info(key, res)
		}
	}
	return nil
}

// List - all task descriptions
func (tasks *TaskFile) List() {
	fmt.Println("variables:")
	for _, v := range tasks.TemplateVars.List {
		fmt.Printf("%s%s: %s\n", spaces(4), v.Key, os.Getenv(v.Key))
	}

	fmt.Println("tasks:")
	for key, task := range tasks.Tasks {
		fmt.Printf("%s%s: %s\n", spaces(4), key, strings.TrimSpace(task.Description))
	}
}

// Export variables for use in other applications.
func (tasks *TaskFile) Export() {
	var pattern string
	switch runtime.GOOS {
	case "windows":
		pattern = "$env:%s='%s'"
	case "darwin", "linux":
		pattern = "export %s='%s'"
	default:
		logger.Fatal("export", fmt.Errorf("unsupported platform %s", runtime.GOOS))
	}
	for _, v := range tasks.TemplateVars.List {
		fmt.Println(fmt.Sprintf(pattern, v.Key, os.Getenv(v.Key)))
	}
}

func (tasks *TaskFile) toJSON() error {
	json, err := json.Marshal(tasks)
	if err != nil {
		return err
	}

	fmt.Println(string(json))

	return nil
}
