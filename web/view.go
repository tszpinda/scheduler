package web

import (
	//"fmt"
	"html/template"
	"log"
	"net/http"
)

func Init() {
	//serve public/static directory
	fileHandler := http.StripPrefix("/ui/static/", http.FileServer(http.Dir("./public/static")))
	http.Handle("/ui/static/", fileHandler)

	http.HandleFunc("/ui/scheduler", schedulerPage)
}

func serveSingle(pattern string, filename string) {
	http.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, filename)
	})
}

func schedulerPage(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		var templateName = "public/tmpl/scheduler.gtpl"
		t, err := template.ParseFiles(templateName)
		if err != nil {
			log.Fatal("unable to parse ", templateName, err)
		}
		t.Execute(w, nil)
	}
}
