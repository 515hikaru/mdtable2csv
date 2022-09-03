![Go](https://github.com/515hikaru/mdtable2csv/workflows/Go/badge.svg)
![release](https://github.com/515hikaru/mdtable2csv/workflows/release/badge.svg)

# mdtable2csv

Convert the markdown table to CSV.

## Screenshot

![](https://cdn-ak.f.st-hatena.com/images/fotolife/h/hikaru515/20191228/20191228175525.gif)

# Installation

If you use Homebrew, run the following:

```
$ brew tap 515hikaru/tap
$ brew install 515hikaru/tap/mdtable2csv
```

Or `go get`:

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

See `mdtable2csv -help` for more details.

# LICENSE

MIT
