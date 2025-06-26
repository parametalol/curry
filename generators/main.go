package main

import (
	_ "embed"
	"fmt"
	"os"
	"slices"
	"strconv"
	"text/template"
)

var (
	//go:embed header.tmpl
	headerTemplate string
	//go:embed bind.tmpl
	bindTemplate string
	//go:embed reverse.tmpl
	reverseTemplate string

	headerT   = template.Must(template.New("header").Parse(headerTemplate))
	templates = map[string]string{
		"bind":    bindTemplate,
		"reverse": reverseTemplate,
	}
)

type Arg struct {
	I    int
	Name string
}

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "usage: go run generators <num_params>")
		os.Exit(1)
	}
	n, err := strconv.Atoi(os.Args[1])
	if err != nil || n < 1 {
		fmt.Fprintln(os.Stderr, "num_params must be a positive integer")
		os.Exit(1)
	}
	n++
	args := make([]Arg, n)
	for i := range n {
		args[i].I = i
		args[i].Name = fmt.Sprint("arg", i)
	}
	rargs := slices.Clone(args)
	slices.Reverse(rargs)
	for name, value := range templates {
		tmpl := template.Must(template.New(name).Parse(value))
		file, err := os.Create(fmt.Sprintf("../%s_gen.go", name))
		if err != nil {
			panic(err)
		}
		defer file.Close()
		headerT.Execute(file, nil)
		for i := range n - 1 {
			tmpl.Execute(file, map[string]any{
				"N":     i + 1,
				"Last":  i,
				"Args":  args[0 : i+1],
				"RArgs": rargs[n-i-1 : n],
			})
		}
	}
}
