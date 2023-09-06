package belajar_go_web

import (
	"html/template"
	"net/http"
)

func SimpleHTML(writer http.ResponseWriter, request *http.Request) {
	templateText := "<html><body>{{.}}</body></html>"
	// t,err:=template.New("SIMPLE").Parse(templateText)
	// if err != nil {
	// 	panic(err)
	// }

	t:=template.Must(template.New("SIMPLE").Parse(templateText))

	t.ExecuteTemplate(writer,"SIMPLE","Hello HTMLTemplate")
}
