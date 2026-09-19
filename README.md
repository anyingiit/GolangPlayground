<!-- Source: Best-README-Template BLANK_README (Unlicense) — https://github.com/othneildrew/Best-README-Template -->
<a id="readme-top"></a>

# GolangPlayground

A personal Go practice playground holding one working exercise, a TCP server that streams the time of day to any client that connects, one empty unfinished exercise stub, and two placeholder directories reserved for future exercises, rather than a single runnable application.

**English** · [简体中文](README.zh-CN.md)

[![License](https://img.shields.io/github/license/anyingiit/GolangPlayground)](LICENSE)

[Report a bug](https://github.com/anyingiit/GolangPlayground/issues/new?template=bug_report.yml) · [Request a feature](https://github.com/anyingiit/GolangPlayground/issues/new?template=feature_request.yml)

<details>
  <summary>Table of Contents</summary>
  <ol>
    <li><a href="#about-the-project">About The Project</a></li>
    <li><a href="#getting-started">Getting Started</a></li>
    <li><a href="#usage">Usage</a></li>
    <li><a href="#contributing">Contributing</a></li>
    <li><a href="#license">License</a></li>
    <li><a href="#contact">Contact</a></li>
  </ol>
</details>

## About The Project

GolangPlayground is a personal collection of small, independent Go exercises rather than one application. The only exercise that runs is `doanything/2022090801/clock2/clock2.go`, a TCP server that listens on port 8080 and, for every client that connects, writes the current time in a tight loop with no pacing between writes until the connection closes. Its companion, `doanything/2022090801/clockwall/clockwall.go`, is an abandoned stub -- a bare `package main` declaration with no function bodies at all -- and the top-level `doanything/` and `questions/` directories otherwise hold nothing but a `.gitkeep` placeholder each.

See the [open issues](https://github.com/anyingiit/GolangPlayground/issues) for planned features and known issues.

## Getting Started

### Prerequisites

- Git, to clone the repository.
- A Go toolchain. The repository has no `go.mod` anywhere, so nothing is fetched as a dependency; the one working exercise imports only the standard library (`fmt`, `io`, `log`, `net`, `time`) and is built directly from its single source file.

### Installation

```sh
git clone https://github.com/anyingiit/GolangPlayground.git
cd GolangPlayground/doanything/2022090801/clock2
go build clock2.go
```

There is no module anywhere in this repository, so each exercise is built straight from its own `.go` file rather than through `go install` or a package path.

## Usage

```sh
./clock2 &
nc localhost 8080
```

The server accepts a connection on port 8080 and immediately starts writing the current time to it, formatted as `15:06:07`, over and over with no delay between writes, until the client disconnects. There is nothing to type at the server; press Ctrl-C to stop the `nc` client when you are done. `doanything/2022090801/clockwall/clockwall.go` has no `func main` to run at all -- it is an unfinished stub, not a second usable program.

## Contributing

Contributions are welcome. Read [CONTRIBUTING.md](CONTRIBUTING.md) for how to open an issue or a pull request, and [CODE_OF_CONDUCT.md](CODE_OF_CONDUCT.md) for the standards expected of everyone taking part.

Please do not report security issues in public issues or pull requests. [SECURITY.md](SECURITY.md) explains how to report them privately.

## License

Distributed under the MIT License. See [LICENSE](LICENSE) for details.

## Contact

Project link: [https://github.com/anyingiit/GolangPlayground](https://github.com/anyingiit/GolangPlayground)

<p align="right">(<a href="#readme-top">back to top</a>)</p>
