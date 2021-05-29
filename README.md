# About

[![Circle CI](https://circleci.com/gh/g1eng/giosh.svg?style=shield)](https://app.circleci.com/pipelines/github/g1eng/giosh)
[![codecov](https://codecov.io/gh/g1eng/giosh/branch/master/graph/badge.svg?token=2F9PE0MC9B)](https://codecov.io/gh/g1eng/giosh)

giosh (generic IO shell) is a multipurpose command line shell written in Go.

giosh is not:

* a POSIX compliant shell
* a shell language with multi-line lexical scope
* a virtual machine with a new (environmental) variable declaration mechanism

# Supported features

* Command invokation with PATH resolving
* pipe

# Usage

1. Simply invoke it. You can use giosh with interactive mode. 

```shell
$ giosh
@G[1]> ls -l | grep m | sed -r -e s/.+[[:blank:]]//g
README.md
cmd
main.go
```

2. Read giosh script 

```shell
$ giosh some_script
```

3. Read an oneliner from the standard input
   (Standard input for giosh must be an oneliner)

```shell
$ awesome_script | awesome_command | giosh
```

# Limitation

At now, giosh is released under limitations (really so limited in the function).

Following mechanisms are not supported:


* comment line
* escape character \\ 
* arrow keys and command history 
* redirection with file descriptor
* hear strings or hear document
* variable declaration
* variable reference (env value is not directly available)
* variable expansion and command replacement
* background jobs (&)
* builtin commands

# Authors

Suzume Nomura [@g1eng](https://github.com/g1eng)

Okadarien Saru -
