# ltries
This is a generalized tries package with clean and transparent API for the Go language.

## Features
- Lightweight and Fast.
- Native Go implementation.
- Support all utf-8 (characters).

## Requirements
- Go 1.9 or higher. We aim to support the 3 latest versions of Go.

## Installation
Simple install the package to your [$GOPATH](https://github.com/golang/go/wiki/GOPATH "GOPATH") with the [go tool](https://golang.org/cmd/go/ "go command") from shell:
```bash
$ go get -u github.com/lovesaroha/ltries
```
Make sure [Git is installed](https://git-scm.com/downloads) on your machine and in your system's `PATH`.

## Usage

### Create Trie

```Golang
  // Create a trie.
  trie := ltries.Create()

  // Insert words in trie with index.
  trie.Insert("love", 1)
  trie.Insert("saroha", 2)
  trie.Insert("git", 3)

  // To find index of given word (if not found return -1).
  index := trie.Get("love")

  // Remove given word from trie.
  trie.Remove("git")

```