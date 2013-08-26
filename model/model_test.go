package model

import (
	"encoding/json"
	//"fmt"
	"log"
	"testing"
	"time"
)

func initTest() {
}

type pjsonTime struct {
	time.Time
}

func (t pjsonTime) MarshalJSON() ([]byte, error) {
	s := t.Time.Format("02-01-2006 3:4")
	return json.Marshal(s)
}

func (t pjsonTime) String() string {
	return time.Time(t.Time).Format("02-01-2006 3:4")
}

func TestJsonDate(t *testing.T) {
	v := pjsonTime{time.Now()}
	log.Println(v.MarshalJSON())
	log.Println("output:", v.String())
}
