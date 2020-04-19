package views

import (
	"html/template"
	"path/filepath"
	"net/http"
)

var (
	LayoutDir = "views/layouts/"
	TemplateDir = "views/"
	TemplateExt = ".gohtml"
)

// View struct stores the parsed template that we want to execute.
type View struct {
	Template *template.Template
	Layout string
}

func layoutFiles() []string {
	// Glob any files in the LayoutDir with TemplatesExt
	files, err := filepath.Glob(LayoutDir + "*" + TemplateExt)
	if err != nil {
		panic(err)
	}
	return files
}

func NewView(layout string, files ...string) *View {
	
	addTemplatePath(files)
	addTemplateExt(files)
	
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

func(v *View) Render(w http.ResponseWriter, data interface{}) error {
	return v.Template.ExecuteTemplate(w, v.Layout, data)
}

func(v *View) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := v.Render(w, nil); err != nil {
		panic(err)
	}
}

// addTemplatePath takes in a slice of strings
// representing file paths for templates, and it prepends // the TemplateDir directory to each string in the slice //
// Eg the input {"home"} would result in the output
// {"views/home"} if TemplateDir == "views/"
func addTemplatePath(files []string) {
	for i, f := range files {
	  	files[i] = TemplateDir + f
  	} 
}

  // addTemplateExt takes in a slice of strings
// representing file paths for templates and it appends // the TemplateExt extension to each string in the slice //
// Eg the input {"home"} would result in the output
// {"home.gohtml"} if TemplateExt == ".gohtml"
func addTemplateExt(files []string) {
  	for i, f := range files {
    	files[i] = f + TemplateExt
	} 
}
