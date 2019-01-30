# Goodnotes Cover

This application will generate a goodnotes cover template that:

* Is 'embossed' with a name/text and a month/year it is used for
* Has a scuffed, faux leather style finish
* Has some masking tape lazily applied that can be used to write titles on

It is written in golang and is powered by a Makefile.

## Dependencies

1. golang
1. latex
1. make

## Usage

### Simple

```bash
$ make
```

This will:

1. Compile the binary
1. Generate 12 tex files, one per month
1. Generate each tex file into a pdf


### Complex(ish)

The `goodnotes-cover` app has a number of options which the Makefile leaves alone:

```bash
$ ./goodnotes-cover -h
Usage of ./goodnotes-cover:
  -input string
        Base template (default "latex/cover.tex.tmpl")
  -name string
        The name to emboss onto covers (default "A. N. Other")
  -output string
        Directory to output to (default "build/")
  -year int
        Year for which to generate (default 2019)
```

This allows for different templates, different names, and so on. It is recommended, though, that compilation to pdf is still done with:

```bash
$ make compile
```
