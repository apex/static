<img src="http://tjholowaychuk.com:6000/svg/title/STATIC/ANTI FRAMEWORK">

The goal of this project is to build a common toolkit for building static sites for a variety of domains. Each program is tailored to one domain, such as a blog, documentation site, photo gallery, and so on. For example. Focusing the UX on each domain as necessary makes for a smoother experience, nicer content structuring, while maintaining an overall common feel and usability through the shared library. The shared library or "static stdlib" will also make it easy to write custom static generators to fit your needs.

I think this technique makes much more sense than fighting tools which are designed for blogs, as we all know how to write code, it can sometime save hours to just write a few lines instead of fighting a complex system that tries to do everything. I don't have much time for OSS right now, so it only has what I need, but hopefully long-term it'll turn into something real.

## Install

```bash
$ go get github.com/apex/static/cmd/static-docs
```

## Usage

The `static-docs` program generates a documentation website from a directory of markdown files. For example the [Up](https://apex.github.io/up/) documentation is generated with:

```
$ static-docs --in docs --out . --title Up --subtitle "Deploy serverless apps in seconds"
```

---

[![GoDoc](https://godoc.org/github.com/apex/static?status.svg)](https://godoc.org/github.com/apex/static)
![](https://img.shields.io/badge/license-MIT-blue.svg)
![](https://img.shields.io/badge/status-stable-green.svg)

<a href="https://apex.sh"><img src="http://tjholowaychuk.com:6000/svg/sponsor"></a>
