package model

import (
	"encoding/json"
	"fmt"
	"time"
)

type JsonTime struct {
	time.Time
}

func (t JsonTime) MarshalJSON() ([]byte, error) {
	s := t.Time.Format("2006-01-02 15:04")
	return json.Marshal(s)
}
func (t JsonTime) String() string {
	fmt.Printf("%+v", t.Day())
	return time.Time(t.Time).Format("2006-01-02 15:04")
}

type Resource struct {
	Id    int64  `json:"id"`
	Value int64  `json:"value"`
	Label string `json:"label"`
}

type Event struct {
	Id         int64    `json:"id"`
	ResourceId int64    `json:"type"`
	Notes      string   `json:"text"`
	StartDate  JsonTime `json:"start_date"`
	EndDate    JsonTime `json:"end_date"`
}
type Collections struct {
	Resources []Resource `json:"type"`
}
type Schedule struct {
	Events      []Event `json:"data"`
	Collections `json:"collections"`
}

type Customer struct {
	Code            string     `bson:"code"`
	Name            string     `bson:"name"`
	Delivery        []Delivery `bson:"delivery"`
	OverCreditLimit bool       `bson:"overcreditlimit"`
}
type Address struct {
	Line1    string `json:"line1"`
	Line2    string `json:"line2"`
	Line3    string `json:"line3"`
	Line4    string `json:"line4"`
	Postcode string `json:"postcode"`
	Phone    string `json:"phone"`
}
type Delivery struct {
	ContractNumber string    `json:"contractNumber"`
	StartTime      time.Time `json:"startTime"`
	EndTime        time.Time `json:"endTime"`
	ResourceId     int       `json:"resourceId"`
	Mix            string    `json:"mix"`
	Qty            string    `json:"qty"`
	CmPrice        float64   `json:"cmPrice"`
	DeliveryPrice  float64   `json:"deliveryPrice"`
	Additives      []string  `json:"additives"`
	Notes          string    `json:"notes"`
	Address        Address   `json:"address"`
}

/*{ //start delivery 1
		customer : {
			“code”        : “CUS-1”
			"name" 	        : "Design Build Limited",
			},
		delivery : {
			“startTime”     	: "2013/03/01 15:38", //if startTime & endTime empty delivery will be in the holding area
			”endTime”       	: "2013/03/01 17:38", //waiting to be dragged onto the scheduler
                                    “resourceId”       : 255, //resource (truck/driver) which must be available on the day of delivery
			”mix”			  	: "C35",
			”qty”			  	: "6.5m",
			”cmPrice”		  	: "265.00",
			”deliveryPrice” 	: "95.30", //not required
			”additives”		: ["Fibres", "Retarder"],
			”contractNumber”	: "014528",
			”notes”			: "Some longish notes can go here, whatever text they want",
			“address”		: {
				"line1"		: "Greyhound Cottage",
				"line2"		: "Old Brighton Road",
				"line3"		: "Wadhurst",
				"line4"		: "",
				"postcode"		: "TN5 4SQ",
				"phone"			: "07772566334"
			}
		}
	};*/
