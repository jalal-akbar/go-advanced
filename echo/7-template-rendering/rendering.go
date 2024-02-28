package templaterendering

import (
	"html/template"
	"io"

	"github.com/labstack/echo"
)

// Template Rendering
type M map[string]interface{}

type Renderer struct {
	template *template.Template
	debug    bool
	location string
}

func NewRender(location string, debug bool) *Renderer {
	tmpl := new(Renderer)

	tmpl.location = location
	tmpl.debug = debug

	tmpl.ReloadTemplates()

	return tmpl
}

func (t *Renderer) ReloadTemplates() {
	t.template = template.Must(template.ParseGlob(t.location))
}

func (t *Renderer) Render(w io.Writer, name string, data interface{}, ctx echo.Context) error {
	if t.debug {
		t.ReloadTemplates()
	}
	return t.template.ExecuteTemplate(w, name, data)
}
