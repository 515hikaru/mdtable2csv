![Go](https://github.com/515hikaru/mdtable2csv/workflows/Go/badge.svg)
![release](https://github.com/515hikaru/mdtable2csv/workflows/release/badge.svg)

# mdtable2csv

Convert markdown table to csv.

## Screen shot

![](https://cdn-ak.f.st-hatena.com/images/fotolife/h/hikaru515/20191228/20191228175525.gif)

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
