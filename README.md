# renum - strongly typed Go enums

## Overview [![GoDoc](https://godoc.org/github.com/gen0cide/renum?status.svg)](https://godoc.org/github.com/gen0cide/renum) [![Sourcegraph](https://sourcegraph.com/github.com/gen0cide/renum/-/badge.svg)](https://sourcegraph.com/github.com/gen0cide/renum?badge)

Go package that provides a rich, descriptive interface for developers to use in order to allow enums to cross package boundries without loosing important details and metadata.

Also a CLI utility to generate idiomatic Go enums with a diverse set of features and options (that allow you to easily satisfy the `renum` interface ^.^)

*NOTE: This library is in it's early stage, so I wouldn't call it production stable yet. But any PRs and comments are welcome!*

## Background

Go's language, while expressive, has shortcomings around propogation of type information with commonly used code. A great example of this is the [error](https://golang.org/pkg/errors/) interface that's built into the language.

While Go lets you define custom error types (`*os.PathError` is one example), generally developers end up simply using the standard `errors.New` to generate a type that satisfies `error`, but is basically a string containing whatever you passed to `New()`.

As Go (rightly) attempts to force you to handle your errors, it often involves passing errors around, with the expectation that the caller likely  wants to make decisions about what to do. This is an incredibly powerful paradigm, and why I fully support **not** including constructs like exception handling into the runtime.

