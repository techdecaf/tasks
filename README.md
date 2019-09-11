<p align="center">
  <img alt="cgen" src="https://images.techdecaf.com/fit-in/100x/tiny/tasks-logo.png" width="100" />
</p>

# Tasks

Tasks is a task runner written in GO. Designed to be a simple task runner supporting both local development and ci/cd pipelines.

- [Tasks](#tasks)
  - [Installing Tasks](#installing-tasks)
  - [Using Tasks](#using-tasks)
    - [Quick Start](#quick-start)
    - [taskfile.yaml](#taskfileyaml)
    - [Go Template Options](#go-template-options)
  - [CLI Options](#cli-options)
  - [Variables](#variables)
- [Credits](#credits)
  - [Application Design](#application-design)
  - [Logo](#logo)
  - [Sponsor](#sponsor)

## Installing Tasks

```bash
sh -c "$(curl -fsSL https://raw.github.com/techdecaf/tasks/master/install.sh)"
```

## Using Tasks

### Quick Start

1. [install tasks](#installing-tasks)
2. create a taskfile.yaml in the root of your project
3. `tasks run task2` with the following yaml would result in task 1, then task2 running.

```yaml
# options can be specified
options:
  log: true # boolean, write logs in addition to stdout / stderr

# variables are resolved in the following order:
# variables declared as part of tasks
# environmental variables that exist before the task is run
# global variables listed here.
# variables are run in order, and also mapped to the environment, so you can feel free to use
# other variables in subsequent commands.

# to see how your variable are being interpreted, you can run tasks list which will evaluate your global
# variables and list tasks with their descriptions
variables:
  # an example of a static variable
  CI_PROJECT_NAME: tasks
  # the EXEC command can be run anywhere string interpretation is done
  CI_COMMIT_TAG: "{{ EXEC(git describe --tags --always --dirty --abbrev=0) }}"
  CI_COMMIT_REF_NAME: "{{EXEC `git rev-parse --abbrev-ref HEAD`}}"
  CI_COMMIT_SHA: "{{EXEC `git rev-parse HEAD`}}"
  S3_BUCKET: github.techdecaf.io

# tasks can be run by performing `tasks run <list of tasks>`
tasks:
  # optionally you can specify a default task that will run if no other tasks are specified.
  default:
    description: is the task that runs when no tasks have been specified. `tasks run` == `tasks run default`
    commands: [tasks list]

  dependencies:
    description: install all required dependencies
    commands: [go get, go install]

  build:
    description: compile window, linux, osx x64
    pre: [clean, dependencies]
    # local variables can be used and will persist even if the task was run in the "pre" step.
    # note, that local variables will always trump environmental variables and global variables
    variables:
      flags: build -ldflags "-X main.VERSION={{.CI_COMMIT_TAG}}"
    commands:
      - GOOS=darwin go {{.flags}} -o dist/darwin/{{.CI_PROJECT_NAME}} -v
      - GOOS=linux go {{.flags}} -o dist/linux/{{.CI_PROJECT_NAME}} -v
      - GOOS=windows go {{.flags}} -o dist/windows/{{.CI_PROJECT_NAME}}.exe -v

  clean:
    description: removes all files listed in .gitignore
    commands: ["rm -rf dist"]

  install:
    description: installs locally to /usr/local/bin
    commands:
      - "chmod +x dist/{{OS}}/{{.CI_PROJECT_NAME}}"
      - "mv dist/{{OS}}/{{.CI_PROJECT_NAME}} /usr/local/bin"

  publish:
    description: moves compiled files to /usr/local/bin/
    commands:
      - "aws s3 sync dist s3://{{.S3_BUCKET}}/{{.CI_PROJECT_NAME}}/{{.CI_COMMIT_TAG}}"
      - "aws s3 sync dist s3://{{.S3_BUCKET}}/{{.CI_PROJECT_NAME}}/latest"
      - "aws s3 cp install.sh s3://{{.S3_BUCKET}}/{{.CI_PROJECT_NAME}}/install.sh"

  fails:
    commands: [does_not_exist]
```

### taskfile.yaml

### Go Template Options

- [Go Template Documentation](https://golang.org/pkg/text/template/)
- Additionally we add functions from the [Sprig](http://masterminds.github.io/sprig/) library for convenience

## CLI Options

```bash
Usage:
  tasks [flags]
  tasks [command]

Available Commands:
  completion  generates a zsh plugin with completion scripts
  exec        execute a command using all resolved variables from your taskfile.yaml
  export      Resolves all global variables and prints an export script
  help        Help about any command
  init        initialize a task file in the current directory.
  list        list available commands and their descriptions from your taskfile.yaml
  run         runs a list of tasks as defined in your taskfile.yaml

Flags:
      --config string      config file (default is $HOME/.tasks.yaml)
  -h, --help               help for tasks
  -l, --log string         tasks run -l <true|false> task1 task2 <overrides log option in taskfile.y
aml>
  -f, --task-file string   use a specific taskfile.
  -v, --version            Prints application version

Use "tasks [command] --help" for more information about a command.
```

## Variables

variables are run in order, and also mapped to the environment, so you can feel free to use your defined variables in subsequent commands.

variables are resolved in the following order:

1. local variables declared as part of a task
2. environmental variables that exist before the task is run
3. global variables listed in the taskfile.yaml

# Credits

## Application Design

Application design taken from [go-task](https://github.com/go-task/task). Lots of key design elements come from this project, the only reason to roll our own was a fundamental breaking change on how to handle variables.

## Logo

The logo for this project provided by [logomakr](https://logomakr.com)

## Sponsor

[![TechDecaf](https://images.techdecaf.com/fit-in/150x/techdecaf/logo_full.png)](https://techdecaf.com)

_Get back to doing what you do best, let us handle the rest._
