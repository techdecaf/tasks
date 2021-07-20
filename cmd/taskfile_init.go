package cmd

var SAMPLE_TASKFILE = `
# documentation for tasks can be found here https://github.com/techdecaf/tasks
options:
  log: true

variables:
  CI_PROJECT_NAME: "{{PWD | base}}"
  CI_COMMIT_TAG: "{{TRY ~~git describe --tags --always --dirty --abbrev=0~~}}"
  CI_COMMIT_REF_NAME: "{{TRY ~~git rev-parse --abbrev-ref HEAD~~}}"
  CI_COMMIT_SHA: "{{TRY ~~git rev-parse HEAD~~}}"

tasks:
  default:
    description: runs when no tasks have been specified.
    commands: [tasks list]

  clean:
    description: remove temporary files or directories
    commands: []

  dependencies:
    description: install all required dependencies
    commands: []

  test:
    description: run tests
    commands: []

  coverage:
    description: run test including coverage
    commands: []

  build:
    description: build current project
    commands: []

  deploy:
    description: deploy the current project
    commands: []

  upgrade:
    description: upgrade the current project
    commands: []
`
