# Releasing

This tool uses Go Releaser to manage release builds.

## Setup

Install Go Releaser.

```bash
brew install goreleaser/tap/goreleaser
```

## Releasing

Tag the repo:

```bash
$ git tag -a v0.1.0 -m "release tag."
$ git push origin v0.1.0
```

Then:

```bash
goreleaser --snapshot
```
