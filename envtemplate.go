package main

import (
	"flag"
	"io"
	"os"
	"strings"
	"text/template"
)

func envToMap() map[string]string {
	data := os.Environ()
	items := make(map[string]string)
	for _, item := range data {
		splits := strings.Split(item, "=")
		items[splits[0]] = splits[1]
	}
	return items
}

func resolveEnv(writer io.Writer, file string) error {
	tpl, err := template.ParseFiles(file)
	if err != nil {
		return err
	}
	return tpl.Execute(writer, envToMap())
}

func main() {
	file := flag.String("file", "", "The file for template")
	flag.Parse()
	if *file == "" {
		panic("envtemplate -file=somefilepath")
	}

	resolveEnv(os.Stdout, *file)
}
