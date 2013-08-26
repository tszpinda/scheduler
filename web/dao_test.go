package web

import (
	//"encoding/json"
	m "github.com/tszpinda/scheduler/model"
	//"labix.org/v2/mgo"
	//"labix.org/v2/mgo/bson"
	"log"
	//"reflect"
	"testing"
	"time"
)

func dropDb() {
	session := getSession()
	err := session.DB(databaseName).DropDatabase()
	if err != nil {
		panic(err)
	}
}

func tearDown() {

}

func getCustomer(code, name, contract1, contract2 string, hour1, min1, hour2, min2 int) m.Customer {
	a1 := m.Address{"Greyhound Cottage", "Old Brighton Road", "Wadhurst", "", "TN5 4SQ", "0777256334"}

	st := todayAt(hour1, min1)
	et := st.Add(time.Duration(2) * time.Hour)

	resId := 5
	d1 := m.Delivery{contract1, st, et, resId, "C35", "6.5m",
		265.00, 95.30, []string{"Fibres", "Retarder"},
		"Make sure to drop of additives.",
		a1}

	st = todayAt(hour2, min2).Add(time.Duration(24) * time.Hour)
	et = st.Add(time.Duration(1) * time.Hour).Add(time.Duration(24) * time.Hour)
	d2 := m.Delivery{contract2, st, et, resId, "C40", "7.5m",
		299.00, 195.30, []string{"Fibres", "Retarder"},
		"Make sure to drop of additives.",
		a1}

	return m.Customer{code, name, []m.Delivery{d1, d2}, false}
}

func initTest() {
	log.Println("##########################################################################################")
	log.Println("initTest()")
	log.Println("##########################################################################################")
	databaseName = "unit-test"
	dropDb()
	session := getSession()
	c := session.DB(databaseName).C("scheduler")

	c.Insert(getCustomer("MX-1001", "Max 1001", "014510", "014511", 8, 10, 10, 20))
	c.Insert(getCustomer("MX-1002", "Max 1002", "014520", "014521", 10, 30, 12, 20))
	c.Insert(getCustomer("MX-1003", "Max 1003", "014530", "014531", 8, 25, 10, 50))
	c.Insert(getCustomer("MX-1004", "Max 1004", "014540", "014541", 9, 10, 12, 20))
}

func TestAllCustomers(t *testing.T) {
	initTest()
	defer tearDown()
	log.Println("TestAllCustomers")
	customers := AllCustomers()

	//log.Printf("%+v", customers)

	if len(customers) != 4 {
		log.Panicf("invalid number of customers expected %v but was %v", 4, len(customers))
	}
}

func TestCreateOrUpdate_Create(t *testing.T) {
	initTest()
	defer tearDown()

	CreateOrUpdate(getCustomer("MX-1005", "Max 1005", "014550", "014551", 8, 10, 10, 20))

	count, _ := getSession().DB(databaseName).C("scheduler").Count()
	if count != 5 {
		log.Panicf("invalid number of customers expected %v but was %v", 5, count)
	}
}

func TestCreateOrUpdate_Update(t *testing.T) {
	initTest()
	defer tearDown()

	CreateOrUpdate(getCustomer("MX-1001", "Max 1001", "014510", "014513", 9, 15, 13, 12))

	count, _ := getSession().DB(databaseName).C("scheduler").Count()
	if count != 4 {
		log.Panicf("invalid number of customers expected %v but was %v", 4, count)
	}

	//check if one contract was added (now should have 3), and the other one was updated
	c := DeliveryCount("MX-1001")
	if c != 3 {
		log.Panicf("invalid number of deliveries expected %v but was %v", 3, c)
	}
}
func TestDeliveryCount(t *testing.T) {
	initTest()
	defer tearDown()
	deliveryCount := DeliveryCount("MX-1001")
	if deliveryCount != 2 {
		log.Panicf("invalid number of deliveries expected %v but was %v", 2, deliveryCount)
	}

	CreateOrUpdate(getCustomer("MX-1001", "Max 1001", "014510", "014513", 9, 15, 13, 12))

	deliveryCount = DeliveryCount("MX-1001")
	if deliveryCount != 3 {
		log.Panicf("invalid number of deliveries expected %v but was %v", 3, deliveryCount)
	}

}

