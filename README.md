<!-- markdownlint-disable -->

<p align="center">
  <img alt="go-branchy logo" src="assets/branchy-logo-transparent.png" height="150" />
  <h3 align="center">go-branchy</h3>
  <p align="center">Written in Go... Logo authored by Midjourney</p>
</p>

---

`go-branchy` or `branchy` is a CLI helper which generates/creates git branch
names from JIRA tickets (summary field) and automatically copies the branch name
to the clipboard.

## Motivation

This project has two motivational aspects. The first is that we use a common
pattern to create git branch names at work (and we use JIRA... of course). 
So this is why I created a little tool to create a branch name by just giving
the branch type (e.g. feat or fix) and a JIRA issue (e.g. ABC-1234).
The result should be the following:

``` shell
feat/ABC-1234/super-cool-branch-name-created-from-a-jira-summary
```

The second aspect is the learning aspect. This little tool and the whole repository
seem to be a little overloaded for what it really achieves. But I had different
goals on what I wanted to learn.
These where the things I wanted to learn or at least try out:

* Building a CLI tool with [Cobra](https://cobra.dev/)
* Using [Task](https://taskfile.dev/) instead of a `Makefile`
* Automating the release process with [GoReleaser](https://goreleaser.com/) and [GitHub Actions](https://github.com/features/actions)
* Generating the `CHANGELOG.md` automatically with [git-chglog](https://github.com/git-chglog/git-chglog)
* Generating a beautiful logo with `ChatGPT` and `Midjourney`
* And of course also utilizing `GitHub Copilot`

## Installation

### Go

``` shell
go install github.com/untcha/go-branchy@latest
```

### Download the binary

You can download the binary from the [GitHub releases page](https://github.com/untcha/go-branchy/releases) and add it to your `$PATH`

The `go-branchy_<version>_checksums.txt` file contains the SHA-256 checksum for each file.

## Configuration

In order to use `go-branchy` you need to export the following two environment
variables:

``` shell
export BRANCHY_JIRA_URL=<the-url-of-your-jira-instance>
export BRANCHY_JIRA_TOKEN=<your-jira-personal-access-token>
```

## Usage

Now you can use `branchy` very simple:

``` shell
go-branchy generate feat ABC-1234
```

or use the alias `g` instead of generate:

``` shell
go-branchy g feat ABC-1234
```

### Help

```
go-branchy help
```

```
A little helper tool to create git branch names from JIRA tickets - written in Go

Branchy is a CLI helper which generates/creates git branch names from JIRA tickets (summary field)
and automatically copies the branch name to the clipboard (e.g. feat/ABC-1234/this-is-my-branch-name)

Usage:
  branchy [command]

Examples:
branchy generate feat ABC-1234
branchy g fix ABC-1234

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  generate    Generate a branch name from a JIRA issue summary field
  help        Help about any command
  version     Print the CLI version

Flags:
  -h, --help   help for branchy

Use "branchy [command] --help" for more information about a command.
```
