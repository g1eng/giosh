# About

giosh (generic IO shell) is a multipurpose command line shell written in Go.
giosh is not:

* a POSIX compliant shell
* a shell language with multi-line lexical scope
* a virtual machine with a new (environmental) variable declaration mechanism

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

At now, giosh is released under following limitations:

* comment line is not supported
* escape character \\ is not supported
* arrow keys and command history are not supported

# Authors

Suzume Nomura @g1eng

Okadarien Saru -