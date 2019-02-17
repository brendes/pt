# pt

[![Build Status](https://travis-ci.org/hoffa/pt.svg?branch=master)](https://travis-ci.org/hoffa/pt) [![Go Report Card](https://goreportcard.com/badge/github.com/hoffa/pt)](https://goreportcard.com/report/github.com/hoffa/pt)

A minimalist static site generator.

Just converts [Markdown](https://daringfireball.net/projects/markdown/syntax) to HTML. Uses Go's [`text/template`](https://golang.org/pkg/text/template/).

## Features

- Tiny
- Straightforward
- Single-file template
- Generates pages instantly
- Generates RSS feed

## Installation

```shell
go get github.com/hoffa/pt
```

## Usage

```shell
pt
```

`pt` takes no arguments.

Markdown files within the working directory are converted to HTML.

## Demo

My [website](https://rehn.me). Generated from [hoffa/hoffa.github.io](https://github.com/hoffa/hoffa.github.io).

## Configuration

Site variables are defined in `pt.toml`. They're accessible through `.Site`.

### Example

```toml
author = "Jane Doe"
email = "jane@doe.com"
baseURL = "https://doe.com"
```

### Variables

- `author`: the site author
- `email`: the author's email
- `baseURL`: the URL used to combine with relative paths

## Front matter

Front matter is in [TOML](https://github.com/toml-lang/toml).

It must be included before the content within an opening and closing `+++`.

In case of missing values, sensible defaults will be used.

### Example

```toml
title = "Hello, world!"
date = 2019-02-11
```

### Variables

- `title`: the content title
- `description`: the content description
- `date`: the content creation date
- `exclude`: if `true`, the page won't be included in `.Site.Pages`
