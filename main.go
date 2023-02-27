package main

import (
	"cvgo/cv"
	"cvgo/renderer"
	"html/template"

	"fmt"
)

func main() {
	myCV := cv.CV{}
	template.New()
	tmpl, err := template.ParseFiles()
	if err != nil {
		panic(err)
	}
	render := renderer.Renderer{}
	render.ToHtml("cv.html", &myCV)
	fmt.Printf("cv: %v\n", myCV)
}
