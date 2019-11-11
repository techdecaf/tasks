# tasks

<p align="center">
  <img
    alt="tasks"
    src="https://images.techdecaf.com/fit-in/100x/techdecaf/tasks_logo.png"
    width="100"
  />
</p>


- [tasks](#ciprojectname)
  - [Download and Install](#download-and-install)
  - [Quick Start](#quick-start)
  - [Contribution Guide](#contribution-guide)
  - [Credits](#credits)

## Download and Install

```bash
sh -c "$(curl -fsSL https://raw.github.com/techdecaf/tasks/master/install.sh)"
```

Download Links

- [windows](http://github.techdecaf.io/tasks/latest/windows/tasks.exe)
- [mac](http://github.techdecaf.io/tasks/latest/latest/darwin/tasks)
- [linux](http://github.techdecaf.io/tasks/latest/latest/linux/tasks)

To install tasks, use the provided script, simlink it or place it in any directory that is part of your path.
i.e. `/usr/local/bin` or `c:\windows`


## Quick Start

```text
Tasks is a task runner written in GO. Designed to be a simple task runner supporting both local development and ci/cd pipelines.

Usage:
  tasks [command]

Available Commands:
  completion  Generates zsh completion scripts
  exec        execute a command using all resolved variables from your taskfile.yaml
  export      Resolves all global variables and prints an export script
  help        Help about any command
  init        initialize a task file in the current directory.
  list        list available commands and their descriptions from your taskfile.yaml
  run         runs a list of tasks as defined in your taskfile.yaml

Flags:
  -h, --help     help for tasks
  -t, --toggle   Help message for toggle

Use "tasks [command] --help" for more information about a command.
```

1. [install tasks](#download-and-install)
2. create a taskfile.yaml in the root of your project
3. `tasks run task2` with the following yaml would result in task 1, then task2 running.

```text
options:
  log: true # debug, info, error, silent

variables:
  CI_PROJECT_NAME: "{{EXEC `echo ${PWD##*/}`}}"
  CI_COMMIT_TAG: "{{TRY `git describe --tags --always --dirty --abbrev=0`}}"
  CI_COMMIT_REF_NAME: "{{TRY `git rev-parse --abbrev-ref HEAD`}}"
  CI_COMMIT_SHA: "{{TRY `git rev-parse HEAD`}}"
  S3_BUCKET: github.techdecaf.io
  DOWNLOAD_URI: http://{{.S3_BUCKET}}/{{.CI_PROJECT_NAME}}/latest

tasks:
  default:
    description: is the task that runs when no tasks have been specified. `tasks run` == `tasks run default`
    commands: [tasks list]

  dependencies:
    description: install all required dependencies
    commands: [go get, go install]

  build:
    description: compile window, linux, osx x64
    pre: [clean, dependencies]
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
      - "cp dist/{{OS}}/{{.CI_PROJECT_NAME}} /usr/local/bin"

  publish:
    description: moves compiled files to /usr/local/bin/
    commands:
      - "aws s3 sync dist s3://{{.S3_BUCKET}}/{{.CI_PROJECT_NAME}}/{{.CI_COMMIT_TAG}}"
      - "aws s3 sync dist s3://{{.S3_BUCKET}}/{{.CI_PROJECT_NAME}}/latest"

  login:
    description: checkout temporary aws access keys
    commands:
      - curl -s "$DECAF_URL/keys/aws/set/env/linux/website-update?jwt=$DECAF_TOKEN"

  docs:
    description: auto generate documentation
    commands:
      # expand the file, and pipe to write, if no errors default to the string "success"
      - "echo {{ExpandFile `docs/README.md` | WriteFile `README.md` | default `docs updated`}}"

  upgrade:
    description: upgrade project from cgen template
    commands: ["cgen upgrade"]

```

### taskfile.yaml

### Go Template Options

- [Go Template Documentation](https://golang.org/pkg/text/template/)
- Additionally we add functions from the [Sprig](http://masterminds.github.io/sprig/) library for convenience

## CLI Options

```bash
<no value>
```

## Variables

variables are run in order, and also mapped to the environment, so you can feel free to use your defined variables in subsequent commands.

variables are resolved in the following order:

1. local variables declared as part of a task
2. environmental variables that exist before the task is run
3. global variables listed in the taskfile.yaml


## Contribution Guide

## Credits

### Application Design

Application design taken from [go-task](https://github.com/go-task/task). Lots of key design elements come from this project, the only reason to roll our own was a fundamental breaking change on how to handle variables.

### Logo

The logo for this project provided by [logomakr](https://logomakr.com)

### Sponsor

[![TechDecaf](https://images.techdecaf.com/fit-in/150x/techdecaf/logo_full.png)](https://techdecaf.com)

_Get back to doing what you do best, let us handle the rest._