func TestRemoveDelivery(t *testing.T) {
	initTest()
	defer tearDown()
	RemoveContract("014510")

	c := DeliveryCount("MX-1001")
	if c != 1 {
		log.Panicf("invalid number of deliveries expected %v but was %v", 1, c)
	}

	RemoveContract("014511")

	c = DeliveryCount("MX-1001")
	if c != 0 {
		log.Panicf("invalid number of deliveries expected %v but was %v", 0, c)
	}

}

func TestCustomerWithAllocatedDelivery(t *testing.T) {
	initTest()
	defer tearDown()
	/*
		8:10 - 10:20
		8:25 - 10:50
		10:30 - 12:20
		9:10 - 12:20 */
	customers := CustomerWithAllocatedDelivery(todayAt(0, 01), tomorrowAt(23, 59))
	c := len(customers)
	if c != 4 {
		log.Panicf("invalid number of deliveries expected %v but was %v", 4, c)
	}

	customers = CustomerWithAllocatedDelivery(todayAt(9, 9), todayAt(23, 59))
	c = len(customers)
	if c != 2 {
		log.Panicf("invalid number of deliveries expected %v but was %v", 2, c)
	}

	customers = CustomerWithAllocatedDelivery(todayAt(7, 10), todayAt(8, 10))
	c = len(customers)
	if c != 1 {
		log.Panicf("invalid number of deliveries expected %v but was %v", 1, c)
	}
}

func TestCustomerWithUnallocatedDelivery(t *testing.T) {
	initTest()
	defer tearDown()

	session := getSession()
	c := session.DB(databaseName).C("scheduler")

	cus := getCustomer("MX-1005", "Max 1005", "014510", "014511", 8, 10, 10, 20)

	var zeroTime time.Time
	cus.Delivery[0].StartTime = zeroTime
	cus.Delivery[0].EndTime = zeroTime
	c.Insert(cus)

	customers := CustomerWithUnallocatedDelivery()
	//should find one customer with one delivery only
	//log.Printf("%+v", customers[0].Delivery)
	count := len(customers)
	if count != 1 {
		log.Panicf("invalid number of customers expected %v but was %v", 1, count)
	}

	count = len(customers[0].Delivery)
	if count != 1 {
		log.Panicf("invalid number of deliveries expected %v but was %v", 1, count)
	}
}

func TestUpdateOverCreditLimit(t *testing.T) {
	customers := AllCustomers()
	for _, c := range customers {
		if c.OverCreditLimit {
			log.Panicln("Expeced all default customers to be in credit limit.")
		}
	}
	UpdateOverCreditLimit("MX-1003", true)
	customers = AllCustomers()
	for _, c := range customers {
		if !c.OverCreditLimit && c.Code == "MX-1003" {
			log.Panicln("Expeced MX-1003 to be over the credit limit.")
		}
		if c.OverCreditLimit && c.Code != "MX-1003" {
			log.Panicln("Expeced non apart MX-1003 to be in the credit limit.")
		}
	}

	UpdateOverCreditLimit("MX-1003", false)
	customers = AllCustomers()
	for _, c := range customers {
		if c.OverCreditLimit {
			log.Panicln("Expeced all default customers to be in credit limit.")
		}
	}
}

func todayAt(hour, min int) time.Time {
	now := time.Now()
	return time.Date(now.Year(), now.Month(), now.Day(), hour, min, 0, 0, time.UTC)
}

func tomorrowAt(hour, min int) time.Time {
	today := todayAt(hour, min)
	return today.Add(time.Duration(24) * time.Hour)
}
