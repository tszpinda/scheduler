package web

import (
	m "github.com/tszpinda/scheduler/model"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	"log"
	//"reflect"
	"time"
)

var (
	mgoSession   *mgo.Session
	databaseName = "test"
)

func getSession() *mgo.Session {
	if mgoSession == nil {
		var err error
		mgoSession, err = mgo.Dial("localhost")
		if err != nil {
			panic(err) // no, not really
		}
	}
	return mgoSession.Clone()
}

func schedulerCollection(s func(*mgo.Collection) error) error {
	return withCollection("scheduler", s)
}
func withCollection(collection string, s func(*mgo.Collection) error) error {
	session := getSession()
	defer session.Close()
	c := session.DB(databaseName).C(collection)
	return s(c)
}

func AllCustomers() (customers []m.Customer) {
	query := func(c *mgo.Collection) error {
		fn := c.Find(nil).All(&customers)
		return fn
	}

	search := func() error {
		return schedulerCollection(query)
	}

	err := search()
	if err != nil {
		panic(err)
	}

	return
}

func CustomerWithAllocatedDelivery(startTime time.Time, endTime time.Time) (customers []m.Customer) {
	session := getSession()
	defer session.Close()
	c := session.DB(databaseName).C("scheduler")

	/*db.scheduler.aggregate(
	{$unwind: "$delivery"},
	{$match : {"delivery.starttime": {$gte : new Date("2013-08-26")}}},
	{$group : {_id:{"name":"$name","code":"$code"}, delivery:{"$push":"$delivery"}}},
	{$project: {"name" : "$_id.name", code:"$_id.code", "delivery" :1, "_id" : 0}})
	*/
	pipe := c.Pipe([]bson.M{
		{"$unwind": "$delivery"},
		{"$match": bson.M{"delivery.starttime": bson.M{"$gte": startTime, "$lte": endTime}}},
		{"$group": bson.M{"_id": bson.M{"code": "$code", "name": "$name", "overcreditlimit": "$overcreditlimit"},
			"delivery": bson.M{"$push": "$delivery"}}},
		{"$project": bson.M{"code": "$_id.code", "name": "$_id.name", "delivery": 1, "_id": 0}},
	})

	/*
		iter := pipe.Iter()
		for {
			result := m.Customer{}
			if iter.Next(&result) {
				for r, i := range result.Delivery {
					log.Println("r:", r, result.Code, i.StartTime)
				}
			} else {
				break
			}
		}
	*/

	result := make([]m.Customer, 0)
	pipe.All(&result)
	log.Printf("Found '%v' customers with deliveries between '%v' - '%v' ", len(result), startTime, endTime)
	return result
}

func CustomerWithUnallocatedDelivery() (customers []m.Customer) {
	var zeroTime time.Time
	return CustomerWithAllocatedDelivery(zeroTime, zeroTime)
}

func FindCustomer(code string) (customer m.Customer, found bool) {
	query := func(c *mgo.Collection) error {
		fn := c.Find(bson.M{"code": code}).One(&customer)
		return fn
	}

	search := func() error {
		return schedulerCollection(query)
	}

	err := search()
	found = true
	if err != nil {
		if err.Error() == "not found" {
			found = false
		} else {
			panic(err)
		}
	}

	return
}
func CreateOrUpdate(customer m.Customer) {
	//TODO - make in a query instead of looping in the code for add/update
	query := func(c *mgo.Collection) error {
		cus, found := FindCustomer(customer.Code)
		if !found {
			fn := c.Insert(customer)
			return fn
		} else {
			for _, newD := range customer.Delivery {
				delFound := false
				for i, currentD := range cus.Delivery {
					//update if found
					if currentD.ContractNumber == newD.ContractNumber {
						cus.Delivery[i] = newD
						delFound = true
					}
				}
				if !delFound {
					cus.Delivery = append(cus.Delivery, newD)
				}
			}

			fn := c.Update(bson.M{"code": customer.Code}, cus)
			return fn
		}
	}
	err := schedulerCollection(query)
	if err != nil {
		panic(err)
	}
}

func DeliveryCount(cusCode string) int {
	/*
		db.scheduler.aggregate(
			{$unwind: "$delivery"},
			{$match : {"code":"MX-1001"}},
			{$project: {"delivery.contractnumber" : 1, _id:0}},
			{$group : {_id:"$delivery.contractnumber"}}
			).result.length
	*/
	session := getSession()
	defer session.Close()
	c := session.DB(databaseName).C("scheduler")

	pipe := c.Pipe([]bson.M{{"$unwind": "$delivery"}, {"$match": bson.M{"code": cusCode}}, {"$project": bson.M{"delivery.contractnumber": 1, "_id": 0}}, {"$group": bson.M{"_id": "$delivery.contractnumber"}}})

	var m bson.D
	pipe.All(&m)
	return len(m)
}

func RemoveContract(contract string) {
	//db.scheduler.update({"delivery.contractnumber": "014511"}, {"$pull": {"delivery" : {"contractnumber":"014511"}}})
	log.Println("remove contract:", contract)

	query := func(c *mgo.Collection) error {
		q := bson.M{"delivery.contractnumber": contract}
		rq := bson.M{"$pull": bson.M{"delivery": bson.M{"contractnumber": contract}}}

		return c.Update(q, rq)
	}

	search := func() error {
		return schedulerCollection(query)
	}

	err := search()
	if err != nil {
		panic(err)
	}

	return
}

func UpdateOverCreditLimit(cusCode string, isOverCreditLimit bool) {
	query := func(c *mgo.Collection) error {

		q := bson.M{"code": cusCode}
		rq := bson.M{"$set": bson.M{"overcreditlimit": isOverCreditLimit}}
		return c.Update(q, rq)
	}

	search := func() error {
		return schedulerCollection(query)
	}

	err := search()
	if err != nil {
		panic(err)
	}

	return
}
