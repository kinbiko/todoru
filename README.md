# TODOる

[![Build Status](https://github.com/kinbiko/todoru/workflows/Go/badge.svg)](https://github.com/kinbiko/todoru/actions)
[![Coverage Status](https://coveralls.io/repos/github/kinbiko/todoru/badge.svg?branch=main)](https://coveralls.io/github/kinbiko/todoru?branch=main)
[![Go Report Card](https://goreportcard.com/badge/github.com/kinbiko/todoru)](https://goreportcard.com/report/github.com/kinbiko/todoru)
[![Latest version](https://img.shields.io/github/tag/kinbiko/todoru.svg?label=latest%20version&style=flat)](https://github.com/kinbiko/todoru/releases)
[![Go Documentation](http://img.shields.io/badge/godoc-documentation-blue.svg?style=flat)](https://pkg.go.dev/github.com/kinbiko/todoru?tab=doc)
[![License](https://img.shields.io/github/license/kinbiko/todoru.svg?style=flat)](https://github.com/kinbiko/todoru/blob/master/LICENSE)

`todoru` is a stack-based To Do list for the command line. Great for integrating with other tools (Alfred workflow included).

## Usage

Add items with `todoru <task>`.
Get the top of the stack with `todoru`.
Pop the top of the stack with `todoru pop`.

```console
$ todoru solve world hunger
$ todoru fix global warming
$ todoru refactor legacy codebase

$ todoru
refactor legacy codebase
$ todoru pop
popped "refactor legacy codebase"

$ todoru
fix global warming
$ todoru pop
popped "fix global warming"

$ todoru
solve world hunger
$ todoru pop
popped "solve world hunger"

$ todoru
Nothing left to do!

$ todoru pop
Stack is empty. Nothing to pop
```
