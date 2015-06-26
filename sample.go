package main

import (
    "gopkg.in/mgo.v2"
    "gopkg.in/mgo.v2/bson"
    // "log"
)

type Person struct {
    Name  string `bson:"name"`
    Phone string `bson:"phone"`
}

var peopleC *mgo.Collection

func Insert(person *Person) {

    if person.Name == "" || GetResult(person.Name) != "" {
        panic("insert conflict")
    }
    // Insert Datas
    err := peopleC.Insert(person)
    if err != nil {
        panic(err)
    }
}

func List() []*Person {
    query := peopleC.Find(nil).Select(bson.M{"_id": 0}).Sort("_id")
    list := []*Person{}
    err := query.All(&list)
    if err != nil {
        panic(err)
    }

    return list
}

func GetResult(name string) string {
    result := &Person{}

    err := peopleC.Find(bson.M{"name": name}).Select(bson.M{"_id": 0}).One(&result)
    if err != nil && err != mgo.ErrNotFound {
        panic(err)
    }

    return result.Name
}
