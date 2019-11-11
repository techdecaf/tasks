1. [install tasks](#download-and-install)
2. create a taskfile.yaml in the root of your project
3. `tasks run task2` with the following yaml would result in task 1, then task2 running.

```text
{{ ReadFile `taskfile.yaml` }}
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
