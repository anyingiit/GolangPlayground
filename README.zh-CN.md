[English](README.md) · **简体中文**

> 英文版是规范版本。本页与 [README.md](README.md) 不一致时，以英文版为准。

<!-- translation-of: README.md sha256:d876e6b9ae532a68 -->

<!-- Source: Best-README-Template BLANK_README (Unlicense) — https://github.com/othneildrew/Best-README-Template -->
<a id="readme-top"></a>

# GolangPlayground

一个个人的 Go 练习场，其中只有一个真正能运行的练习——一个把当前时间持续发送给任意连接客户端的 TCP 服务器——外加一个未完成的空练习桩，以及两个留给未来练习的占位目录，而不是一个单一的可运行应用。

[![License](https://img.shields.io/github/license/anyingiit/GolangPlayground)](LICENSE)

[报告问题](https://github.com/anyingiit/GolangPlayground/issues/new?template=bug_report.yml) · [提出需求](https://github.com/anyingiit/GolangPlayground/issues/new?template=feature_request.yml)

<details>
  <summary>目录</summary>
  <ol>
    <li><a href="#about-the-project">关于本项目</a></li>
    <li><a href="#getting-started">开始使用</a></li>
    <li><a href="#usage">用法</a></li>
    <li><a href="#contributing">参与贡献</a></li>
    <li><a href="#license">许可证</a></li>
    <li><a href="#contact">联系方式</a></li>
  </ol>
</details>

## 关于本项目

GolangPlayground 是一批彼此独立的小型 Go 练习，而不是一个应用程序。唯一能运行的练习是 `doanything/2022090801/clock2/clock2.go`，它是一个监听 8080 端口的 TCP 服务器：每当有客户端连接上来，它就会在一个没有任何间隔的紧密循环里不断写入当前时间，直到连接断开。与它相邻的 `doanything/2022090801/clockwall/clockwall.go` 是一个被放弃的空文件——只有一行 `package main` 声明，没有任何函数体——而顶层的 `doanything/` 和 `questions/` 目录里，也分别只剩下一个 `.gitkeep` 占位文件。

计划中的功能与已知问题，见 [open issues](https://github.com/anyingiit/GolangPlayground/issues)。

## 开始使用

### 环境要求

- Git，用于克隆本仓库。
- 一个 Go 工具链。仓库里任何地方都没有 `go.mod`，所以不需要拉取任何依赖；唯一能运行的练习只导入了标准库（`fmt`、`io`、`log`、`net`、`time`），直接从它自己的单个源文件构建即可。

### 安装

```sh
git clone https://github.com/anyingiit/GolangPlayground.git
cd GolangPlayground/doanything/2022090801/clock2
go build clock2.go
```

这个仓库里没有任何模块（module），因此每个练习都是直接从自己的 `.go` 文件构建的，而不是通过 `go install` 或某个包路径。

## 用法

```sh
./clock2 &
nc localhost 8080
```

服务器在 8080 端口接受连接后，会立即开始向连接写入当前时间，格式为 `15:06:07`，不做任何等待地反复写入，直到客户端断开连接为止。服务器本身没有任何需要输入的内容；用完之后按 Ctrl-C 结束 `nc` 客户端即可。`doanything/2022090801/clockwall/clockwall.go` 根本没有 `func main`，无法运行——它只是一个未完成的空文件，而不是第二个可用的程序。

## 参与贡献

欢迎参与。[CONTRIBUTING.md](CONTRIBUTING.md) 说明如何提交 issue 或 pull request，[CODE_OF_CONDUCT.md](CODE_OF_CONDUCT.md) 说明对所有参与者的行为要求。

请不要在公开的 issue 或 pull request 中报告安全问题。[SECURITY.md](SECURITY.md) 说明了私下报告的方式。

## 许可证

以 MIT 许可证分发。详见 [LICENSE](LICENSE)。

## 联系方式

项目地址：[https://github.com/anyingiit/GolangPlayground](https://github.com/anyingiit/GolangPlayground)

<p align="right">(<a href="#readme-top">back to top</a>)</p>
