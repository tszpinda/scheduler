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
	s := t.Time.Format("02-01-2006 15:04")
	return json.Marshal(s)
}

func (t jsonTime) String() string {
	return time.Time(t.Time).Format("02-01-2006 15:04")
}

type Truck struct {
	Id    int64  `json:"key"`
	Label string `json:"label"`
}

type Event struct {
	Id        int64    `json:"id"`
	TruckId   int64    `json:"truck_id"`
	Text      string   `json:"text"`
	StartDate jsonTime `json:"start_date"`
	EndDate   jsonTime `json:"end_date"`
}

type Schedule struct {
	Events []Event `json:"events"`
	Trucks []Truck `json:"trucks"`
}

type SchedulerService struct {
	gorest.RestService `root:"/scheduler/" consumes:"application/json" produces:"application/json"`
	getSchedule        gorest.EndPoint `method:"GET" path:"/events/{resourceId:int}" output:"Schedule"`
	//getSchedule        gorest.EndPoint `method:"GET" path:"/resources" output:"[]Resource"`
}

func (serv SchedulerService) GetSchedule(resourceId int) (s Schedule) {
	e := make([]Event, 0)
	t := make([]Truck, 0)

	startTime1 := time.Now().Add(time.Duration(2) * time.Hour)
	endTime1 := time.Now().Add(time.Duration(2) * time.Hour)
	startTime2 := time.Now().Add(time.Duration(-3) * time.Hour)
	endTime2 := time.Now().Add(time.Duration(1) * time.Hour)

	e1 := Event{1, 1, "BA1 3AX - 15T Concrete", jsonTime{startTime1}, jsonTime{endTime1}}
	e2 := Event{2, 2, "BA1 3AX - Hire Pickup", jsonTime{startTime2}, jsonTime{endTime2}}
	e = append(e, e1)
	e = append(e, e2)

	t1 := Truck{1, "Truck 1"}
	t2 := Truck{2, "Truck 2"}
	t = append(t, t1)
	t = append(t, t2)

	s.Events = e
	s.Trucks = t
	fmt.Println("incoming request v%", s)
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
