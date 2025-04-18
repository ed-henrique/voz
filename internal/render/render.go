package render

import (
	"html/template"
	"io/fs"
	"path"

	"github.com/ed-henrique/voz/internal/errkit"
	"github.com/ed-henrique/voz/internal/shortener"
)

var tmplFuncs = template.FuncMap{
	"shorten": shortener.ShortenNumber,
}

func Load(views fs.FS, files ...string) *template.Template {
	mainFilename := path.Base(files[0])
	templatePaths := make([]string, len(files))
	for i, filename := range files {
		templatePaths[i] = "views/" + filename
	}

	tmpl, err := template.New(mainFilename).Funcs(tmplFuncs).ParseFS(views, templatePaths...)
	if err != nil {
		errkit.FinalErr(err)
		return nil
	}

	return tmpl
}
