package welcome

import (
	"html/template"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	var tmplPath string = "views/template/"
	var welcomePath string = "views/welcome/"
	t, _ := template.ParseFiles(
		tmplPath+"header.tmpl",
		welcomePath+"index.html",
		tmplPath+"footer.tmpl",
	)
	var data = map[string]interface{}{
		"title": "Golang server works!",
	}
	t.ExecuteTemplate(w, "body", data)
	t.Execute(w, nil)
}
