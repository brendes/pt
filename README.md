# pt

[![Build Status](https://travis-ci.org/hoffa/pt.svg?branch=master)](https://travis-ci.org/hoffa/pt)
[![Go Report Card](https://goreportcard.com/badge/github.com/hoffa/pt)](https://goreportcard.com/report/github.com/hoffa/pt)

A minimalist static blog generator.

- Super simple
- Write pages in [Markdown](https://daringfireball.net/projects/markdown/syntax)
- Generates valid [RSS 2.0](https://validator.w3.org/feed/docs/rss2.html) feed

## Installation

```shell
go get github.com/hoffa/pt
```

## Usage

```shell
pt -base-url https://my.site *.md
```

## Front matter

```markdown
---
title: Hello, world!
date: 2019-02-11
---
```

Each page can contain a YAML front matter. It must be placed at the top within `---` delimiters.

Valid variables are:

- `title`: the content title
- `date`: the content creation date
- `exclude`: if `yes`, the page won't be included in `.Pages`

## Themes

Just add your CSS in the `<head>`.
For example:

```css
body {
  line-height: 1.5;
  max-width: 40em;
  margin: auto;
  padding: 1em;
}
```

## Example

First, get the page and RSS feed templates:

```shell
curl -L https://github.com/hoffa/pt/archive/master.tar.gz \
  | tar xz --strip-components=1 pt-master/templates
```

Create the index page as `index.md`:

```Markdown
---
title: Jane Doe
exclude: yes
---

Subscribe via [RSS](/feed.xml).
```

And a post within a file called `my-first-post.md`:

```Markdown
---
title: My first post
date: 2019-04-20
---

This is an example **Markdown** _post_.
I like `turtles`.
```

Finally, build:

```shell
pt *.md
```
