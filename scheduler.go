package main

import (
	"fmt"
	web "github.com/tszpinda/scheduler/web"
	"log"
	"net/http"
	"os"
)

func main() {
	web.InitView()
	web.InitApi()

	var port = os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	fmt.Println("starting app on port: " + port)
	log.Fatal(http.ListenAndServe(":"+port, nil))

}
