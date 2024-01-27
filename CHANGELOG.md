<!-- markdownlint-disable MD012 MD013 MD022 MD024 MD032 MD033 -->

# CHANGELOG

This CHANGELOG is a format conforming to [keep-a-changelog](https://github.com/olivierlacan/keep-a-changelog).

<a name="unreleased"></a>
## [Unreleased]


<a name="v0.1.6"></a>
## [v0.1.6] - 2024-01-27
### Bug Fixes
- go version in go.mod

### Other Work
- bump github actions


<a name="v0.1.5"></a>
## [v0.1.5] - 2024-01-27
### Other Work
- update changelog for v0.1.5
- bump go version to v1.21.6
- add goland to .gitignore
- bump go module dependencies
- update pre-commit hooks to the latest versions


<a name="v0.1.4"></a>
## [v0.1.4] - 2023-06-16
### Build Process
- fix wrong url in release footer

### Other Work
- update changelog for v0.1.4


<a name="v0.1.3"></a>
## [v0.1.3] - 2023-06-16
### Build Process
- rename binary consisten to 'go-branchy'
- add variable for CGO

### Other Work
- update changelog for v0.1.3


<a name="v0.1.2"></a>
## [v0.1.2] - 2023-06-11
### Refactoring
- replace golang.design/x/clipboard with github.com/atotto/clipboard to prevent CGO dependency and goreleaser cross-compile hassle

### Build Process
- removed CGO and cross-compile dependencies after refactoring

### Other Work
- update changelog for v0.1.2


<a name="v0.1.1"></a>
## [v0.1.1] - 2023-06-11
### Bug Fixes
- broken builds on MacOS due to cgo dependencies for the clipboard module

### Other Work
- update changelog for v0.1.1


<a name="v0.1.0"></a>
## [v0.1.0] - 2023-06-10
### New Features
- add first version of go-branchy

### Documentation
- update README.md and add go-branchy logo without transparent background
- update README.md
- update README.md
- update README.md
- update README.md

### Build Process
- add Taskfile.yml
- add .goreleaser.yaml and github action release workflow

### Other Work
- update changelog for v0.1.0
- add git-chglog template, config and an empty CHANGELOG.md
- add LICENSE
- add .gitignore and .pre-commit-config.yaml


<a name="v0.0.0"></a>
## v0.0.0 - 2023-06-08
### Other Work
- initial commit


[Unreleased]: https://github.com/untcha/go-branchy/compare/v0.1.6...HEAD
[v0.1.6]: https://github.com/untcha/go-branchy/compare/v0.1.5...v0.1.6
[v0.1.5]: https://github.com/untcha/go-branchy/compare/v0.1.4...v0.1.5
[v0.1.4]: https://github.com/untcha/go-branchy/compare/v0.1.3...v0.1.4
[v0.1.3]: https://github.com/untcha/go-branchy/compare/v0.1.2...v0.1.3
[v0.1.2]: https://github.com/untcha/go-branchy/compare/v0.1.1...v0.1.2
[v0.1.1]: https://github.com/untcha/go-branchy/compare/v0.1.0...v0.1.1
[v0.1.0]: https://github.com/untcha/go-branchy/compare/v0.0.0...v0.1.0
