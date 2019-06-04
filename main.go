package main

import (
	"html/template"
	"log"
	"os"

	"github.com/gobuffalo/packr"
)

func main() {
	box := packr.NewBox("./templates")
	t, err := template.New("dalongdemo").Parse(box.String("app.t"))
	if err != nil {
		log.Fatal(err)
	}
	err = t.Execute(os.Stdout, "dalongdemo")
	if err != nil {
		log.Fatal(err)
	}
}
