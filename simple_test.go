package go_template_demos

import (
	"os"
	"testing"
	"text/template"
)

func TestSimple(t *testing.T) {
	type Inventory struct {
		Material string
		Count    uint
	}
	sweaters := Inventory{"wool", 17}
	tmpl, err := template.New("test").Parse("{{.Count}} items are made of {{.Material}}")
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(os.Stdout, sweaters)
	if err != nil {
		panic(err)
	}
}

func TestSpace(t *testing.T) {
	tmpl, err := template.New("space").Parse("{{23 -}} < {{- 45}}")
	if err != nil {
		panic(err)
	}
	var s struct{}
	err = tmpl.Execute(os.Stdout, s)
	if err != nil {
		panic(err)
	}
}
