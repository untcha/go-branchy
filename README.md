<!-- markdownlint-disable -->

<p align="center">
  <img alt="go-branchy logo" src="assets/branchy-logo-transparent.png" height="150" />
  <h3 align="center">go-branchy</h3>
  <p align="center">Written in Go...</p>
</p>

---

`go-branchy` or `branchy` is a CLI helper which generates/creates git branch
names from JIRA tickets (summary field) and automatically copies the branch name
to the clipboard.

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
go-branchy generate feat PROJ-1234
```
