package template

import (
	tpl "html/template"
	"os"
	"path/filepath"
	"strings"
)

var (
	funcMap = tpl.FuncMap{
		"add":               AddFunction,
		"as_html":           AsHTML,
		"to_persian_number": ToPersianNumber,
	}
)

func New(path string) *tpl.Template {
	t := tpl.New("").Funcs(funcMap)
	filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if strings.Contains(path, ".html") {
			_, err = t.ParseFiles(path)
			if err != nil {
				return err
			}
		}
		return err
	})

	return t
}
