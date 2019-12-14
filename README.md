# mdtable2csv

Convert markdown table to csv.

# Installation

```
$ go get -u github.com/515hikaru/mdtable2csv
```

# Usage

```
$ mdtable2csv
|foo|bar|boo|
|---|---|---|
|a|b|c|^D
foo,bar,boo
a,b,c
```

Of course, you can use standard input with `|`, for example:

```
$ cat foo.md
|foo|bar|boo|
|---|---|---|
| a | b | c |

$ cat foo.md | mdtable2csv
foo,bar,boo
a,b,c
```

# LICENSE

MIT
