# fetakv

[![Build Status](https://travis-ci.com/frankgreco/fetakv.svg?token=MkdavBWRqQGB4gWqK2cR&branch=master)](https://travis-ci.com/frankgreco/fetakv)

> a command line REPL that drives a thread safe, in-memory kv storage system

## Quick Start
```sh
$ mkdir -p $GOPATH/src/github.com/frankgreco
$ cd $GOPATH/src/github.com/frankgreco
$ git clone git@github.com:frankgreco/fetakv.git
$ cd fetakv
$ make binary
$ fetakv
>
```

## Advanced Usage

It may be useful to process transactions in bulk. To do this, simply set the `--stdin` flag to the location of your transactions.

Below is an example taken from this project's integration tests. Note that both the `stdout` and `stderr` file descriptors have been redirected to simulate a real world use case.

```sh
$ fetakv --stdin examples/stdin.txt 1>>examples/stdout.txt 2>>examples/stderr.txt
$ cat examples/stdout.txt
bar
one
...
$ cat examples/stderr.txt
There are no current transactions to abort.
Key not found: foo
...
```

## Development Guide
`featkv` uses [`dep`](https://github.com/golang/dep) for dependencies.

```sh
$ mkdir -p $GOPATH/src/github.com/frankgreco
$ cd $GOPATH/src/github.com/frankgreco
$ git clone git@github.com:frankgreco/fetakv.git
$ cd fetakv
$ make # format, lint, and test everything
```

## Usage
```sh
$ fetakv
> HELP
fetakv a command line REPL that drives an in-memory kv storage system.

Available Commands:
  HELP                Display usage information.
  READ    <key>       Read the value associated with the given key.
  WRITE   <key> <val> Stores val in key.
  DELETE  <key>       Removes key from the store.
  START               Start a transation (nested transactions supported).
  COMMIT              Commit a transaction.
  ABORT               Aborts a transation.
  QUIT                Exit fetakv.
>
```