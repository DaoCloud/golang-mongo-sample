package main

import (
	"fmt"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	"os"
	"reflect"
	"time"
)

type Person struct {
	ID        bson.ObjectId `bson:"_id,omitempty"`
	Name      string
	Phone     string
	Timestamp time.Time
}

var (
	IsDrop = true
)

func GetResult() string {
	conn := ""
	if len(os.Getenv("MONGODB_USERNAME")) > 0 {
		conn += os.Getenv("MONGODB_USERNAME")

		if len(os.Getenv("MONGODB_PASSWORD")) > 0 {
			conn += ":" + os.Getenv("MONGODB_PASSWORD")
		}

		conn += "@"
	}

	if len(os.Getenv("MONGODB_PORT_27017_TCP_ADDR")) > 0 {
		conn += os.Getenv("MONGODB_PORT_27017_TCP_ADDR")
	} else {
		conn += "localhost"
	}

	if len(os.Getenv("MONGODB_PORT_27017_TCP_PORT")) > 0 {
		conn += ":" + os.Getenv("MONGODB_PORT_27017_TCP_PORT")
	} else {
		conn += ":27017"
	}
	// defaultly using "test" as the db instance
	db := "test"

	if len(os.Getenv("MONGODB_INSTANCE_NAME")) > 0 {
		db = os.Getenv("MONGODB_INSTANCE_NAME")
	}

	conn += "/" + db

	session, err := mgo.Dial(conn)
	if err != nil {
		panic(err)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)

	// Drop Database
	if IsDrop {
		err = session.DB(db).DropDatabase()
		if err != nil {
			panic(err)
		}
	}

	// Collection People
	c := session.DB(db).C("people")

	// Index
	index := mgo.Index{
		Key:        []string{"name", "phone"},
		Unique:     true,
		DropDups:   true,
		Background: true,
		Sparse:     true,
	}
	err = c.EnsureIndex(index)
	if err != nil {
		panic(err)
	}

	// Insert Datas
	err = c.Insert(&Person{Name: "Ale", Phone: "+55 53 1234 4321", Timestamp: time.Now()},
		&Person{Name: "Cla", Phone: "+66 33 1234 5678", Timestamp: time.Now()})
	if err != nil {
		panic(err)
	}

	result := Person{}
	err = c.Find(bson.M{"name": "Ale"}).Select(bson.M{"phone": 0}).One(&result)
	if err != nil {
		panic(err)
	}
	return fmt.Sprintf("%v", reflect.ValueOf(&result).Elem().Field(1))
}
