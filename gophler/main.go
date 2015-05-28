package main

import (
	"bytes"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"html/template"
	"log"
	"net/http"
)

var layoutFuncs = template.FuncMap{
	"yield": func() (string, error) {
		return "", fmt.Errorf("yield called inappropriately")
	},
}

var layout = template.Must(
	template.
		New("layout.html").
		Funcs(layoutFuncs).
		ParseFiles("templates/layout.html"),
)

var templates = template.Must(template.New("t").ParseGlob("templates/**/*.html"))

func RenderTemplate(w http.ResponseWriter, r *http.Request, name string, data interface{}) {
	funcs := template.FuncMap{
		"yield": func() (template.HTML, error) {
			buf := bytes.NewBuffer(nil)
			err := templates.ExecuteTemplate(buf, name, data)
			return template.HTML(buf.String()), err
		},
	}

	layoutClone, _ := layout.Clone()
	layoutClone.Funcs(funcs)

	err := layoutClone.Execute(w, data)

	if err != nil {
		http.Error(w, fmt.Sprintf(errorTemplate, name, err), http.StatusInternalServerError)
	}
}

var errorTemplate = `
<html>
<body>
<h1> Error rendering template %s</h1>
<p>%s</p>
</body>
</html>
`

func HandleUserImage(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	//userID := params.ByName("user_id")
	//imageID := params.ByName("image_id")
}

func main() {
	router := httprouter.New()
	router.Handle("GET", "/", HandleHome)
	router.ServeFiles("/assets/*filepath", http.Dir("assets/"))
	router.Handle("GET", "/register", HandleUserNew)
	router.Handle("POST", "/register", HandleUserCreate)

	middleware := Middleware{}
	middleware.Add(router)

	log.Fatal(http.ListenAndServe(":3000", middleware))
}
