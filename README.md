# tasks

<p align="center">
  <img
    alt="tasks"
    src="https://images.techdecaf.com/fit-in/100x/tiny/tasks-logo.png"
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

```powershell
Set-ExecutionPolicy Bypass -Scope Process -Force; iex ((New-Object System.Net.WebClient).DownloadString('https://raw.github.com/techdecaf/tasks/master/install.ps1'))
```

Download Links

- [windows](http://github.techdecaf.io/tasks/latest/windows/tasks.exe)
- [mac](http://github.techdecaf.io/tasks/latest/darwin/tasks)
- [linux](http://github.techdecaf.io/tasks/latest/linux/tasks)

To install tasks, use the provided script, simlink it or place it in any directory that is part of your path.
i.e. `/usr/local/bin` or `c:\windows`


## Quick Start

```text
runs a list of tasks as defined in your taskfile.yaml

Usage:
  tasks [flags]
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
  -h, --help               help for tasks
  -s, --silent             suppress log messages <overrides log option in taskfile.yaml>
  -f, --task-file string   use a specific taskfile.
  -v, --version            Prints application version

Use "tasks [command] --help" for more information about a command.
```

1. [install tasks](#download-and-install)
2. create a taskfile.yaml in the root of your project
3. `tasks run task2` with the following yaml would result in task 1, then task2 running.

### taskfile.yaml

```text
options:
  log: true # debug, info, error, silent

# all task variables are environmental variables, if they do not already exist in the current
# environment, then you can set them here using the golang template syntax
variables:
  # Get the current working directory and extract the base path
  CI_PROJECT_NAME: "{{ReadFile `.cgen.yaml` | YQ `answers.Name`}}" # === tasks
  # TRY to execute git describe --tags, defaults to an empty string
  CI_COMMIT_TAG: "{{TRY `git describe --tags --always --abbrev=0`}}"
  # TRY to get the current branch name using git rev-parse
  CI_COMMIT_REF_NAME: "{{TRY `git rev-parse --abbrev-ref HEAD`}}"
  # TRY to get the current commit sha
  CI_COMMIT_SHA: "{{TRY `git rev-parse HEAD`}}"
  # Sets a static value
  S3_BUCKET: github.techdecaf.io
  # use the dot variable syntax to template the following url from variables
  DOWNLOAD_URI: http://{{.S3_BUCKET}}/{{.CI_PROJECT_NAME}}
  # generate a docker image based on the current version
  DOCKER_IMAGE_TAG: techdecaf/tasks:{{.CI_COMMIT_TAG | replace `v` ``}}
  # # find files using glob pattern in current directory matching which returns a list
  INSTALL_SCRIPTS: "{{GlobMatch PWD `install.*` | join `|`}}"

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
      GOARCH: amd64
    commands:
      - GOOS=darwin go {{.flags}} -o build/darwin/{{.CI_PROJECT_NAME}} -v
      - GOOS=linux go {{.flags}} -o build/linux/{{.CI_PROJECT_NAME}} -v
      - GOOS=windows go {{.flags}} -o build/windows/{{.CI_PROJECT_NAME}}.exe -v
      # - docker build . -t {{.DOCKER_IMAGE_TAG}}

  clean:
    description: removes all files listed in .gitignore
    commands: ["rm -rf build temp"]

  install:
    description: installs locally to /usr/local/bin
    commands:
      - "chmod +x build/{{OS}}/{{.CI_PROJECT_NAME}}"
      - "cp build/{{OS}}/{{.CI_PROJECT_NAME}} /usr/local/bin"

  publish:
    description: publish artifacts to S3
    commands:
      - "aws s3 sync --acl bucket-owner-full-control build s3://{{.S3_BUCKET}}/{{.CI_PROJECT_NAME}}/{{.CI_COMMIT_TAG}}"
      - "aws s3 sync --acl bucket-owner-full-control build s3://{{.S3_BUCKET}}/{{.CI_PROJECT_NAME}}/latest"
      - "docker push {{.DOCKER_IMAGE_TAG}}"

  login:
    description: checkout temporary aws access keys
    commands:
      - curl -s "$DECAF_URL/keys/aws/set/env/linux/website-update?jwt=$DECAF_TOKEN"

  docs:
    description: auto generate documentation
    commands:
      # expand the file, and pipe to write, if no errors default to the string "success"
      - "echo {{ExpandFile `docs/README.md` | WriteFile `README.md` | default `docs updated`}}"

  test:
    description: run tests
    pre: [clean]
    commands:
      - go run . help
      - go run . list

  upgrade:
    description: upgrade project from cgen template
    commands: ["cgen upgrade"]

  # coverage:
  #   description: run test coverage
  #   commands:
  #     - "go test ./app -coverprofile coverage.out && go tool cover -func=coverage.out"

  release:
    description: bump version and release for deployment
    commands:
      - cgen bump --level {{.type}} --push

  pre-release:
    description: bump the prerelease version and deploy skipping production
    commands: [tasks run release --variable type=pre-release]

  release-patch:
    description: bump minor version and release for deployment
    commands: [tasks run release --variable type=patch]

  release-feature:
    description: bump minor version and release for deployment
    commands: [tasks run release --variable type=minor]

  release-breaking-change:
    description: bump major version and release for deployment
    commands: [tasks run release --variable type=major]

  oops:
    description: undo last commit
    commands: [git reset HEAD~1]

```

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

