package render

import (
	"html/template"
	"io/fs"

	"github.com/ed-henrique/voz/internal/errkit"
	"github.com/ed-henrique/voz/internal/shortener"
)

var tmplFuncs = template.FuncMap{
	"shorten": shortener.ShortenNumber,
}

func Render(views fs.FS, files ...string) *template.Template {
	mainFilename := files[0]
	templatePaths := make([]string, len(files))
	for i, filename := range files {
		templatePaths[i] = "views/" + filename
	}

	tmpl, err := template.New(mainFilename).Funcs(tmplFuncs).ParseFS(views, templatePaths...)
	if err != nil {
		errkit.FinalErr(err.Error())
		return nil
	}

	return tmpl
}
