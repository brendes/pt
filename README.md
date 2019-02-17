# pt

[![Build Status](https://travis-ci.org/hoffa/pt.svg?branch=master)](https://travis-ci.org/hoffa/pt)

## Features

- Extremely straightforward
- Tiny
- Accessible

## Installation

```shell
go get github.com/hoffa/pt
```

## Usage

```shell
pt
```

`pt` takes no arguments. Markdown files within the working directory are converted to HTML.

## Front matter

Front matter is in TOML. It must be included within an opening and closing `+++` before the content.

### Example

```toml
title = "PKI for busy people"
date = 2019-02-11
```

### Variables

#### title

The content title.

#### description

The content description.

#### date

The content creation date.

#### exclude

If `true`, the page won't be included in `.Site.Pages`.
