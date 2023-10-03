# AOC

Advent of Code workspace. Aim is to have an introduction to golang.

# Requirements

golang docker image or golang installed.

# Commands

```bash
$ cd path/to/repo/root
$ docker run -it -v ${PWD}:/aoc -w /aoc golang:1.21  # optional
```

## Initialize, build and run a module

```bash
> cd path/to/module
> go mod init path/to/module                 # Writes the go.mod file to current working dir.
> go mod edit -replace \
    import/statement=../../import/statement  # Fixes relative imports
> go mod tidy                                # Finds the dependency and tidies up the go.mod file.
> go build -o /tmp/module . && /tmp/module   # --> or equivalently `go run .`
```

## Format a file or recursively all files under a directory

```bash
> go fmt path/to/module/file.go  # Format a file. (Calls gofmt under the hood.)
> gofmt -w -l -s /aoc            # write, list, simplify, recursively all files, given a dir.
```
