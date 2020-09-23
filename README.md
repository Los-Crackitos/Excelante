![Excelante](img/excelante.png)

![Code Quality Status](https://github.com/Los-Crackitos/Excelante/workflows/Code%20Quality/badge.svg?branch=master)
![Tests Status](https://github.com/Los-Crackitos/Excelante/workflows/Tests/badge.svg?branch=master)
[![Codecov](https://codecov.io/gh/Los-Crackitos/Excelante/branch/master/graph/badge.svg)](https://codecov.io/gh/Los-Crackitos/Excelante)
[![Go Report Card](https://goreportcard.com/badge/github.com/Los-Crackitos/Excelante)](https://goreportcard.com/report/github.com/Los-Crackitos/Excelante)
![Code Size](https://img.shields.io/github/languages/code-size/Los-Crackitos/Excelante)
[![Go Doc](https://godoc.org/github.com/Los-Crackitos/Excelante?status.svg)](https://godoc.org/github.com/Los-Crackitos/Excelante)
[![License](https://img.shields.io/github/license/Los-Crackitos/Excelante)](LICENSE)

Excelante is an open source API that allow users to interact with XLSX files, extract data from it or create files from a given input. Only a HTTP request is needed to use our API, it can be useful for website/application that needs to deal with XLSX files or for automatization.

## Documentation

All API documentation can be found at [Excelante Doc](https://app.gitbook.com/@loscrackitos/s/excelante/)

Swagger API documentation can be found at `/api/v1/swagger/`

## Benchmark

Benchmark was realized on MacBook Pro (2015) laptop:
2,7 GHz Intel Core i5 double core
8 Go 1867 MHz DDR3

### Writer

| Tests                       | Run 1   | Run 2   | Run 3   | Run 4     | Run 5    | Moy     |
| --------------------------- | ------- | ------- | ------- | --------- | -------- | ------- |
| 10x10 (100 cells)           | 382ms   | 361ms   | 249ms   | 269ms     | 214ms    | 295ms   |
| 100x100 (10 000 cells)      | 627ms   | 501ms   | 517ms   | 512ms     | 475ms    | 5264ms  |
| 1000x1000 (1 000 000 cells) | 83,299s | 95,253s | 81,792s | 104,130ms | 101,336s | 93,162s |

### Reader

| Tests                           | Run 1          | Run 2          | Run 3          | Run 4          | Run 5          | Moy     |
| ------------------------------- | -------------- | -------------- | -------------- | -------------- | -------------- | ------- |
| 15(cols)x25(rows) (375 cells)   | columns: 168ms | columns: 140ms | columns: 199ms | columns: 224ms | columns: 90ms  | 3262ms  |
|                                 | lines: 45ms    | lines: 28ms    | lines: 39ms    | lines: 79ms    | lines: 37ms    | 45,6ms  |
| 30(cols)x50(rows) (1500 cells)  | columns: 512ms | columns: 597ms | columns: 582ms | columns: 609ms | columns: 414ms | 542,8ms |
|                                 | lines: 48ms    | lines: 97ms    | lines: 69ms    | lines: 172ms   | lines: 104ms   | 98ms    |
| 60(cols)x100(rows) (6000 cells) | columns: 2,17s | columns: 2,74s | columns: 2,10s | columns: 2,14s | columns: 2,25s | 2,28s   |
|                                 | lines: 167ms   | lines: 146ms   | lines: 392ms   | lines: 201ms   | lines: 241ms   | 229,4ms |

## Website

Work in progress [here](https://github.com/Los-Crackitos/excelante-site)

## Want to contribute

### Getting Started

Clone the repository and run `make init` to copy all configuration files.
Then, run `docker-compose up` to start the project.

### Roadmap

[See the roadmap »](roadmap.md)

### Code of conduct

As contributors and maintainers of this project, we pledge to respect all people who contribute through reporting issues, posting feature requests, updating documentation, submitting pull requests or patches, and other activities.

We are committed to making participation in this project a harassment-free experience for everyone, regardless of level of experience, gender, gender identity and expression, sexual orientation, disability, personal appearance, body size, race, ethnicity, age, or religion.

Examples of unacceptable behavior by participants include the use of sexual language or imagery, derogatory comments or personal attacks, trolling, public or private harassment, insults, or other unprofessional conduct.

Project maintainers have the right and responsibility to remove, edit, or reject comments, commits, code, wiki edits, issues, and other contributions that are not aligned to this Code of Conduct. Project maintainers who do not follow the Code of Conduct may be removed from the project team.

Instances of abusive, harassing, or otherwise unacceptable behavior may be reported by opening an issue or contacting one or more of the project maintainers.

This Code of Conduct is adapted from the Contributor Covenant, version 1.0.0, available at <https://www.contributor-covenant.org/version/1/0/0/code-of-conduct.html>

## Sponsors

Support our open source work 😄

## License

Excelante is under MIT license. Feel free to use it as you want.
