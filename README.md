![orbi](https://github.com/bronzdoc/orbi/blob/master/logo.png)
[![Build Status](https://travis-ci.org/bronzdoc/orbi.svg?branch=master)](https://travis-ci.org/bronzdoc/orbi)

> Project structure generator.

Generate project structures using yaml and language agnostic templates.

# Install
> NOTE: orbi just works with unix-like operating systems, windows is not supported for now.

### Binaries

- **linux** [386](https://github.com/bronzdoc/orbi/releases/download/v0.0.0/orbi-linux-386) / [amd64](https://github.com/bronzdoc/orbi/releases/download/v0.0.0/orbi-linux-amd64) / [arm](https://github.com/bronzdoc/orbi/releases/download/v0.0.0/orbi-linux-arm) / [arm64](https://github.com/bronzdoc/orbi/releases/download/v0.0.0/orbi-linux-arm64)
- **darwin** [386](https://github.com/bronzdoc/orbi/releases/download/v0.0.0/orbi-darwin-386) / [amd64](https://github.com/bronzdoc/orbi/releases/download/v0.0.0/orbi-darwin-amd64)
- **freebsd** [386](https://github.com/bronzdoc/orbi/releases/download/v0.0.0/orbi-freebsd-386) / [amd64](https://github.com/bronzdoc/orbi/releases/download/v0.0.0/orbi-freebsd-amd64)

### Via Go

```shell
$ go get github/bronzdoc/orbi
```

# Usage

Orbi defines project structures using a `definition.yml`.

e.g:

```yaml
---
context: .
resources:
  - dir:
     name: dir_1
     files:
      - file_a

  - dir:
     name: dir_2
     files:
      - file_b
     dir:
      name: dir_3
      files:
        - file_c
        - file_d

  - files:
     - file_e
     - file_f
```

A `context` is where your `resources` structure will be created.

The way orbi organize definitions is with something called a `plan`, you can create a new plan by doing:

```shell
$ orbi plan new my_plan
```

This command will generate the following:

```shell
$ tree ~/.orbi/plans/my_plan

my_plan
├── definition.yml
└── templates
```

You can notice a `templates` directory, your templates go there...

In order tu template a file all you need to do is create a file named the same as a file resource.

e.g:

```yaml
context: .
resources:
  files:
    - file_a
```

```shell
├── definition.yml
└── templates
    └── file_a
```

Ok, so we have a plan with our definition and templates, so how we create all that stuff we defined?

All you need to do is tell orbi to execute a `plan`:

```shell
$ orbi exec my_plan
```

this command will generate the file structure defined in your plan definition.

If your plan templates happen to have variables, you can pass arguments to those variables with the `--vars` flag.

e.g:

in `.orbi/plans/tiesto/templates/file_a`

```
{{.name}} is awesome
```

```shell
$ orbi exec my_plan --vars="name=Tarantino"
```

that command will generate the file named `file_a` with content `Tarantino is awesome`

## Contributing

Contributions are greatly appreciated, and encouraged
