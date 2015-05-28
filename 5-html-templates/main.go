package main

import (
	"fmt"
	"html/template"
	"os"
)

type Article struct {
	Name       string
	AuthorName string
	Draft      bool
}

func (a Article) Byline() string {
	return fmt.Sprintf("Written by %s", a.AuthorName)
}

func main() {

	goArticle := Article{
		Name:       "The Go html/template package",
		AuthorName: "Mal Curtis",
	}

	tmpl, err := template.New("Foo").Parse("'{{.Name}}' {{if .Draft}} (Draft) {{else}} (Published) {{end}} {{.Byline}}\n")
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(os.Stdout, goArticle)
	if err != nil {
		panic(err)
	}
}
