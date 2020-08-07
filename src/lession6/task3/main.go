package main

import (
	"html/template"
	"log"
	"net/http"
)

func indexHandler(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("Index page handler"))
}

func helloHandler(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "text/html")
	data := struct {
		Name string
	}{req.FormValue("name")}

	t := `<doctype html>
<html>
	<head>
    	<title>Hello Page</title>
	</head>
	<body>
    	Hello{{if .Name}} {{.Name}}{{end}}!
	</body>
</html>`

	tpl := template.Must(template.New("").Parse(t))
	_ = tpl.Execute(res, data)

}

func main() {

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/hello", helloHandler)
	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		log.Fatal(err)
	}
}
