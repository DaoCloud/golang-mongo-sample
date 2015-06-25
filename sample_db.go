package main

import (
    "gopkg.in/mgo.v2"
    "log"
    "os"
)

var DB *mgo.Database

func MustConnectMongo() {
    if err := ConnectMongo(); err != nil {
        panic(err)
    }
    InitDB()
}

func InitDB() {
    defer func() {
        if e := recover(); e != nil {
            log.Println(e)
        }
    }()
    Insert(&Person{Name: "Ale", Phone: "+55 53 1234 4321"})
    Insert(&Person{Name: "Cla", Phone: "+66 33 1234 5678"})
}

func ConnectMongo() error {
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
        return err
    }

    DB = session.DB(db)
    peopleC = DB.C("people")

    return nil
}

func Drop() {
    // Drop Collection
    err := peopleC.DropCollection()
    if err != nil {
        log.Println(err)
    }
}
