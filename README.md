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

## Installing Tasks

```bash
curl -sL http://github.techdecaf.io/tasks/install.sh | sh
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

`tasks run [tasks]` run any number of tasks in order
`tasks list` resolve all task variables and list all tasks with their descriptions
`tasks --log true|false <cmd>` enables / disables additional log output, all regular command output is still sent to stdout
ie: `[task_name][running] <command string>`

## Variables

## Credits

### Application Design

Application design taken from [go-task](https://github.com/go-task/task). Lots of key design elements come from this project, the only reason to roll our own was a fundamental breaking change on how to handle variables.

The logo for this project provided by https://logomakr.com
