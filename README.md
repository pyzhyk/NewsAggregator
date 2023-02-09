# [News Aggregator](https://github.com/pyzhyk/NewsAggregator/)

[![License](https://img.shields.io/badge/license-GPL-yellow.svg)][license]

[license]: https://www.gnu.org/licenses/gpl.html


#### Simple RSS news aggregator written in Go

## Prerequisites

- Go — [golang.org/doc/install](https://golang.org/doc/install)

## Getting started

- Add your favorite RSS feeds to `news.txt`
- Run aggregator:
```bash
go run main.go news.txt
```

![Screenshot-1](Images/Screenshot-1.png)

#### Default port is 800. It can be changed in main.go at line 75 (	`var Port string = "800"` ).
