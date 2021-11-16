package page

import (
	"github.com/Oppodelldog/go-simple-webapp-template/assets"
	"html/template"
	"path"
	"path/filepath"
)

const templatesBasePath = "templates"

func getPartials() []string {

	partialsPath := path.Join(templatesBasePath, "partials")
	return []string{
		path.Join(partialsPath, "logo.html"),
		path.Join(partialsPath, "head.html"),
	}
}

func NewPageTemplate(pageFileName string) (*template.Template, error) {
	pageFilePath := path.Join(templatesBasePath, pageFileName)
	name := filepath.Base(pageFilePath)

	t := template.New(name)

	b, err := assets.Templates.ReadFile(pageFilePath)
	if err != nil {
		return nil, err
	}
	_, err = t.Parse(string(b))
	if err != nil {
		return nil, err
	}

	for _, partial := range getPartials() {
		b, err := assets.Templates.ReadFile(partial)
		if err != nil {
			return nil, err
		}
		tmpl := t.New(partial)
		_, err = tmpl.Parse(string(b))
		if err != nil {
			return nil, err
		}
	}

	return t, nil
}
