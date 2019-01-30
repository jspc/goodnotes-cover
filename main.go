package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"
	"time"
)

var (
	name      = flag.String("name", "A. N. Other", "The name to emboss onto covers")
	year      = flag.Int("year", time.Now().Year(), "Year for which to generate")
	inputFile = flag.String("input", "latex/cover.tex.tmpl", "Base template")
	outputDir = flag.String("output", "build/", "Directory to output to")
)

type vars struct {
	Name      string
	Timestamp string
}

func main() {
	flag.Parse()

	err := os.MkdirAll(*outputDir, 0755)
	if err != nil {
		log.Panic(err)
	}

	v := vars{
		Name: *name,
	}

	input, err := ioutil.ReadFile(*inputFile)
	if err != nil {
		log.Panic(err)
	}

	t, err := template.New("input").Parse(string(input))
	if err != nil {
		log.Panic(err)
	}

	for _, m := range []string{
		"January", "February", "March", "April", "May", "June",
		"July", "August", "September", "October", "November", "December",
	} {
		v.Timestamp = fmt.Sprintf("%s %d", m, *year)

		filename := filepath.Join(*outputDir, fmt.Sprintf("%s.tex", strings.ToLower(v.Timestamp)))

		output, err := os.Create(filename)
		if err != nil {
			log.Panic(err)
		}

		err = t.Execute(output, v)
		if err != nil {
			log.Panic(err)
		}
	}
}
