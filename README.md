![orbi-project-generator](http://i.imgur.com/mtUfTaV.png)

> Project structure generator


## Usage

Orbi defines project structures using a `definition.yml` i.e:

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

Orbi uses the `GO` template system to template files, in order tu template a file all you need to do
is create a file named the same as a file resource, i.e:

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

If your plan templates happen to have variables, you can pass arguments to those variables with the `--vars` flag, i.e

in `.orbi/plans/tiesto/templates/file_a`

```
{{.name}} is awesome
```

```shell
$ orbi exec my_plan --vars="name=Tarantino"
```

that command will generate the file named `file_a` with content `Tarantino is awesome`

## Install

```shell
curl -L "https://github.com/bronzdoc/orbi/releases/download/0.0.1/orbi-$(uname -s)-$(uname -m)" -o /usr/local/bin/orbi; chmod +x /usr/local/bin/orbi
```

If you have a GO environment you can simply:

```shell
$ go get github/bronzdoc/orbi
```

## Contributing

Contributions are greatly appreciated

$ go get github.com/onsi/ginkgo/ginkgo
$ go get github.com/onsi/gomega

