package web

import (
	//"fmt"
	"html/template"
	"log"
	"net/http"
)

func Mount() {
	pfx := "/ui/public/static/img/"
	h := http.StripPrefix(pfx, http.FileServer(http.Dir("./public/static/img")))
	http.Handle(pfx, h)

	serveSingle("/ui/static/js/main.js", "./public/static/js/main.js")
	serveSingle("/ui/static/js/jquery.js", "./public/static/js/jquery.js")
	serveSingle("/ui/static/bootstrap/js/bootstrap.js", "./public/static/bootstrap/js/bootstrap.js")
	serveSingle("/ui/static/bootstrap/css/bootstrap.css", "./public/static/bootstrap/css/bootstrap.min.css")
	serveSingle("/ui/static/bootstrap/css/bootstrap-responsive.css", "./public/static/bootstrap/css/bootstrap-responsive.min.css")
	serveSingle("/ui/static/bootstrap/css/main.css", "./public/static/bootstrap/css/main.css")
	serveSingle("/favicon.ico", "./public/static/favicon.ico")

	serveSingle("/ui/static/dhtmlxscheduler.css", "./public/static/dhtmlxscheduler.css")
	serveSingle("/ui/static/dhtmlxscheduler.js", "./public/static/dhtmlxscheduler.js")
	serveSingle("/ui/static/dhtmlxscheduler-units.js", "./public/static/dhtmlxscheduler-units.js")

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
