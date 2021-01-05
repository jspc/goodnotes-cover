package main

import (
	"flag"
	"io/ioutil"
	"log"
	"os"
	"text/template"
	"time"
)

var (
	line1 = flag.String("line1", "A. N. Other", "The first line to print on the cover")
	line2 = flag.String("line2", time.Now().Format("2006"), "The second line of text on the cover")

	inputFile  = flag.String("input", "latex/cover.tex.tmpl", "Base template")
	outputFile = flag.String("output", "build/cover.tex", "File to output to")
)

type vars struct {
	Line1 string
	Line2 string
}

func main() {
	flag.Parse()

	input, err := ioutil.ReadFile(*inputFile)
	if err != nil {
		log.Panic(err)
	}

	t, err := template.New("input").Parse(string(input))
	if err != nil {
		log.Panic(err)
	}

	v := vars{
		Line1: *line1,
		Line2: *line2,
	}

	output, err := os.Create(*outputFile)
	if err != nil {
		log.Panic(err)
	}

	err = t.Execute(output, v)
	if err != nil {
		log.Panic(err)
	}

}
