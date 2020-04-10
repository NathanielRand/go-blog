package views

import (
	"html/template"
	"path/filepath"
	"net/http"
)

var (
	LayoutDir = "views/layouts/"
	TemplateExt = ".gohtml"
)

func layoutFiles() []string {
	// Glob any files in the LayoutDir with TemplatesExt
	files, err := filepath.Glob(LayoutDir + "*" + TemplateExt)
	if err != nil {
		panic(err)
	}
	return files
}

func NewView(layout string, files ...string) *View {
	// Append list of files provided to common files.
	files = append(files, layoutFiles()...)
	// Parse the appended template files
	t, err := template.ParseFiles(files...)
	// Error check
	if err != nil {
		panic(err)
	}
	// Contrust pointer to view and return it.
	return &View{
		Template: t,
		Layout: layout,
	}
}

// View struct stores the parsed template that we want to execute.
type View struct {
	Template *template.Template
	Layout string
}

func(v *View) Render(w http.ResponseWriter, data interface{}) error {
	return v.Template.ExecuteTemplate(w, v.Layout, data)
}