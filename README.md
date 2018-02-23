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

## Testing
`featkv` uses [`dep`](https://github.com/golang/dep) for dependencies.
```sh
$ cd $GOPATH/src/frankgreco/fetakv
$ make install_ci       # install all dependencies for testing
$ make                  # format, lint, and test everything
$ make fmt              # format all files
$ make lint             # lint all files
$ make unit_test        # only unit tests
$ make integration_test # only integration tests
$ make test             # all tests
```
