package main

import (
	"code.google.com/p/gorest"
	"encoding/json"
	"fmt"
	//mux "github.com/gorilla/mux"
	view "github.com/tszpinda/scheduler/web"
	"net/http"
	"os"
	"time"
)

type jsonTime struct {
	time.Time
}

func (t jsonTime) MarshalJSON() ([]byte, error) {
	s := t.Time.Format("02-01-2006 3:4")
	return json.Marshal(s)
}

func (t jsonTime) String() string {
	return time.Time(t.Time).Format("02-01-2006 3:4")
}

type Event struct {
	Id        float64  `json:"id"`
	Text      string   `json:"text"`
	StartDate jsonTime `json:"start_date"`
	EndDate   jsonTime `json:"end_date"`
}

type SchedulerService struct {
	gorest.RestService `root:"/scheduler/" consumes:"application/json" produces:"application/json"`
	getSchedule        gorest.EndPoint `method:"GET" path:"/event/{resourceId:int}" output:"[]Event"`
}

func (serv SchedulerService) GetSchedule(resourceId int) (e []Event) {
	endTime1 := time.Now().Add(time.Duration(2) * time.Hour)

	e1 := Event{1, "Replace boiler", jsonTime{time.Now()}, jsonTime{endTime1}}
	e2 := Event{2, "Replace bitchen", jsonTime{time.Now()}, jsonTime{time.Now()}}
	e = append(e, e1)
	e = append(e, e2)
	fmt.Println("incoming request v%", e)
	return
}

func main() {
	gorest.RegisterService(new(SchedulerService))

	view.Mount()
	http.Handle("/", gorest.Handle())

	var port = os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	fmt.Println("starting app on port: " + port)
	http.ListenAndServe(":"+port, nil)
	fmt.Println("ups")
}
