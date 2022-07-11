# istty

[![build-img]][build-url]
[![pkg-img]][pkg-url]
[![reportcard-img]][reportcard-url]
[![coverage-img]][coverage-url]
[![version-img]][version-url]

Go library to report when program runs in terminal.

## Features

* Simple API.
* Dependency-free.

## Install

Go version 1.17+

```
go get github.com/cristalhq/istty
```

## Example

```go
if istty.IsTerminal(os.Stdin.Fd()) {
	// do fancy stuff
}
```

## Documentation

See [these docs][pkg-url].

## License

[MIT License](LICENSE).

[build-img]: https://github.com/cristalhq/istty/workflows/build/badge.svg
[build-url]: https://github.com/cristalhq/istty/actions
[pkg-img]: https://pkg.go.dev/badge/cristalhq/istty
[pkg-url]: https://pkg.go.dev/github.com/cristalhq/istty
[reportcard-img]: https://goreportcard.com/badge/cristalhq/istty
[reportcard-url]: https://goreportcard.com/report/cristalhq/istty
[coverage-img]: https://codecov.io/gh/cristalhq/istty/branch/main/graph/badge.svg
[coverage-url]: https://codecov.io/gh/cristalhq/istty
[version-img]: https://img.shields.io/github/v/release/cristalhq/istty
[version-url]: https://github.com/cristalhq/istty/releases
