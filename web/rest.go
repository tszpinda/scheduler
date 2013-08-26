package web

import (
	"code.google.com/p/gorest"
	"fmt"
	m "github.com/tszpinda/scheduler/model"
	"net/http"

//	"time"
)

func InitApi() {
	gorest.RegisterService(new(SchedulerService))
	http.Handle("/", gorest.Handle())

}

type SchedulerService struct {
	gorest.RestService `root:"/api/" consumes:"application/json" produces:"application/json"`
	getSchedule        gorest.EndPoint `method:"GET" path:"/events/{resourceId:int}" output:"Schedule"`
	/*
		PUT  /api/user/login/${username}/${fullName} - creates schedulerKey: "XXX111"
		GET  /api/user/logout/${schedulerKey}
		GET  /api/user/auth/${schedulerKey} - logs in and creates cookie
		POST /api/delivery/import
		GET  /api/delivery/${startDate}
		GET  /api/delivery/${startDate}/${endDate}
		GET  /api/delivery/${contractCode}
		PUT  /api/delivery/${contractCode}/fullAllocation/${fullAllocation}
		PUT  /api/customer/${customerCode}/overCredit/${overCreditLimit}
	*/
}

func (serv SchedulerService) GetSchedule(resourceId int) (s m.Schedule) {

	//e := make([]m.Event, 0)

	t := make([]m.Resource, 0)
	/*
		const format = "2006-01-02 15:04"
		startTime1, _ := time.Parse(format, "2013-08-23 08:05")
		endTime1 := startTime1.Add(time.Duration(2) * time.Hour)

		startTime2, _ := time.Parse(format, "2013-08-23 08:35")
		endTime2 := startTime2.Add(time.Duration(1) * time.Hour)

		e1 := m.Event{1, 1, "BA1 3AX - 15T Concrete", m.JsonTime{startTime1}, m.JsonTime{endTime1}}
		e2 := m.Event{2, 2, "BA1 3AX - Hire Pickup", m.JsonTime{startTime2}, m.JsonTime{endTime2}}
		e = append(e, e1)
		e = append(e, e2)
	*/
	t1 := m.Resource{1, 1, "Truck 1"}
	t2 := m.Resource{2, 2, "Truck 2"}
	t = append(t, t1)
	t = append(t, t2)

	//s.Events = e
	s.Collections.Resources = t

	fmt.Println("incoming request v+%", s)
	return
}
