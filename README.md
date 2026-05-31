# phase-shift

phase-shift is a static analysis tool for Go.

The first analyzer checks functions annotated with `//phase:nonmutating`.
An annotated function must not mutate caller-visible state through parameters or method receivers.

## Installation

Standard `go install` installation,

```sh
GOPRIVATE=github.com/phasemerge go install github.com/phasemerge/go-phase-shift/cmd/phase-shift@latest
```

## Usage

Run `phase-shift` on your Go packages,

```sh
phase-shift ./...
```

## Example

This function is reported because it mutates through a pointer parameter,

```go
//phase:nonmutating
func F(p *int) {
	*p = 1
}
```

This function is allowed because it only reads through the pointer parameter,

```go
//phase:nonmutating
func F(p *int) int {
	return *p
}
```

This method is allowed because a scalar field assignment on a value receiver mutates only the receiver copy,

```go
type Count struct{ n int }

//phase:nonmutating
func (c Count) Increment() Count {
	c.n++
	return c
}
```

This method is reported because the slice field is mutated for the caller,

```go
type Buffer struct{ data []byte }

//phase:nonmutating
func (b Buffer) ClearFirstByte() {
	b.data[0] = 0
}
```

## Development

Run all fixers,

```sh
mise run fixers
```

Run all tests,

```sh
mise run tests
```

Run all linters,

```sh
mise run linters
```
