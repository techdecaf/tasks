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
variables:
  # you can use static values
  MY_VARIABLE: default_value
  # variables are executed in order
  # you can execute shell commands
  # you can use go template syntax
  MY_CMD_VAR: exec(echo {{.MY_VARIABLE}})

tasks:
  task1:
    description: help text for `tasks list`
    commands:
      - echo hello $MY_CMD_VAR
  task2:
    description: runs task1 first!
    pre: [task1] # run task 1 before running this task
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