An example of this occurred for myself recently with the [github.com/masterzen/winrm](https://github.com/masterzen/winrm) package. I was using it to make WinRM connections to a Windows host, but I kept getting an error relating to response header timeouts. Is this happening within `winrm` or `net/http` or `net`? The only way to answer that question was to print the error to the console and begin grepping through source trees, looking for string literals that use those words.

While generically this system is cheap, efficient, and allows broad adoption - it begins to age when working with large, complex codebases where error propagation becomes a lot of manual logging of error messages, with human review consuming considerable time. Forgetting how much `return nil, errors.New("this is bad")` you see, typically this is the "idiomatic" way to define error types in Go:

```go
var (
  // ErrUnauthorized is thrown when a request is not authorized to perform a function.
  ErrUnauthorized = errors.New("request unauthorized")

  // ErrInvalidSQLQuery is thrown when the provided SQL query was not a valid SQL expression.
  ErrInvalidSQLQuery = errors.New("invalid sql query")
)
```

This creats code that is easy to read and now is comparable (even type comparable), hear me out. Imagine a situation where this is printed to a log. You'd see a message of "requested unauthorized". What happens though when another package does this:

```go
return nil, errors.New("request unauthorized")
```

### Solution

`renum` aims to solve this by allow users define "constant" (types that don't change after compilation) types that conform to a more machine friendly and descriptive interface. While errors are a great use case, they certainly aren't the **only** paradigm where this benefits. The interface aims to push users not to write these type definitions by hand, but to generate them with codegen. You certainly could write a type that satisfies `renum.Enum` or `renum.Error`, but after seeing how easy it is to generate, I think you'll gladly let the `renum` utility do the heavy lifting :smiley:

Simply define your types in in YAML:

```yaml
# Enum configuration
go:
  name: ErrorCode
  package_name: lib
  package_path: github.com/gen0cide/renum/example/lib
plugins:
  error: true
  text: true
  json: true
  yaml: true
  sql: true
  description: true
values:
  - name: unauthorized
    message: The request was unauthorized.
    comment: Unauthorized is thrown when the request action cannot be taken.
    description: This error is used to signify that the request was made by an *authenticated* requester, but that requester is not authorized to perform the requested action.
  - name: invalid_sql_query
    message: The provided query was not valid SQL.
    comment: InvalidSQLQuery is thrown when a user supplied SQL query is not valid.
    description: This error often means the caller should perform further validation in order to locate situations where they're taking unsanitized input from users and interpolating that value directly into the SQL query.


```

and use the `renum generate` command in order to codegen a much better error paradigm:

```sh
$ renum -c error_code.yaml generate -o .
[✓] parsed configuration
[✓] initialized generator
[✓] generated Go code
[✓] successfully wrote code to generated_error_codes.go
$
```

And if you opened up `generated_error_codes.go`, you'd see something that looks like this:

```go
// ErrorCode is a generated type alias for the ErrorCode enum.
type ErrorCode int

const (
  // ErrorCodeUnknown is an enum value for type ErrorCode.
  // ErrorCodeUnknown is the default value for enum type ErrorCode. It is meant to be a placeholder and default for unknown values.
  // This value is a default placeholder for any unknown type for the lib.ErrorCode enum.
  ErrorCodeUnknown ErrorCode = iota

  // ErrorCodeUnauthorized is an enum value for type ErrorCode.
  // Unauthorized is thrown when the request action cannot be taken.
  // This error is used to signify that the request was made by an *authenticated* requester, but that requester is not authorized to perform the requested action.
  ErrorCodeUnauthorized

  // ErrorCodeInvalidSQLQuery is an enum value for type ErrorCode.
  // InvalidSQLQuery is thrown when a user supplied SQL query is not valid.
  // This error often means the caller should perform further validation in order to locate situations where they're taking unsanitized input from users and interpolating that value directly into the SQL query.
  ErrorCodeInvalidSQLQuery

// ... more code below
```

To demonstrate how this is now a much richer error interface, I've created a small example program that shows how this now looks to the human eye. I've pasted the output of the example program to demonstrate what features you now have:

```sh
$ go run main.go
[+] renum.Coder interface
[✓] Code() = 2

[+] renum.Namespacer interface
[✓] Namespace() = github.com.gen0cide.renum.cmd.renum.example.lib
[✓]      Path() = github.com.gen0cide.renum.cmd.renum.example.lib.error_code_invalid_sql_query

[+] renum.Typer interface
[✓]        Kind() = lib.ErrorCodeInvalidSQLQuery
[✓]      Source() = github.com/gen0cide/renum/cmd/renum/example/lib.ErrorCodeInvalidSQLQuery
[✓] PackageName() = lib
[✓]  ImportPath() = github.com/gen0cide/renum/cmd/renum/example/lib

[+] renum.Descriptioner interface
[✓] Description() = This error often means the caller should perform further validation in order to locate situations where they're taking unsanitized input from users and interpolating that value directly into the SQL query.

[+] fmt.Stringer interface
[✓] String() = invalid_sql_query

[+] error interface
[✓] Error() = github.com.gen0cide.renum.cmd.renum.example.lib.error_code_invalid_sql_query (2): The provided query was not valid SQL.
$
```

We've effectively created a situation where errors are isolated into their namespace - they have identity, lineage, descriptive information, and satisfy the interface correctly. And of course, great [Godocs](https://godoc.org/github.com/gen0cide/renum/example/lib). This is the power of strongly typed enums in Go.

## Generating Enums w/ CLI

To install the CLI:

```sh
go get github.com/gen0cide/renum/cmd/renum
```

### YAML Configuration Format

An example and YAML configuration file outlining all options can be found in the `config_spec.yaml` file inside this repo.

### Example

Below shows an example of how you can write your enums in YAML, then with the renum codegen tool, generate your Go code. The example files can be found in the examples folder of the CLI.

```sh
# write your YAML configuration file...
$ renum -c error_code.yaml generate
[✓] parsed configuration
[✓] initialized generator
[✓] generated Go code
[✓] successfully wrote code to generated_error_codes.go
$
```

## Library

To use the interfaces, simply install the library with:

```sh
go get github.com/gen0cide/renum
```

## Inspiration / Prior Works

Both [go-enum](https://github.com/abice/go-enum) and [enumer](https://github.com/alvaroloes/enumer) do very similar things, but have some shortcomings. Both of them rely on AST parsing - meaning if your code is not parsable due to errors, you cannot generate your enums. Secondly, they don't provide easy mechanisms to enrich your types with additional methods and functionality.

This project started as a fork of `go-enum`, but ended up on it's own trajectory given the number of interfaces I wanted the enums to be able to implement. Their work is great and I think they have their uses, but were too limited to implement the strict paradigm of `renum`.

## Author

* <https://github.com/gen0cide>
* <https://twitter.com/alexlevinson>
* <https://www.linkedin.com/in/alexlevinson/>

## Shoutouts

* mbm
* davehughes
* ychen
* emperorcow
* m0
* vyrus001
* hecfblog
