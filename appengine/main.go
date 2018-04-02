package main

import (
	"CardGame/controller"
	"html/template"
	"os"

	"google.golang.org/appengine"
)

var allTemplates map[string]*template.Template

func main() {
	allTemplates = populateTemplates()
	controller.Startup(allTemplates)
	appengine.Main()
}

func populateTemplates() map[string]*template.Template {
	result := make(map[string]*template.Template)
	const basePath = "templates"
	dir, err := os.Open(basePath)
	if err != nil {
		panic("Failed to open template blocks directory: " + err.Error())
	}
	fis, err := dir.Readdir(-1)
	if err != nil {
		panic("Failed to read contents of content directory: " + err.Error())
	}
	for _, fi := range fis {
		/*		f, err := os.Open(basePath + "/" + fi.Name())
				if err != nil {
					panic("Failed to open template '" + fi.Name() + "'")
				}
				content, err := ioutil.ReadAll(f)
				if err != nil {
					panic("Failed to read content from file '" + fi.Name() + "'")
				}
				f.Close() */
		tmpl := template.Must(template.ParseFiles(basePath + "/" + fi.Name()))
		result[fi.Name()] = tmpl
	}
	return result
}
