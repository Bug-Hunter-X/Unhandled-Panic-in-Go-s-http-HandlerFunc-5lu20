# Unhandled Panic in Go's http.HandlerFunc

This repository demonstrates a common error in Go web servers: unhandled panics within `http.HandlerFunc`.

When a panic occurs within a handler, the entire server crashes unless proper error handling is implemented.

The `bug.go` file showcases the problem, while `bugSolution.go` provides a solution using `recover()`.

## Bug

The `handleRequest` function in `bug.go` intentionally panics. Running this code will result in the server crashing.

## Solution

`bugSolution.go` demonstrates how to use `recover()` within a `defer` statement to catch panics and gracefully handle them, preventing the server from crashing.
