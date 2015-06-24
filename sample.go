package main

import (
    "fmt"
    "gopkg.in/mgo.v2"
    "gopkg.in/mgo.v2/bson"
    "log"
    "time"
)

type Person struct {
    ID        bson.ObjectId `bson:"_id,omitempty"`
    Name      string        `bson:"name"`
    Phone     string        `bson:"phone"`
    Timestamp time.Time
}

var (
    IsDrop = false
)
var PeopleC *mgo.Collection

func Insert() {
    // Insert Datas
    err := PeopleC.Insert(&Person{Name: "Ale", Phone: "+55 53 1234 4321", Timestamp: time.Now()},
        &Person{Name: "Cla", Phone: "+66 33 1234 5678", Timestamp: time.Now()})
    if err != nil {
        log.Println(err)
    }
}

func GetResult() string {

    result := &Person{}

    err := PeopleC.Find(bson.M{"name": "Ale"}).Select(bson.M{"phone": 0}).One(&result)
    if err != nil {
        panic(err)
    }

    return fmt.Sprintf("%s", result.Name)
}
